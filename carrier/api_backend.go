package carrier

import (
	"errors"
	"fmt"
	"github.com/RosettaFlow/Carrier-Go/grpclient"
	libTypes "github.com/RosettaFlow/Carrier-Go/lib/types"
	"github.com/RosettaFlow/Carrier-Go/types"
)

// CarrierAPIBackend implements rpc.Backend for Carrier
type CarrierAPIBackend struct {
	carrier *Service
}

func NewCarrierAPIBackend(carrier *Service) *CarrierAPIBackend {
	return &CarrierAPIBackend{carrier: carrier}
}

func (s *CarrierAPIBackend) SendMsg(msg types.Msg) error {
	return s.carrier.mempool.Add(msg)
}

// system (the yarn node self info)
func (s *CarrierAPIBackend) GetNodeInfo() (*types.YarnNodeInfo, error) {
	jobNodes, err := s.carrier.carrierDB.GetRegisterNodeList(types.PREFIX_TYPE_JOBNODE)
	if nil != err {
		log.Error("Failed to get all job nodes, on GetNodeInfo(), err:", err)
		return nil, err
	}
	dataNodes, err := s.carrier.carrierDB.GetRegisterNodeList(types.PREFIX_TYPE_DATANODE)
	if nil != err {
		log.Error("Failed to get all data nodes, on GetNodeInfo(), err:", err)
		return nil, err
	}
	jobsLen := len(jobNodes)
	datasLen := len(dataNodes)
	length := jobsLen + datasLen
	registerNodes := make([]*types.RegisteredNodeDetail, length)
	if len(jobNodes) != 0 {
		for i, v := range jobNodes {
			n := &types.RegisteredNodeDetail{
				NodeType: types.PREFIX_TYPE_JOBNODE.String(),
			}
			n.RegisteredNodeInfo = &types.RegisteredNodeInfo{
				Id:           v.Id,
				InternalIp:   v.InternalIp,
				InternalPort: v.InternalPort,
				ExternalIp:   v.ExternalIp,
				ExternalPort: v.ExternalPort,
				ConnState:    v.ConnState,
			}
			registerNodes[i] = n
		}
	}
	if len(dataNodes) != 0 {
		for i, v := range dataNodes {
			n := &types.RegisteredNodeDetail{
				NodeType: types.PREFIX_TYPE_DATANODE.String(),
			}
			n.RegisteredNodeInfo = &types.RegisteredNodeInfo{
				Id:           v.Id,
				InternalIp:   v.InternalIp,
				InternalPort: v.InternalPort,
				ExternalIp:   v.ExternalIp,
				ExternalPort: v.ExternalPort,
				ConnState:    v.ConnState,
			}
			registerNodes[jobsLen+i] = n
		}
	}
	name, err := s.carrier.carrierDB.GetYarnName()
	if nil != err {
		log.Error("Failed to get yarn nodeName, on GetNodeInfo(), err:", err)
		return nil, err
	}
	identity, err := s.carrier.carrierDB.GetIdentityId()
	if nil != err {
		log.Error("Failed to get identity, on GetNodeInfo(), err:", err)
		return nil, err
	}
	seedNodes, err := s.carrier.carrierDB.GetSeedNodeList()
	return &types.YarnNodeInfo{
		NodeType:     types.PREFIX_TYPE_YARNNODE.String(),
		NodeId:       "",       //TODO 读配置
		InternalIp:   "",       // TODO 读配置
		ExternalIp:   "",       // TODO 读p2p
		InternalPort: "",       // TODO 读配置
		ExternalPort: "",       //TODO 读p2p
		IdentityType: "",       // TODO 读配置
		IdentityId:   identity, // TODO 读接口
		Name:         name,
		Peers:        registerNodes,
		SeedPeers:    seedNodes,
		State:        "", // TODO 读系统状态
	}, nil
}

func (s *CarrierAPIBackend) GetRegisteredPeers() (*types.YarnRegisteredNodeDetail, error) {
	// all dataNodes on yarnNode
	dataNodes, err := s.carrier.carrierDB.GetRegisterNodeList(types.PREFIX_TYPE_DATANODE)
	if nil != err {
		return nil, err
	}
	// all jobNodes on yarnNode
	jobNodes, err := s.carrier.carrierDB.GetRegisterNodeList(types.PREFIX_TYPE_JOBNODE)
	if nil != err {
		return nil, err
	}
	jns := make([]*types.YarnRegisteredJobNode, len(jobNodes))
	for i, v := range jobNodes {
		n := &types.YarnRegisteredJobNode{
			Id:           v.Id,
			InternalIp:   v.InternalIp,
			ExternalIp:   v.ExternalIp,
			InternalPort: v.InternalPort,
			ExternalPort: v.ExternalPort,
			//ResourceUsage:  &types.ResourceUsage{},
			Duration: 0, // TODO 添加运行时长 ...
		}
		n.Task.Count = s.carrier.carrierDB.GetRunningTaskCountOnJobNode(v.Id)
		n.Task.TaskIds = s.carrier.carrierDB.GetJobNodeRunningTaskIdList(v.Id)
		jns[i] = n
	}
	dns := make([]*types.YarnRegisteredDataNode, len(jobNodes))
	for i, v := range dataNodes {
		n := &types.YarnRegisteredDataNode{
			Id:           v.Id,
			InternalIp:   v.InternalIp,
			ExternalIp:   v.ExternalIp,
			InternalPort: v.InternalPort,
			ExternalPort: v.ExternalPort,
			//ResourceUsage:  &types.ResourceUsage{},
			Duration: 0, // TODO 添加运行时长 ...
		}
		n.Delta.FileCount = 0
		n.Delta.FileTotalSize = 0
		dns[i] = n
	}
	return &types.YarnRegisteredNodeDetail{
		JobNodes:  jns,
		DataNodes: dns,
	}, nil
}

func (s *CarrierAPIBackend) SetSeedNode(seed *types.SeedNodeInfo) (types.NodeConnStatus, error) {
	// current node need to connect with seed node.
	return s.carrier.carrierDB.SetSeedNode(seed)
}

func (s *CarrierAPIBackend) DeleteSeedNode(id string) error {
	return s.carrier.carrierDB.DeleteSeedNode(id)
}

func (s *CarrierAPIBackend) GetSeedNode(id string) (*types.SeedNodeInfo, error) {
	return s.carrier.carrierDB.GetSeedNode(id)
}

func (s *CarrierAPIBackend) GetSeedNodeList() ([]*types.SeedNodeInfo, error) {
	return s.carrier.carrierDB.GetSeedNodeList()
}

func (s *CarrierAPIBackend) SetRegisterNode(typ types.RegisteredNodeType, node *types.RegisteredNodeInfo) (types.NodeConnStatus, error) {
	switch typ {
	case types.PREFIX_TYPE_DATANODE:
	case types.PREFIX_TYPE_JOBNODE:
	default:
		return types.NONCONNECTED, errors.New("invalid nodeType")
	}
	if typ == types.PREFIX_TYPE_JOBNODE {
		client, err := grpclient.NewJobNodeClient(s.carrier.ctx, fmt.Sprint("%s:%s", node.ExternalIp, node.ExternalPort), node.Id)
		if err != nil {
			return types.NONCONNECTED, err
		}
		s.carrier.jobNodes[node.Id] = client
	}
	if typ == types.PREFIX_TYPE_DATANODE {
		client, err := grpclient.NewDataNodeClient(s.carrier.ctx, fmt.Sprint("%s:%s", node.InternalIp, node.InternalPort), node.Id)
		if err != nil {
			return types.NONCONNECTED, err
		}
		s.carrier.dataNodes[node.Id] = client
	}
	_, err := s.carrier.carrierDB.SetRegisterNode(typ, node)
	if err != nil {
		return types.NONCONNECTED, err
	}
	return types.CONNECTED, nil
}

func (s *CarrierAPIBackend) DeleteRegisterNode(typ types.RegisteredNodeType, id string) error {
	if typ == types.PREFIX_TYPE_JOBNODE {
		if client, ok := s.carrier.jobNodes[id]; ok {
			client.Close()
			delete(s.carrier.jobNodes, id)
		}
	}
	if typ == types.PREFIX_TYPE_DATANODE {
		if client, ok := s.carrier.dataNodes[id]; ok {
			client.Close()
			delete(s.carrier.dataNodes, id)
		}
	}
	return s.carrier.carrierDB.DeleteRegisterNode(typ, id)
}

func (s *CarrierAPIBackend) GetRegisterNode(typ types.RegisteredNodeType, id string) (*types.RegisteredNodeInfo, error) {
	return s.carrier.carrierDB.GetRegisterNode(typ, id)
}

func (s *CarrierAPIBackend) GetRegisterNodeList(typ types.RegisteredNodeType) ([]*types.RegisteredNodeInfo, error) {
	return s.carrier.carrierDB.GetRegisterNodeList(typ)
}

func (s *CarrierAPIBackend) SendTaskEvent(event *types.TaskEventInfo) error {
	// return s.carrier.resourceManager.SendTaskEvent(evengine)
	return s.carrier.taskManager.SendTaskEvent(event)
}

// metadata api
func (s *CarrierAPIBackend) GetMetaDataDetail(identityId, metaDataId string) (*types.OrgMetaDataInfo, error) {
	metadata, err := s.carrier.carrierDB.GetMetadataByDataId(metaDataId)
	return types.NewOrgMetaDataInfoFromMetadata(metadata), err
}

func (s *CarrierAPIBackend) GetMetaDataDetailList() ([]*types.OrgMetaDataInfo, error) {
	metadataArray, err := s.carrier.carrierDB.GetMetadataList()
	return types.NewOrgMetaDataInfoArrayFromMetadataArray(metadataArray), err
}

func (s *CarrierAPIBackend) GetMetaDataDetailListByOwner(identityId string) ([]*types.OrgMetaDataInfo, error) {
	return nil, nil
}

// power api
func (s *CarrierAPIBackend) GetPowerTotalDetailList() ([]*types.OrgPowerDetail, error) {
	return nil, nil
}

func (s *CarrierAPIBackend) GetPowerSingleDetailList() ([]*types.NodePowerDetail, error) {
	return nil, nil // TODO 未完成,  需要查自己参与过的任务信息
}

// identity api
func (s *CarrierAPIBackend) ApplyIdentityJoin(identity *types.Identity) error {
	return s.carrier.carrierDB.InsertIdentity(identity)
}

func (s *CarrierAPIBackend) RevokeIdentityJoin(identity *types.Identity) error {
	return s.carrier.carrierDB.RevokeIdentity(identity)
}

func (s *CarrierAPIBackend) GetNodeIdentity() (*types.Identity, error) {
	nodeAlias, err := s.carrier.carrierDB.GetIdentity()
	return types.NewIdentity(&libTypes.IdentityData{
		Identity:             nodeAlias.IdentityId,
		NodeId:               nodeAlias.NodeId,
		NodeName:             nodeAlias.Name,
	}), err
}

func (s *CarrierAPIBackend) GetIdentityList() ([]*types.Identity, error) {
	return s.carrier.carrierDB.GetIdentityList()
}

// task api
func (s *CarrierAPIBackend) GetTaskDetailList() ([]*types.TaskDetailShow, error) {
	taskArray, err := s.carrier.carrierDB.GetTaskList()
	return types.NewTaskDetailShowArrayFromTaskDataArray(taskArray), err
}

func (s *CarrierAPIBackend) GetTaskEventList(taskId string) ([]*types.TaskEvent, error) {
	taskEvent, err := s.carrier.carrierDB.GetTaskEventListByTaskId(taskId)
	return types.NewTaskEventFromAPIEvent(taskEvent), err
}

