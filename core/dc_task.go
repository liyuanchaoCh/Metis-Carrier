package core

import (
	"errors"
	"fmt"
	rawdb "github.com/datumtechs/datum-network-carrier/carrierdb/rawdb"
	carriertypespb "github.com/datumtechs/datum-network-carrier/pb/carrier/types"
	datacenterapipb "github.com/datumtechs/datum-network-carrier/pb/datacenter/api"
	"github.com/datumtechs/datum-network-carrier/types"
	"strings"
)

// about task on local
// local task
func (dc *DataCenter) StoreLocalTask(task *types.Task) error {
	if task == nil {
		return errors.New("invalid params for task")
	}
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreLocalTask(dc.db, task)
}

func (dc *DataCenter) RemoveLocalTask(taskId string) error {
	if taskId == "" {
		return errors.New("invalid params for taskId to DelLocalTask")
	}
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveLocalTask(dc.db, taskId)
}

func (dc *DataCenter) HasLocalTask(taskId string) (bool, error) {
	if taskId == "" {
		return false, nil
	}
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	//log.Debugf("QueryLocalTask, taskId: {%s}", taskId)
	_, err := rawdb.QueryLocalTask(dc.db, taskId)
	if rawdb.IsNoDBNotFoundErr(err) {
		return false, err
	}

	if rawdb.IsDBNotFoundErr(err) {
		return false, nil
	}
	return true, nil
}

func (dc *DataCenter) QueryLocalTask(taskId string) (*types.Task, error) {
	if taskId == "" {
		return nil, errors.New("invalid params taskId for QueryLocalTask")
	}
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	//log.Debugf("QueryLocalTask, taskId: {%s}", taskId)
	return rawdb.QueryLocalTask(dc.db, taskId)
}

func (dc *DataCenter) QueryLocalTaskListByIds(taskIds []string) (types.TaskDataArray, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryLocalTaskByIds(dc.db, taskIds)
}

func (dc *DataCenter) QueryLocalTaskList() (types.TaskDataArray, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryAllLocalTasks(dc.db)
}

func (dc *DataCenter) QueryLocalTaskAndEvents(taskId string) (*types.Task, error) {
	if taskId == "" {
		return nil, errors.New("invalid params taskId for QueryLocalTask")
	}
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	task, err := rawdb.QueryLocalTask(dc.db, taskId)
	if nil != err {
		return nil, err
	}
	list, err := rawdb.QueryTaskEvent(dc.db, taskId)
	if nil != err {
		return nil, err
	}
	task.SetEventList(list)
	return task, nil
}

func (dc *DataCenter) QueryLocalTaskAndEventsListByIds(taskIds []string) (types.TaskDataArray, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	tasks, err := rawdb.QueryLocalTaskByIds(dc.db, taskIds)
	if nil != err {
		return nil, err
	}
	for i, task := range tasks {
		list, err := rawdb.QueryTaskEvent(dc.db, task.GetTaskId())
		if nil != err {
			return nil, err
		}
		task.SetEventList(list)
		tasks[i] = task
	}
	return tasks, nil
}

func (dc *DataCenter) QueryLocalTaskAndEventsList() (types.TaskDataArray, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()

	tasks, err := rawdb.QueryAllLocalTasks(dc.db)
	if nil != err {
		return nil, err
	}
	for i, task := range tasks {
		list, err := rawdb.QueryTaskEvent(dc.db, task.GetTaskId())
		if nil != err {
			return nil, err
		}
		task.SetEventList(list)
		tasks[i] = task
	}
	return tasks, nil
}

func (dc *DataCenter) StoreTaskEvent(event *carriertypespb.TaskEvent) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreTaskEvent(dc.db, event)
}

func (dc *DataCenter) QueryTaskEventList(taskId string) ([]*carriertypespb.TaskEvent, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryTaskEvent(dc.db, taskId)
}

func (dc *DataCenter) QueryTaskEventListByPartyId(taskId, partyId string) ([]*carriertypespb.TaskEvent, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryTaskEventByPartyId(dc.db, taskId, partyId)
}

func (dc *DataCenter) RemoveTaskEventList(taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveTaskEvent(dc.db, taskId)
}

func (dc *DataCenter) RemoveTaskEventListByPartyId(taskId, partyId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveTaskEventByPartyId(dc.db, taskId, partyId)
}

func sprintPowerOrgs(powers []*carriertypespb.TaskOrganization) string {
	arr := make([]string, 0)
	for _, power := range powers {
		arr = append(arr, power.String())
	}
	if len(arr) != 0 {
		return "[" + strings.Join(arr, ",") + "]"
	}
	return "[]"
}

func sprintPowerResources(powers []*carriertypespb.TaskPowerResourceOption) string {
	arr := make([]string, 0)
	for _, power := range powers {
		arr = append(arr, power.String())
	}
	if len(arr) != 0 {
		return "[" + strings.Join(arr, ",") + "]"
	}
	return "[]"
}

// about task on datacenter
func (dc *DataCenter) InsertTask(task *types.Task) error {
	dc.serviceMu.Lock()
	defer dc.serviceMu.Unlock()
	log.Debugf("Start save task to datacenter, taskId: {%s}, powerOrgs: %s, powerResources: %s",
		task.GetTaskId(), sprintPowerOrgs(task.GetTaskData().GetPowerSuppliers()), sprintPowerResources(task.GetTaskData().GetPowerResourceOptions()))
	response, err := dc.client.SaveTask(dc.ctx, types.NewSaveTaskRequest(task))
	if err != nil {
		log.WithError(err).WithField("taskId", task.GetTaskId()).Errorf("InsertTask failed")
		return err
	}
	if response.GetStatus() != 0 {
		return fmt.Errorf("insert task, taskId: {%s},  error: %s", task.GetTaskId(), response.GetMsg())
	}
	return nil
}

func (dc *DataCenter) QueryGlobalTaskList(lastUpdate, pageSize uint64) (types.TaskDataArray, error) {
	dc.serviceMu.RLock()
	defer dc.serviceMu.RUnlock()
	taskListResponse, err := dc.client.ListTask(dc.ctx, &datacenterapipb.ListTaskRequest{
		LastUpdated: lastUpdate,
		PageSize:    pageSize,
	})
	return types.NewTaskArrayFromResponse(taskListResponse), err
}

func (dc *DataCenter) QueryTaskListByIdentityId(identityId string, lastUpdate, pageSize uint64) (types.TaskDataArray, error) {
	dc.serviceMu.RLock()
	defer dc.serviceMu.RUnlock()
	taskListResponse, err := dc.client.ListTaskByIdentity(dc.ctx, &datacenterapipb.ListTaskByIdentityRequest{
		IdentityId:  identityId,
		LastUpdated: lastUpdate,
		PageSize:    pageSize,
	})
	return types.NewTaskArrayFromResponse(taskListResponse), err
}

func (dc *DataCenter) QueryTaskListByTaskIds(taskIds []string) (types.TaskDataArray, error) {
	dc.serviceMu.RLock()
	defer dc.serviceMu.RUnlock()
	taskListResponse, err := dc.client.ListTaskByTaskIds(dc.ctx, &datacenterapipb.ListTaskByTaskIdsRequest{
		TaskIds: taskIds,
	})
	return types.NewTaskArrayFromResponse(taskListResponse), err
}

func (dc *DataCenter) QueryTaskEventListByTaskId(taskId string) ([]*carriertypespb.TaskEvent, error) {
	dc.serviceMu.RLock()
	defer dc.serviceMu.RUnlock()
	taskEventResponse, err := dc.client.ListTaskEvent(dc.ctx, &datacenterapipb.ListTaskEventRequest{
		TaskId: taskId,
	})
	if nil != err {
		return nil, err
	}
	if nil == taskEventResponse {
		return nil, rawdb.ErrNotFound
	}
	return taskEventResponse.GetTaskEvents(), nil
}

func (dc *DataCenter) QueryTaskEventListByTaskIds(taskIds []string) ([]*carriertypespb.TaskEvent, error) {
	dc.serviceMu.RLock()
	defer dc.serviceMu.RUnlock()

	eventList := make([]*carriertypespb.TaskEvent, 0)
	for _, taskId := range taskIds {
		taskEventResponse, err := dc.client.ListTaskEvent(dc.ctx, &datacenterapipb.ListTaskEventRequest{
			TaskId: taskId,
		})
		if nil != err {
			continue
		}
		if nil == taskEventResponse {
			continue
		}
		eventList = append(eventList, taskEventResponse.GetTaskEvents()...)
	}
	return eventList, nil
}

// about TaskPowerUsed
func (dc *DataCenter) StoreLocalTaskPowerUsed(taskPowerUsed *types.LocalTaskPowerUsed) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreLocalTaskPowerUsed(dc.db, taskPowerUsed)
}

func (dc *DataCenter) StoreLocalTaskPowerUseds(taskPowerUseds []*types.LocalTaskPowerUsed) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreLocalTaskPowerUseds(dc.db, taskPowerUseds)
}

func (dc *DataCenter) HasLocalTaskPowerUsed(taskId, partyId string) (bool, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.HasLocalTaskPowerUsed(dc.db, taskId, partyId)
}

func (dc *DataCenter) RemoveLocalTaskPowerUsed(taskId, partyId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveLocalTaskPowerUsed(dc.db, taskId, partyId)
}

func (dc *DataCenter) RemoveLocalTaskPowerUsedByTaskId(taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveLocalTaskPowerUsedByTaskId(dc.db, taskId)
}

func (dc *DataCenter) QueryLocalTaskPowerUsed(taskId, partyId string) (*types.LocalTaskPowerUsed, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryLocalTaskPowerUsed(dc.db, taskId, partyId)
}

func (dc *DataCenter) QueryLocalTaskPowerUsedsByTaskId(taskId string) ([]*types.LocalTaskPowerUsed, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryLocalTaskPowerUsedsByTaskId(dc.db, taskId)
}

func (dc *DataCenter) QueryLocalTaskPowerUseds() ([]*types.LocalTaskPowerUsed, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryLocalTaskPowerUseds(dc.db)
}

// prefix + jobNodeId + taskId -> [partyId, ..., partyId]
func (dc *DataCenter) StoreJobNodeTaskPartyId(jobNodeId, taskId, partyId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreJobNodeTaskPartyId(dc.db, jobNodeId, taskId, partyId)
}

func (dc *DataCenter) RemoveJobNodeTaskPartyId(jobNodeId, taskId, partyId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveJobNodeTaskPartyId(dc.db, jobNodeId, taskId, partyId)
}

func (dc *DataCenter) RemoveJobNodeTaskIdAllPartyIds(jobNodeId, taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveJobNodeTaskIdAllPartyIds(dc.db, jobNodeId, taskId)
}

func (dc *DataCenter) QueryJobNodeRunningTaskIdList(jobNodeId string) ([]string, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryJobNodeRunningTaskIds(dc.db, jobNodeId)
}

func (dc *DataCenter) QueryJobNodeRunningTaskCount(jobNodeId string) (uint32, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryJobNodeRunningTaskIdCount(dc.db, jobNodeId)
}


func (dc *DataCenter) QueryJobNodeRunningTaskIdsAndPartyIdsPairs(jobNodeId string) (map[string][]string, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryJobNodeRunningTaskIdsAndPartyIdsPairs(dc.db, jobNodeId)
}

func (dc *DataCenter) QueryJobNodeRunningTaskAllPartyIdList(jobNodeId, taskId string) ([]string, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryJobNodeTaskAllPartyIds(dc.db, jobNodeId, taskId)
}

func (dc *DataCenter) HasJobNodeRunningTaskId(jobNodeId, taskId string) (bool, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.HasJobNodeRunningTaskId(dc.db, jobNodeId, taskId)
}

func (dc *DataCenter) HasJobNodeTaskPartyId(jobNodeId, taskId, partyId string) (bool, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.HasJobNodeTaskPartyId(dc.db, jobNodeId, taskId, partyId)
}

func (dc *DataCenter) QueryJobNodeTaskPartyIdCount(jobNodeId, taskId string) (uint32, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryJobNodeTaskPartyIdCount(dc.db, jobNodeId, taskId)
}

// about jobNode history taskId and count
func (dc *DataCenter) StoreJobNodeHistoryTaskId(jobNodeId, taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreJobNodeHistoryTaskId(dc.db, jobNodeId, taskId)
}

func (dc *DataCenter) HasJobNodeHistoryTaskId(jobNodeId, taskId string) (bool, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.HasJobNodeHistoryTaskId(dc.db, jobNodeId, taskId)
}

func (dc *DataCenter) QueryJobNodeHistoryTaskCount(jobNodeId string) (uint32, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryJobNodeHistoryTaskCount(dc.db, jobNodeId)
}

// about TaskResultDataMetadataId
func (dc *DataCenter) StoreTaskUpResultData(turf *types.TaskUpResultData) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreTaskUpResultData(dc.db, turf)
}

func (dc *DataCenter) QueryTaskUpResulData(taskId string) (*types.TaskUpResultData, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryTaskUpResultData(dc.db, taskId)
}

func (dc *DataCenter) QueryTaskUpResultDataList() ([]*types.TaskUpResultData, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryTaskUpResultDataList(dc.db)
}

func (dc *DataCenter) RemoveTaskUpResultData(taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveTaskUpResultData(dc.db, taskId)
}

func (dc *DataCenter) StoreTaskPartnerPartyIds(taskId string, partyIds []string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreTaskPartnerPartyIds(dc.db, taskId, partyIds)
}

func (dc *DataCenter) HasTaskPartnerPartyIds(taskId string) (bool, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.HasTaskPartnerPartyIds(dc.db, taskId)
}

func (dc *DataCenter) QueryTaskPartnerPartyIds(taskId string) ([]string, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryTaskPartnerPartyIds(dc.db, taskId)
}

func (dc *DataCenter) RemoveTaskPartnerPartyId(taskId, partyId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	log.Debugf("Start remove partyId of local task's partner arr, taskId: {%s}, partyId: {%s}", taskId, partyId)
	return rawdb.RemoveTaskPartnerPartyId(dc.db, taskId, partyId)
}

func (dc *DataCenter) RemoveTaskPartnerPartyIds(taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveTaskPartnerPartyIds(dc.db, taskId)
}

// v 1.0 -> v 2.0 about task exec status (prefix + taskId + partyId -> "cons"|"exec")
func (dc *DataCenter) StoreLocalTaskExecuteStatusValConsensusByPartyId(taskId, partyId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreLocalTaskExecuteStatusValConsByPartyId(dc.db, taskId, partyId)
}

func (dc *DataCenter) StoreLocalTaskExecuteStatusValRunningByPartyId(taskId, partyId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	log.Debugf("Store task execute status, taskId: {%s}, partyId: {%s}", taskId, partyId)
	return rawdb.StoreLocalTaskExecuteStatusValExecByPartyId(dc.db, taskId, partyId)
}

func (dc *DataCenter) StoreLocalTaskExecuteStatusValTerminateByPartyId(taskId, partyId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	log.Debugf("Store task terminate status, taskId: {%s}, partyId: {%s}", taskId, partyId)
	return rawdb.StoreLocalTaskExecuteStatusValTerminateByPartyId(dc.db, taskId, partyId)
}

func (dc *DataCenter) RemoveLocalTaskExecuteStatusByPartyId(taskId, partyId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	log.Debugf("Remove task execute status, taskId: {%s}, partyId: {%s}", taskId, partyId)
	return rawdb.RemoveLocalTaskExecuteStatusByPartyId(dc.db, taskId, partyId)
}

func (dc *DataCenter) HasLocalTaskExecuteStatusParty(taskId string) (bool, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.HasLocalTaskExecuteStatusParty(dc.db, taskId)
}

func (dc *DataCenter) HasLocalTaskExecuteStatusByPartyId(taskId, partyId string) (bool, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.HasLocalTaskExecuteStatusByPartyId(dc.db, taskId, partyId)
}

func (dc *DataCenter) HasLocalTaskExecuteStatusConsensusByPartyId(taskId, partyId string) (bool, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.HasLocalTaskExecuteStatusConsensusByPartyId(dc.db, taskId, partyId)
}

func (dc *DataCenter) HasLocalTaskExecuteStatusRunningByPartyId(taskId, partyId string) (bool, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.HasLocalTaskExecuteStatusRunningByPartyId(dc.db, taskId, partyId)
}

func (dc *DataCenter) HasLocalTaskExecuteStatusTerminateByPartyId(taskId, partyId string) (bool, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.HasLocalTaskExecuteStatusTerminateByPartyId(dc.db, taskId, partyId)
}

// v 2.0 about local needExecuteTask
func (dc *DataCenter) StoreNeedExecuteTask(task *types.NeedExecuteTask) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreNeedExecuteTask(dc.db, task)
}

func (dc *DataCenter) RemoveNeedExecuteTaskByPartyId(taskId, partyId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveNeedExecuteTaskByPartyId(dc.db, taskId, partyId)
}

func (dc *DataCenter) RemoveNeedExecuteTask(taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveNeedExecuteTask(dc.db, taskId)
}

func (dc *DataCenter) ForEachNeedExecuteTaskWithPrefix(prifix []byte, f func(key, value []byte) error) error {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.ForEachNeedExecuteTaskWwithPrefix(dc.db, prifix, f)
}

func (dc *DataCenter) ForEachNeedExecuteTask(f func(key, value []byte) error) error {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.ForEachNeedExecuteTask(dc.db, f)
}

// v 2.0 about taskbullet
func (dc *DataCenter) StoreTaskBullet(bullet *types.TaskBullet) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreTaskBullet(dc.db, bullet)
}

func (dc *DataCenter) RemoveTaskBullet(taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveTaskBullet(dc.db, taskId)
}

func (dc *DataCenter) ForEachTaskBullets(f func(key, value []byte) error) error {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.ForEachTaskBullets(dc.db, f)
}
