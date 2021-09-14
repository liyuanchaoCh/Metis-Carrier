package grpclient

import (
	"context"
	"github.com/RosettaFlow/Carrier-Go/lib/center/api"
	apicommonpb "github.com/RosettaFlow/Carrier-Go/lib/common"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

const DefaultGrpcRequestTimeout = 2 * time.Second

// Client defines typed wrapper for the CenterData GRPC API.
type GrpcClient struct {
	c *grpc.ClientConn

	// grpc service
	metadataService api.MetadataServiceClient
	resourceService api.ResourceServiceClient
	identityService api.IdentityServiceClient
	taskService 	api.TaskServiceClient
}

// NewClient creates a client that uses the given GRPC client.
func NewGrpcClient(ctx context.Context, addr string) (*GrpcClient, error) {
	ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)  //TODO 默认连接 dataCenter 的 grpc client 超时时间设为 2 s
	_ = cancel
	conn, err := dialContext(ctx, addr)
	if err != nil {
		return nil, err
	}
	return &GrpcClient{
		c: conn,
		metadataService: api.NewMetadataServiceClient(conn),
		resourceService: api.NewResourceServiceClient(conn),
		identityService: api.NewIdentityServiceClient(conn),
		taskService:     api.NewTaskServiceClient(conn),
	}, nil
}

func (gc *GrpcClient) Close() {
	gc.c.Close()
}

func (gc *GrpcClient) GetClientConn() *grpc.ClientConn {
	return gc.c
}

// MetadataSave saves new metadata to database.
func (gc *GrpcClient) SaveMetadata(ctx context.Context, request *api.SaveMetadataRequest) (*apicommonpb.SimpleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return gc.metadataService.SaveMetadata(ctx, request)
}

func (gc *GrpcClient) GetMetadataSummaryList(ctx context.Context) (*api.ListMetadataSummaryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return gc.metadataService.ListMetadataSummary(ctx, &emptypb.Empty{})
}

func (gc *GrpcClient) GetMetadataList(ctx context.Context, request *api.ListMetadataRequest) (*api.ListMetadataResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 600 * time.Second)
	defer cancel()
	return gc.metadataService.ListMetadata(ctx, request)
}

func (gc *GrpcClient) GetMetadataById(ctx context.Context, request *api.FindMetadataByIdRequest) (*api.FindMetadataByIdResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return gc.metadataService.FindMetadataById(ctx, request)
}

func (gc *GrpcClient) RevokeMetadata(ctx context.Context, request *api.RevokeMetadataRequest) (*apicommonpb.SimpleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return gc.metadataService.RevokeMetadata(ctx, request)
}

// ************************************** Resource module *******************************************************

func (gc *GrpcClient) SaveResource(ctx context.Context, request *api.PublishPowerRequest) (*apicommonpb.SimpleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return gc.resourceService.PublishPower(ctx, request)
}

func (gc *GrpcClient) SyncPower(ctx context.Context, request *api.SyncPowerRequest) (*apicommonpb.SimpleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return gc.resourceService.SyncPower(ctx, request)
}

func (gc *GrpcClient) RevokeResource(ctx context.Context, request *api.RevokePowerRequest) (*apicommonpb.SimpleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return gc.resourceService.RevokePower(ctx, request)
}

func (gc *GrpcClient) GetPowerSummaryByIdentityId(ctx context.Context, request *api.GetPowerSummaryByIdentityRequest) (*api.PowerSummaryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return gc.resourceService.GetPowerSummaryByIdentityId(ctx, request)
}

func (gc *GrpcClient) GetPowerTotalSummaryList(ctx context.Context) (*api.ListPowerSummaryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	return gc.resourceService.ListPowerSummary(ctx, &emptypb.Empty{})
}

func (gc *GrpcClient) GetPowerList(ctx context.Context, request *api.ListPowerRequest) (*api.ListPowerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return gc.resourceService.ListPower(ctx, request)
}

// ************************************** Identity module *******************************************************

func (gc *GrpcClient) SaveIdentity(ctx context.Context, request *api.SaveIdentityRequest) (*apicommonpb.SimpleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultGrpcRequestTimeout)
	defer cancel()
	return gc.identityService.SaveIdentity(ctx, request)
}

func (gc *GrpcClient) RevokeIdentityJoin(ctx context.Context, request *api.RevokeIdentityRequest) (*apicommonpb.SimpleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultGrpcRequestTimeout)
	defer cancel()
	return gc.identityService.RevokeIdentity(ctx, request)
}

func (gc *GrpcClient) GetIdentityList(ctx context.Context, request *api.ListIdentityRequest) (*api.ListIdentityResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return gc.identityService.ListIdentity(ctx, request)
}

// 存储元数据鉴权申请记录
func (gc *GrpcClient) SaveMetadataAuthority(ctx context.Context, request *api.MetadataAuthorityRequest) (*apicommonpb.SimpleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultGrpcRequestTimeout)
	defer cancel()
	return gc.identityService.SaveMetadataAuthority(ctx, request)
}

// 数据授权审核，规则：
// 1、授权后，可以将审核结果绑定到原有申请记录之上
func (gc *GrpcClient) AuditMetadataAuthority(ctx context.Context, request *api.MetadataAuthorityRequest) (*apicommonpb.SimpleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultGrpcRequestTimeout)
	defer cancel()
	return gc.identityService.UpdateMetadataAuthority(ctx, request)
}

// 获取数据授权申请列表
// 规则：参数存在时根据条件获取，参数不存在时全量返回
func (gc *GrpcClient) GetMetadataAuthorityList(ctx context.Context, request *api.ListMetadataAuthorityRequest) (*api.ListMetadataAuthorityResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultGrpcRequestTimeout)
	defer cancel()
	return gc.identityService.ListMetadataAuthority(ctx, request)
}

// ************************************** GetTask module *******************************************************

func (gc *GrpcClient) SaveTask(ctx context.Context, request *api.SaveTaskRequest) (*apicommonpb.SimpleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	return gc.taskService.SaveTask(ctx, request)
}

func (gc *GrpcClient) GetDetailTask(ctx context.Context, request *api.GetTaskDetailRequest) (*api.GetTaskDetailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx,2*time.Second)
	defer cancel()
	return gc.taskService.GetTaskDetail(ctx, request)
}

func (gc *GrpcClient) ListTask(ctx context.Context, request *api.ListTaskRequest) (*api.ListTaskResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return gc.taskService.ListTask(ctx, request)
}

func (gc *GrpcClient) ListTaskByIdentity(ctx context.Context, request *api.ListTaskByIdentityRequest) (*api.ListTaskResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return gc.taskService.ListTaskByIdentity(ctx, request)
}

func (gc *GrpcClient) ListTaskEvent(ctx context.Context, request *api.ListTaskEventRequest) (*api.ListTaskEventResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return gc.taskService.ListTaskEvent(ctx, request)
}