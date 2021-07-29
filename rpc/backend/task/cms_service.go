package task

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/RosettaFlow/Carrier-Go/lib/api"
	libTypes "github.com/RosettaFlow/Carrier-Go/lib/types"
	"github.com/RosettaFlow/Carrier-Go/rpc/backend"
	"github.com/RosettaFlow/Carrier-Go/types"
)

//func (svr *TaskServiceServer) GetTaskSummaryList(ctx context.Context, req *pb.EmptyGetParams) (*pb.GetTaskSummaryListResponse, error) {
//	return nil, nil
//}
//func (svr *TaskServiceServer) GetTaskJoinSummaryList(ctx context.Context, req *pb.GetTaskJoinSummaryListRequest) (*pb.GetTaskJoinSummaryListResponse, error) {
//	return nil, nil
//}
//func (svr *TaskServiceServer) GetTaskDetail(ctx context.Context, req *pb.GetTaskDetailRequest) (*pb.GetTaskDetailResponse, error) {
//	return nil, nil
//}
func (svr *TaskServiceServer) GetTaskDetailList(ctx context.Context, req *pb.EmptyGetParams) (*pb.GetTaskDetailListResponse, error) {
	tasks, err := svr.B.GetTaskDetailList()
	if nil != err {
		log.WithError(err).Error("RPC-API:GetTaskDetailList failed")
		return nil, ErrGetNodeTaskList
	}
	arr := make([]*pb.GetTaskDetailResponse, len(tasks))
	for i, task := range tasks {
		t := &pb.GetTaskDetailResponse{
			Information: types.ConvertTaskDetailShowToPB(task),
			Role: task.Role,
		}
		arr[i] = t
	}
	log.Debugf("RPC-API:GetTaskDetailList succeed, taskList len: {%d}", len(arr))
	return &pb.GetTaskDetailListResponse{
		Status:   0,
		Msg:      backend.OK,
		TaskList: arr,
	}, nil
}

func (svr *TaskServiceServer) GetTaskEventList(ctx context.Context, req *pb.GetTaskEventListRequest) (*pb.GetTaskEventListResponse, error) {

	events, err := svr.B.GetTaskEventList(req.TaskId)
	if nil != err {
		log.WithError(err).Errorf("RPC-API:GetTaskEventList failed, taskId: {%s}", req.TaskId)
		return nil, ErrGetNodeTaskEventList
	}
	log.Debugf("RPC-API:GetTaskEventList succeed, taskId: {%s},  eventList len: {%d}", req.TaskId, len(events))
	return &pb.GetTaskEventListResponse{
		Status:        0,
		Msg:           backend.OK,
		TaskEventList: types.ConvertTaskEventArrToPB(events),
	}, nil
}


func (svr *TaskServiceServer) GetTaskEventListByTaskIds (ctx context.Context, req *pb.GetTaskEventListByTaskIdsRequest) (*pb.GetTaskEventListResponse, error) {

	events, err := svr.B.GetTaskEventListByTaskIds(req.TaskIds)
	if nil != err {
		log.WithError(err).Errorf("RPC-API:GetTaskEventListByTaskIds failed, taskId: {%v}", req.TaskIds)
		return nil, ErrGetNodeTaskEventList
	}
	log.Debugf("RPC-API:GetTaskEventListByTaskIds succeed, taskId: {%v},  eventList len: {%d}", req.TaskIds, len(events))
	return &pb.GetTaskEventListResponse{
		Status:        0,
		Msg:           backend.OK,
		TaskEventList: types.ConvertTaskEventArrToPB(events),
	}, nil
}

func (svr *TaskServiceServer) PublishTaskDeclare(ctx context.Context, req *pb.PublishTaskDeclareRequest) (*pb.PublishTaskDeclareResponse, error) {
	if req == nil || req.Owner == nil {
		return nil, errors.New("required owner")
	}
	if req.OperationCost == nil {
		return nil, errors.New("required operationCost")
	}
	if len(req.Receivers) == 0 {
		return nil, errors.New("required receivers")
	}
	if len( req.DataSupplier) == 0 {
		return nil, errors.New("required partners")
	}

	_, err := svr.B.GetNodeIdentity()
	if nil != err {
		log.WithError(err).Errorf("RPC-API:PublishTaskDeclare failed, query local identity failed, can not publish task")
		return nil, ErrSendTaskMsg
	}


	taskMsg := types.NewTaskMessageFromRequest(req)

	// add  dataSuppliers
	dataSuppliers := make([]*libTypes.TaskMetadataSupplierData, len(req.DataSupplier))
	for i, v := range req.DataSupplier {

		metaData, err := svr.B.GetMetaDataDetail(v.MemberInfo.IdentityId, v.MetaDataInfo.MetaDataId)
		if nil != err {
			log.WithError(err).Errorf("RPC-API:PublishTaskDeclare failed, query metadata of partner failed, identityId: {%s}, metadataId: {%s}",
				v.MemberInfo.IdentityId, v.MetaDataInfo.MetaDataId)
			return nil, fmt.Errorf("failed to query metadata of partner, identityId: {%s}, metadataId: {%s}",
				v.MemberInfo.IdentityId, v.MetaDataInfo.MetaDataId)
		}

		columnArr := make([]*libTypes.ColumnMeta, len(v.MetaDataInfo.ColumnIndexList))
		for j, index := range v.MetaDataInfo.ColumnIndexList {
			col := metaData.MetaData.ColumnMetas[index]
			columnArr[j] = &libTypes.ColumnMeta{
				Cindex:   col.Cindex,
				Ctype:    col.Ctype,
				Cname:    col.Cname,
				Csize:    col.Csize,
				Ccomment: col.Ccomment,
			}
		}

		dataSuppliers[i] = &libTypes.TaskMetadataSupplierData{
			Organization: &libTypes.OrganizationData{
				PartyId:  v.MemberInfo.PartyId,
				NodeName: v.MemberInfo.Name,
				NodeId:   v.MemberInfo.NodeId,
				Identity: v.MemberInfo.IdentityId,
			},
			MetaId:     v.MetaDataInfo.MetaDataId,
			MetaName:   metaData.MetaData.MetaDataSummary.TableName,
			ColumnList: columnArr,
		}
	}

	//// TODO mock dataSuppliers
	//dataSuppliers := make([]*libTypes.TaskMetadataSupplierData, 0)

	taskMsg.Data.SetMetadataSupplierArr(dataSuppliers)

	// add receivers
	receivers := make([]*libTypes.TaskResultReceiverData, len(req.Receivers))
	for i, v := range req.Receivers {

		providers := make([]*libTypes.OrganizationData, len(v.Providers))
		for j, val := range v.Providers {
			providers[j] = &libTypes.OrganizationData{
				PartyId:  val.PartyId,
				NodeName: val.Name,
				NodeId:   val.NodeId,
				Identity: val.IdentityId,
			}
		}

		receivers[i] = &libTypes.TaskResultReceiverData{
			Receiver: &libTypes.OrganizationData{
				PartyId:  v.MemberInfo.PartyId,
				NodeName: v.MemberInfo.Name,
				NodeId:   v.MemberInfo.NodeId,
				Identity: v.MemberInfo.IdentityId,
			},
			Provider: providers,
		}
	}
	taskMsg.Data.SetReceivers(receivers)

	// add empty powerSuppliers
	taskMsg.Data.TaskData().ResourceSupplier = make([]*libTypes.TaskResourceSupplierData, 0)


	// add taskId
	taskId := taskMsg.SetTaskId()
	taskMsg.Data.TaskData().TaskId = taskId

	err = svr.B.SendMsg(taskMsg)
	if nil != err {
		log.WithError(err).Errorf("RPC-API:PublishTaskDeclare failed, query metadata of partner failed, taskId: {%s}",
			taskId)
		return nil, ErrSendTaskMsg
	}
	log.Debugf("RPC-API:PublishTaskDeclare succeed, taskId: {%s}", taskId)
	return &pb.PublishTaskDeclareResponse{
		Status: 0,
		Msg:    backend.OK,
		TaskId: taskId,
	}, nil
}
