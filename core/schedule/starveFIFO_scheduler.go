package schedule

import (
	"container/heap"
	"fmt"
	"github.com/RosettaFlow/Carrier-Go/common"
	"github.com/RosettaFlow/Carrier-Go/common/timeutils"
	towTypes "github.com/RosettaFlow/Carrier-Go/consensus/twopc/types"
	"github.com/RosettaFlow/Carrier-Go/core/evengine"
	"github.com/RosettaFlow/Carrier-Go/core/iface"
	"github.com/RosettaFlow/Carrier-Go/core/resource"
	"github.com/RosettaFlow/Carrier-Go/grpclient"
	pb "github.com/RosettaFlow/Carrier-Go/lib/api"
	apipb "github.com/RosettaFlow/Carrier-Go/lib/common"
	libTypes "github.com/RosettaFlow/Carrier-Go/lib/types"
	"github.com/RosettaFlow/Carrier-Go/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

const (
	ReschedMaxCount             = 8
	StarveTerm                  = 3

	electionOrgCondition        = 10000
	electionLocalSeed           = 2
	//taskComputeOrgCount         = 3
)

var (
	ErrEnoughResourceOrgCountLessCalculateCount = fmt.Errorf("the enough resource org count is less calculate count")
	ErrEnoughInternalResourceCount              = fmt.Errorf("has not enough internal resource count")
)

type SchedulerStarveFIFO struct {
	internalNodeSet *grpclient.InternalResourceClientSet
	resourceMng     *resource.Manager
	// the local task into this queue, first
	queue *types.TaskBullets
	// the very very starve local task by priority
	starveQueue *types.TaskBullets
	// the scheduling task, it is ejected from the queue (taskId -> taskBullet)
	schedulings   map[string]*types.TaskBullet
	queueMutex  sync.Mutex

	// fetch local task from taskManager`
	//localTaskMsgCh chan types.TaskMsgs
	//// send local task scheduled to `Consensus`
	//needConsensusTaskCh chan<- *types.ConsensusTaskWrap
	//// receive remote task to replay from `Consensus`
	//replayScheduleTaskCh <-chan *types.ReplayScheduleTaskWrap
	//// 发送经过调度好的 task 交给 taskManager 去分发给自己的 Fighter-Py
	//doneSchedTaskCh chan<- *types.DoneScheduleTaskChWrap
	quit            chan struct{}
	eventEngine     *evengine.EventEngine
	//dataCenter      iface.ForResourceDB
	err             error
}

func NewSchedulerStarveFIFO(
	internalNodeSet *grpclient.InternalResourceClientSet,
	eventEngine *evengine.EventEngine,
	mng *resource.Manager,
) *SchedulerStarveFIFO {

	return &SchedulerStarveFIFO{
		internalNodeSet:      internalNodeSet,
		resourceMng:          mng,
		queue:                new(types.TaskBullets),
		starveQueue:          new(types.TaskBullets),
		schedulings: 		  make(map[string]*types.TaskBullet),
		//localTaskMsgCh:       localTaskMsgCh,

		//dataCenter:           dataCenter,
		eventEngine:          eventEngine,
		quit:                 make(chan struct{}),
	}
}
func (sche *SchedulerStarveFIFO) loop() {

	for {
		select {
		// From taskManager
		// 新task Msg 到来, 主动触发 调度
		case tasks := <-sche.localTaskMsgCh:

			for _, task := range tasks {

				log.Debugf("Received local task, taskId: {%s}, partyId: {%s}", task.TaskId, task.Data.TaskData().PartyId)

				if err := sche.resourceMng.GetDB().StoreLocalTask(task.Data); nil != err {

					e := fmt.Errorf("store local task failed, taskId {%s}, %s", task.Data.TaskData().TaskId, err)

					log.Errorf("failed to call StoreLocalTask on SchedulerStarveFIFO with schedule task, err: {%s}", e.Error())

					task.Data.TaskData().EndAt = uint64(timeutils.UnixMsec())
					task.Data.TaskData().Reason = e.Error()
					task.Data.TaskData().State = types.TaskStateFailed.String()

					identityId, _ := sche.resourceMng.GetDB().GetIdentityId()
					event := sche.eventEngine.GenerateEvent(evengine.TaskDiscarded.Type, task.Data.TaskData().TaskId, identityId, e.Error())
					task.Data.TaskData().EventCount = 1
					task.Data.TaskData().TaskEventList = []*libTypes.TaskEvent{event}

					if err = sche.resourceMng.GetDB().InsertTask(task.Data); nil != err {
						log.Errorf("Failed to save task to datacenter, taskId: {%s}", task.Data.TaskData().TaskId)
						continue
					}
				}

				bullet := types.NewTaskBulletByTaskMsg(task)
				sche.addTaskBullet(bullet)
				sche.trySchedule()
			}

		// From Consensus Engine, from remote peer
		// 让自己的Scheduler 重演选举
		case replayScheduleTask := <-sche.replayScheduleTaskCh:

			log.Debugf("Received remote task, taskId: {%s}, partyId: {%s}, myself task role: {%s}",
				replayScheduleTask.Task.TaskId(), replayScheduleTask.PartyId, replayScheduleTask.Role.String())

			if err := sche.resourceMng.GetDB().StoreLocalTask(replayScheduleTask.Task); nil != err {
				e := fmt.Errorf("store remote task failed, taskId: {%s}, {%s}",
					replayScheduleTask.Task.TaskId(), err)

				log.Errorf("failed to call StoreLocalTask on SchedulerStarveFIFO with replay schedule task, err: {%s}", err)

				replayScheduleTask.SendFailedResult(replayScheduleTask.Task.TaskId(),
					e)
				continue
			}

			go sche.replaySchedule(replayScheduleTask)



		case <-sche.quit:
			log.Info("Stopped SchedulerStarveFIFO ...")
			return
		}

	}
}

func (sche *SchedulerStarveFIFO) Start() error {
	go sche.loop()
	log.Info("Started SchedulerStarveFIFO ...")
	return nil
}
func (sche *SchedulerStarveFIFO) Stop() error {
	close(sche.quit)
	return nil
}
func (sche *SchedulerStarveFIFO) Error() error { return sche.err }
func (sche *SchedulerStarveFIFO) Name() string { return "SchedulerStarveFIFO" }
func (sche *SchedulerStarveFIFO) addTaskBullet(bullet *types.TaskBullet) {
	heap.Push(sche.queue, bullet)
}
func (sche *SchedulerStarveFIFO) trySchedule() error {

	sche.increaseTaskTerm()

	var bullet *types.TaskBullet

	if sche.starveQueue.Len() != 0 {
		x := heap.Pop(sche.starveQueue)
		bullet = x.(*types.TaskBullet)
	} else {
		if sche.queue.Len() != 0 {
			x := heap.Pop(sche.queue)
			bullet = x.(*types.TaskBullet)
		} else {
			//log.Info("There is not task on FIFO schedule, finished try schedule timer ...")
			return nil
		}
	}

	go func() {
		task := bullet.UnschedTask

		repushFn := func(bullet *types.TaskBullet) {

			bullet.IncreaseResched()
			if bullet.Resched > ReschedMaxCount {
				// 被丢弃掉的 task  也要清理掉  本地任务的资源, 并提交到数据中心 ...
				log.Errorf("The number of times the task has been rescheduled exceeds the expected threshold, taskId: {%s}, reschedCount: {%d}, max threshold: {%d}",
					bullet.UnschedTask.Data.TaskId(), bullet.Resched, ReschedMaxCount)
				sche.eventEngine.StoreEvent(sche.eventEngine.GenerateEvent(evengine.TaskDiscarded.Type,
					bullet.UnschedTask.Data.TaskId(), bullet.UnschedTask.Data.TaskData().IdentityId, fmt.Sprintf(
						"Task rescheduled exceeds the expected threshold")))

				failedTask := &types.DoneScheduleTaskChWrap{
					ProposalId:   common.Hash{},
					SelfTaskRole: types.TaskOwner,
					SelfIdentity: &apipb.TaskOrganization{
						PartyId:  task.Data.TaskData().PartyId,
						IdentityId: task.Data.TaskData().IdentityId,
						NodeId:   task.Data.TaskData().NodeId,
						NodeName: task.Data.TaskData().NodeName,
					},
					Task: &types.ConsensusScheduleTask{
						TaskDir:   types.SendTaskDir,
						TaskState: types.TaskStateFailed,
						SchedTask: types.ConvertTaskMsgToTaskWithPowers(task.Data, nil),
					},
					ResultCh: make(chan *types.TaskResultMsgWrap, 0),
				}
				sche.SendTaskToTaskManager(failedTask)
			} else {
				if bullet.Starve {
					// 被丢弃掉的 task  也要清理掉  本地任务的资源, 并提交到数据中心 ...
					log.Debugf("Task repush  into starve queue, taskId: {%s}, reschedCount: {%d}, max threshold: {%d}",
						bullet.UnschedTask.Data.TaskId(), bullet.Resched, ReschedMaxCount)
					heap.Push(sche.starveQueue, bullet)
				} else {
					log.Debugf("Task repush  into queue, taskId: {%s}, reschedCount: {%d}, max threshold: {%d}",
						bullet.UnschedTask.Data.TaskId(), bullet.Resched, ReschedMaxCount)
					heap.Push(sche.queue, bullet)
				}
			}
		}

		cost := &towTypes.TaskOperationCost{
			Mem:       task.Data.TaskData().OperationCost.CostMem,
			Processor: uint64(task.Data.TaskData().OperationCost.CostProcessor),
			Bandwidth: task.Data.TaskData().OperationCost.CostBandwidth,
		}

		log.Debugf("Call trySchedule start, taskId: {%s}, partyId: {%s}, taskCost: {%s}",
			task.Data.TaskData().TaskId, task.Data.TaskData().PartyId, cost.String())



		dataIdentityIdCache := make(map[string]struct{})

		// 选出 关于自己 metaDataId 所在的 dataNode
		var metaDataId string

		for _, dataSupplier := range task.Data.TaskData().DataSupplier {
			dataIdentityIdCache[dataSupplier.MemberInfo.IdentityId] = struct{}{}

			// 取出 自己的 disk used 信息, identity 和 partyId 都一致, 才是同一个人 ..
			if task.Data.TaskData().IdentityId == dataSupplier.MemberInfo.IdentityId &&
				task.Data.TaskData().PartyId == dataSupplier.MemberInfo.PartyId {
				// 【选出 发起方 自己的 metaDataId 的 file 对应的  dataNode [ip:port]】
				metaDataId = dataSupplier.MetadataId
			}

		}

		for _, receiver := range task.Data.TaskData().Receivers {
			dataIdentityIdCache[receiver.Receiver.IdentityId] = struct{}{}
		}
		// 【选出 其他组织的算力】
		powers, err := sche.electionConputeOrg(task.PowerPartyIds, dataIdentityIdCache, cost)
		if nil != err {
			log.Errorf("Failed to election powers org on trySchedule, taskId: {%s}, err: {%s}", task.Data.TaskId(), err)
			sche.eventEngine.StoreEvent(sche.eventEngine.GenerateEvent(evengine.TaskFailedConsensus.Type,
				task.Data.TaskData().TaskId, task.Data.TaskData().IdentityId, err.Error()))
			repushFn(bullet)
			return
		}

		log.Debugf("Succeed to election powers org on trySchedule, taskId {%s}, powers: %s", task.Data.TaskId(), utilOrgPowerArrString(powers))

		// 获取 metaData 所在的dataNode 资源
		dataResourceDiskUsed, err := sche.resourceMng.GetDB().QueryDataResourceDiskUsed(metaDataId)
		if nil != err {
			log.Errorf("Failed to query localResourceId By MetaDataId of task owner, taskId: {%s}, metaDataId: {%s}, err: {%s}",
				task.Data.TaskId(), metaDataId, err)
			sche.eventEngine.StoreEvent(sche.eventEngine.GenerateEvent(evengine.TaskFailedConsensus.Type,
				bullet.UnschedTask.Data.TaskId(), bullet.UnschedTask.Data.TaskData().IdentityId, err.Error()))
			repushFn(bullet)
			return
		}
		dataNodeResource, err := sche.resourceMng.GetDB().GetRegisterNode(pb.PrefixTypeDataNode, dataResourceDiskUsed.GetNodeId())
		if nil != err {
			log.Errorf("Failed to query localResourceInfo By dataNodeId: {%s}, taskId: {%s}, err: {%s}",
				dataResourceDiskUsed.GetNodeId(), task.Data.TaskId(), err)
			sche.eventEngine.StoreEvent(sche.eventEngine.GenerateEvent(evengine.TaskFailedConsensus.Type,
				bullet.UnschedTask.Data.TaskId(), bullet.UnschedTask.Data.TaskData().IdentityId, err.Error()))
			repushFn(bullet)
			return
		}

		// Send task to consensus Engine to consensus.
		scheduleTask := types.ConvertTaskMsgToTaskWithPowers(task.Data, powers)
		// restore task by power
		if err := sche.resourceMng.GetDB().StoreLocalTask(task.Data); nil != err {
			log.Errorf("Failed tp update local task by election powers on `trySchedule()`, taskId: {%s}, err: {%s}", task.Data.TaskId(), err)
		}

		log.Debugf("Succeed dataSupplier dataNode on trySchedule(), taskId: {%s}, dataNode: %s", task.Data.TaskId(), dataNodeResource.String())

		toConsensusTask := &types.ConsensusTaskWrap{
			Task: scheduleTask,
			OwnerDataResource: &types.PrepareVoteResource{
				Id:      dataNodeResource.Id,
				Ip:      dataNodeResource.ExternalIp,
				Port:    dataNodeResource.ExternalPort,
				PartyId: task.Data.TaskData().PartyId,
			},
			ResultCh: make(chan *types.ConsensusResult, 0),
		}
		sche.SendTaskToConsensus(toConsensusTask)
		consensusRes := toConsensusTask.RecvResult()

		log.Debugf("Received task result from consensus, taskId: {%s}, result status: {%s}", consensusRes.TaskId, consensusRes.Status)

		// Consensus failed, task needs to be suspended and rescheduled
		if consensusRes.Status == types.TaskConsensusInterrupt {
			sche.eventEngine.StoreEvent(sche.eventEngine.GenerateEvent(evengine.TaskFailedConsensus.Type,
				task.Data.TaskId(), task.Data.TaskData().IdentityId, consensusRes.Err.Error()))
			repushFn(bullet)
			return
		}

	}()

	return nil
}
func (sche *SchedulerStarveFIFO) replaySchedule(replayScheduleTask *types.ReplayScheduleTaskWrap) {

	cost := &towTypes.TaskOperationCost{
		Mem:       replayScheduleTask.Task.GetTaskData().OperationCost.CostMem,
		Processor: uint64(replayScheduleTask.Task.GetTaskData().OperationCost.CostProcessor),
		Bandwidth: replayScheduleTask.Task.GetTaskData().OperationCost.CostBandwidth,
	}

	log.Debugf("Call replaySchedule start, taskId: {%s}, taskRole: {%s}, partyId: {%s}, taskCost: {%s}",
		replayScheduleTask.Task.GetTaskId(), replayScheduleTask.Role.String(), replayScheduleTask.PartyId, cost.String())

	selfIdentityId, err := sche.resourceMng.GetDB().GetIdentityId()
	if nil != err {
		log.Errorf("Failed to query self identityInfo, taskId: {%s}, err: {%s}", replayScheduleTask.Task.GetTaskId(), err)
		replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(), err)
		return
	}
	// 任务的 重演者 不应该是 任务的发起者
	if selfIdentityId == replayScheduleTask.Task.GetTaskData().IdentityId {
		log.Errorf("failed to validate task, self cannot be task owner, taskId: {%s}", replayScheduleTask.Task.GetTaskId())
		replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(), fmt.Errorf("task ower can not replay schedule task"))
		return
	}

	switch replayScheduleTask.Role {

	// 如果 当前参与方为 DataSupplier   [重新 演算 选 powers]
	case types.DataSupplier:

		powerPartyIds := make([]string, len(replayScheduleTask.Task.GetTaskData().PowerSupplier))
		for i, power := range replayScheduleTask.Task.GetTaskData().PowerSupplier {
			powerPartyIds[i] = power.Organization.PartyId
		}

		dataIdentityIdCache := make(map[string]struct{})
		// 选出 关于自己 metaDataId 所在的 dataNode
		var metaDataId string

		for _, dataSupplier := range replayScheduleTask.Task.GetTaskData().DataSupplier {
			dataIdentityIdCache[dataSupplier.MemberInfo.IdentityId] = struct{}{}

			// 取出 自己的 disk used 信息, identity 和 partyId 都一致, 才是同一个人 ..
			if selfIdentityId == dataSupplier.MemberInfo.IdentityId && replayScheduleTask.PartyId == dataSupplier.MemberInfo.PartyId {
				metaDataId = dataSupplier.MetadataId
			}
		}
		for _, receiver := range replayScheduleTask.Task.GetTaskData().Receivers {
			dataIdentityIdCache[receiver.Receiver.IdentityId] = struct{}{}
		}
		// mock election power orgs
		powers, err := sche.electionConputeOrg(powerPartyIds, dataIdentityIdCache, cost)
		if nil != err {
			log.Errorf("Failed to election powers org on replaySchedule task, taskId: {%s}, err: {%s}", replayScheduleTask.Task.GetTaskId(), err)
			replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(), fmt.Errorf("failed to election power org on replay schedule task, %s", err))
			return
		}

		log.Debugf("Succeed to election powers org on replaySchedule(), taskId {%s}, powers: %s", replayScheduleTask.Task.GetTaskId(), utilOrgPowerArrString(powers))

		// compare powerSuppliers of task And powerSuppliers of election
		if len(powers) != len(replayScheduleTask.Task.GetTaskData().PowerSupplier) {
			log.Errorf("reschedule powers len and task powers len is not match on replay schedule task, taskId: {%s}, reschedule power len: {%d}, task powers len: {%d}",
				replayScheduleTask.Task.GetTaskId(), len(powers), len(replayScheduleTask.Task.GetTaskData().PowerSupplier))
			replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(),
				fmt.Errorf("reschedule powers len and task powers len is not match on replay schedule task"))
			return
		}

		tmp := make(map[string]struct{}, len(powers))

		for _, power := range powers {
			tmp[power.Organization.IdentityId] = struct{}{}
		}
		for _, power := range replayScheduleTask.Task.GetTaskData().PowerSupplier {
			if _, ok := tmp[power.Organization.IdentityId]; !ok {
				log.Errorf("task power identityId not found on reschedule powers on replay schedule task, taskId: {%s}, task power identityId: {%s}",
					replayScheduleTask.Task.GetTaskId(), power.Organization.IdentityId)
				replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(),
					fmt.Errorf("task power identityId not found on reschedule powers on replay schedule task"))
				return
			}
		}

		// 获取 metaData 所在的dataNode 资源
		dataResourceDiskUsed, err := sche.resourceMng.GetDB().QueryDataResourceDiskUsed(metaDataId)
		if nil != err {
			log.Errorf("failed query internal data node by metaDataId, taskId: {%s}, metaDataId: {%s}", replayScheduleTask.Task.GetTaskId(), metaDataId)
			replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(),
				fmt.Errorf("failed query internal data node by metaDataId on replay schedule task"))
			return
		}
		dataNode, err := sche.resourceMng.GetDB().GetRegisterNode(pb.PrefixTypeDataNode, dataResourceDiskUsed.GetNodeId())
		if nil != err {
			log.Errorf("failed query internal data node by metaDataId, taskId: {%s}, metaDataId: {%s}", replayScheduleTask.Task.GetTaskId(), metaDataId)
			replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(),
				fmt.Errorf("failed query internal data node by metaDataId on replay schedule task"))
			return
		}

		log.Debugf("Succeed dataSupplier dataNode on replaySchedule(), taskId: {%s}, dataNode: %s", replayScheduleTask.Task.GetTaskId(), dataNode.String())

		replayScheduleTask.SendResult(&types.ScheduleResult{
			TaskId: replayScheduleTask.Task.GetTaskId(),
			Status: types.TaskSchedOk,
			Resource: &types.PrepareVoteResource{
				Id:      dataNode.Id,
				Ip:      dataNode.ExternalIp,
				Port:    dataNode.ExternalPort,
				PartyId: replayScheduleTask.PartyId,
			},
		})

	// 如果 当前参与方为 PowerSupplier  [选出自己的 内部 power 资源, 并锁定, todo 在最后 DoneXxxxWrap 中解锁]
	case types.PowerSupplier:
		needSlotCount := sche.resourceMng.GetSlotUnit().CalculateSlotCount(cost.Mem, cost.Bandwidth, cost.Processor)
		jobNode, err := sche.electionComputeNode(needSlotCount)
		if nil != err {
			log.Errorf("Failed to election internal power resource, taskId: {%s}, err: {%s}", replayScheduleTask.Task.GetTaskId(), err)
			replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(),
				fmt.Errorf("failed to replay sched myself local power on replay schedule task"))
			return
		}

		log.Debugf("Succeed powerSupplier jobNode on replaySchedule(), taskId: {%s}, jobNode: %s", replayScheduleTask.Task.GetTaskId(), jobNode.String())

		if err := sche.resourceMng.LockLocalResourceWithTask(jobNode.Id, needSlotCount,
			replayScheduleTask.Task); nil != err {
			log.Errorf("Failed to Lock LocalResource {%s} With Task {%s}, err: {%s}",
				jobNode.Id, replayScheduleTask.Task.GetTaskId(), err)
			replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(),
				fmt.Errorf("failed to lock localresource, {%s}", err))
			return
		}

		replayScheduleTask.SendResult(&types.ScheduleResult{
			TaskId: replayScheduleTask.Task.GetTaskId(),
			Status: types.TaskSchedOk,
			Resource: &types.PrepareVoteResource{
				Id:      jobNode.Id,
				Ip:      jobNode.ExternalIp,
				Port:    jobNode.ExternalPort,
				PartyId: replayScheduleTask.PartyId,
			},
		})

	// 如果 当前参与方为 ResultSupplier  [仅仅是选出自己可用的 dataNode]
	case types.ResultSupplier:

		// TODO 判断 task 中对应自己的 privors  是否符合 自己预期 (如: 是否和  powerSuppliers 一致?? 一期 先不做校验了 ...)

		dataResourceTables, err := sche.resourceMng.GetDB().QueryDataResourceTables()
		if nil != err {
			log.Errorf("Failed to election internal data resource with replay schedule task, taskId: {%s}, err: {%s}",
				replayScheduleTask.Task.GetTaskId(), err)
			replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(),
				fmt.Errorf("failed to election internal data resource with replay schedule task, %s", err))
			return
		}

		log.Debugf("QueryDataResourceTables on replaySchedule by taskRole is the resuler, dataResourceTables: %s", utilDataResourceArrString(dataResourceTables))

		resource := dataResourceTables[len(dataResourceTables)-1]
		resourceInfo, err := sche.resourceMng.GetDB().GetRegisterNode(pb.PrefixTypeDataNode, resource.GetNodeId())
		if nil != err {
			log.Errorf("Failed to query internal data node resource,taskId: {%s}, dataNodeId: {%s}, err: {%s}",
				replayScheduleTask.Task.GetTaskId(), resource.GetNodeId(), err)
			replayScheduleTask.SendFailedResult(replayScheduleTask.Task.GetTaskId(),
				fmt.Errorf("failed to query internal data node resource, %s", err))
			return
		}

		log.Debugf("Succeed resultReceiver dataNode on replaySchedule(), taskId: {%s}, dataNode: %s", replayScheduleTask.Task.GetTaskId(), resourceInfo.String())

		replayScheduleTask.SendResult(&types.ScheduleResult{
			TaskId: replayScheduleTask.Task.GetTaskId(),
			Status: types.TaskSchedOk,
			Resource: &types.PrepareVoteResource{
				Id:      resourceInfo.Id,
				Ip:      resourceInfo.ExternalIp,
				Port:    resourceInfo.ExternalPort,
				PartyId: replayScheduleTask.PartyId,
			},
		})
	}
	return
}

func (sche *SchedulerStarveFIFO) increaseTaskTerm() {
	// handle starve queue
	sche.starveQueue.IncreaseTerm()

	// handle queue
	i := 0
	for {
		if i == sche.queue.Len() {
			return
		}
		bullet := (*(sche.queue))[i]
		bullet.IncreaseTerm()

		// When the task in the queue meets hunger, it will be transferred to starveQueue
		if bullet.Term >= StarveTerm {
			bullet.Starve = true
			heap.Push(sche.starveQueue, bullet)
			heap.Remove(sche.queue, i)
			i = 0
			continue
		}
		(*(sche.queue))[i] = bullet
		i++
	}
}

func (sche *SchedulerStarveFIFO) electionComputeNode(needSlotCount uint64) (*pb.YarnRegisteredPeerDetail, error) {

	if nil == sche.internalNodeSet || 0 == sche.internalNodeSet.JobNodeClientSize() {
		return nil, errors.New("not found alive jobNode")
	}

	resourceNodeIdArr := make([]string, 0)

	tables, err := sche.resourceMng.GetLocalResourceTables()
	if nil != err {
		return nil, err
	}
	log.Debugf("GetLocalResourceTables on electionConputeNode, localResources: %s", utilLocalResourceArrString(tables))
	for _, r := range tables {
		if r.IsEnough(uint32(needSlotCount)) {

			jobNodeClient, find := sche.internalNodeSet.QueryJobNodeClient(r.GetNodeId())
			if find && jobNodeClient.IsConnected() {
				resourceNodeIdArr = append(resourceNodeIdArr, r.GetNodeId())
			}
		}
	}

	if len(resourceNodeIdArr) == 0 {
		return nil, ErrEnoughInternalResourceCount
	}

	resourceId := resourceNodeIdArr[len(resourceNodeIdArr)%electionLocalSeed]
	jobNode, err := sche.resourceMng.GetDB().GetRegisterNode(pb.PrefixTypeJobNode, resourceId)
	if nil != err {
		return nil, err
	}
	if nil == jobNode {
		return nil, errors.New("not found jobNode information")
	}
	return jobNode, nil
}

func (sche *SchedulerStarveFIFO) electionConputeOrg(
	powerPartyIds []string,
	dataIdentityIdCache map[string]struct{},
	cost *towTypes.TaskOperationCost,
) ([]*libTypes.TaskPowerSupplier, error) {

	calculateCount := len(powerPartyIds)
	identityIds := make([]string, 0)

	remoteResources := sche.resourceMng.GetRemoteResourceTables()
	log.Debugf("GetRemoteResouceTables on electionConputeOrg, remoteResources: %s", utilRemoteResourceArrString(remoteResources))
	for _, r := range remoteResources {

		// Skip the mock identityId
		if sche.resourceMng.IsMockIdentityId(r.GetIdentityId()) {
			continue
		}

		// 计算方不可以是任务发起方 和 数据参与方 和 接收方
		if _, ok := dataIdentityIdCache[r.GetIdentityId()]; ok {
			continue
		}
		// 还需要有足够的 资源
		if r.IsEnough(cost.Mem,cost.Bandwidth,  cost.Processor) {
			identityIds = append(identityIds, r.GetIdentityId())
		}
	}

	if calculateCount > len(identityIds) {
		return nil, ErrEnoughResourceOrgCountLessCalculateCount
	}

	// Election
	index := electionOrgCondition % len(identityIds)
	identityIdTmp := make(map[string]struct{}, calculateCount)
	for i := calculateCount; i > 0; i-- {
		identityIdTmp[identityIds[index]] = struct{}{}
		index++

	}

	if len(identityIdTmp) != calculateCount {
		return nil, ErrEnoughResourceOrgCountLessCalculateCount
	}

	identityInfoArr, err := sche.resourceMng.GetDB().GetIdentityList()
	if nil != err {
		return nil, err
	}

	log.Debugf("GetIdentityList by dataCenter on electionConputeOrg, identityList: %s", identityInfoArr.String())
	identityInfoTmp := make(map[string]*types.Identity, calculateCount)
	for _, identityInfo := range identityInfoArr {

		// Skip the mock identityId
		if sche.resourceMng.IsMockIdentityId(identityInfo.IdentityId()) {
			continue
		}

		if _, ok := identityIdTmp[identityInfo.IdentityId()]; ok {
			identityInfoTmp[identityInfo.IdentityId()] = identityInfo
		}
	}

	if len(identityInfoTmp) != calculateCount {
		return nil, ErrEnoughResourceOrgCountLessCalculateCount
	}

	resourceArr, err := sche.resourceMng.GetDB().GetResourceList()
	if nil != err {
		return nil, err
	}

	log.Debugf("GetResourceList by dataCenter on electionConputeOrg, resources: %s", resourceArr.String())

	orgs := make([]*libTypes.TaskPowerSupplier, calculateCount)
	i := 0
	for _, iden := range resourceArr {

		if i == calculateCount {
			break
		}

		if info, ok := identityInfoTmp[iden.GetIdentityId()]; ok {
			orgs[i] = &libTypes.TaskPowerSupplier{
				Organization: &apipb.TaskOrganization {
					PartyId:  powerPartyIds[i],
					NodeName: info.Name(),
					NodeId:   info.NodeId(),
					IdentityId: info.IdentityId(),
				},
				// TODO 这里的 task 资源消耗是事先加上的 先在这里直接加上 写死的(任务定义的)
				ResourceUsedOverview: &libTypes.ResourceUsageOverview{
					TotalMem:       iden.GetTotalMem(),
					UsedMem:        cost.Mem,
					TotalProcessor: uint32(iden.GetTotalProcessor()),
					UsedProcessor:  uint32(cost.Processor),
					TotalBandwidth: iden.GetTotalBandWidth(),
					UsedBandwidth:  cost.Bandwidth,
				},
			}
			i++
			delete(identityInfoTmp, iden.GetIdentityId())
		}
	}

	return orgs, nil
}

func (sche *SchedulerStarveFIFO) SendTaskToConsensus(task *types.ConsensusTaskWrap) {
	sche.needConsensusTaskCh <- task
}

func (sche *SchedulerStarveFIFO) SendTaskToTaskManager(task *types.DoneScheduleTaskChWrap) {
	sche.doneSchedTaskCh <- task
}

func utilOrgPowerArrString(powers []*libTypes.TaskPowerSupplier) string {
	arr := make([]string, len(powers))
	for i, power := range powers {
		arr[i] = power.String()
	}
	if len(arr) != 0 {
		return "[" + strings.Join(arr, ",") + "]"
	}
	return "[]"
}
func utilLocalResourceArrString(resources []*types.LocalResourceTable) string {
	arr := make([]string, len(resources))
	for i, r := range resources {
		arr[i] = r.String()
	}
	if len(arr) != 0 {
		return "[" + strings.Join(arr, ",") + "]"
	}
	return "[]"
}

func utilRemoteResourceArrString(resources []*types.RemoteResourceTable) string {
	arr := make([]string, len(resources))
	for i, r := range resources {
		arr[i] = r.String()
	}
	if len(arr) != 0 {
		return "[" + strings.Join(arr, ",") + "]"
	}
	return "[]"
}

func utilDataResourceArrString(resources []*types.DataResourceTable) string {
	arr := make([]string, len(resources))
	for i, r := range resources {
		arr[i] = r.String()
	}
	if len(arr) != 0 {
		return "[" + strings.Join(arr, ",") + "]"
	}
	return "[]"
}