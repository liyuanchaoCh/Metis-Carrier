package backend

type RpcBizErr struct {
	Code int32
	Msg  string
}

func NewRpcBizErr(code int32, msg string) *RpcBizErr { return &RpcBizErr{Code: code, Msg: msg} }

func (e *RpcBizErr) Error() string {
	return e.Msg
}

func (e *RpcBizErr) ErrCode() int32 {
	return e.Code
}

// common
var (
	//OK               = &RpcBizErr{Code: 0, Msg: "OK"}
	ErrSystemPanic   = &RpcBizErr{Code: 00001, Msg: "System error"}
	ErrRequireParams = &RpcBizErr{Code: 00002, Msg: "Require params"}
)

// auth
var (
	ErrApplyIdentityMsg        = &RpcBizErr{Code: 10001, Msg: "ApplyIdentity failed"}
	ErrRevokeIdentityMsg       = &RpcBizErr{Code: 10002, Msg: "RevokeIdentity failed"}
	ErrQueryNodeIdentity       = &RpcBizErr{Code: 10003, Msg: "Query local identity failed"}
	ErrQueryIdentityList       = &RpcBizErr{Code: 10004, Msg: "Query all identity list failed"}
	ErrQueryAuthorityList      = &RpcBizErr{Code: 10005, Msg: "Query all authority list failed"}
	ErrApplyMetadataAuthority  = &RpcBizErr{Code: 10006, Msg: "ApplyMetadataAuthority failed"}
	ErrRevokeMetadataAuthority = &RpcBizErr{Code: 10007, Msg: "RevokeMetadataAuthority failed"}
	ErrAuditMetadataAuth       = &RpcBizErr{Code: 10008, Msg: "AuditMetadataAuth failed"}
)

// metadata
var (
	ErrQueryMetadataDetail         = &RpcBizErr{Code: 11001, Msg: "Query metadata detail failed"}
	ErrQueryMetadataDetailList     = &RpcBizErr{Code: 11002, Msg: "Query metadata detail list failed"}
	ErrPublishMetadataMsg          = &RpcBizErr{Code: 11003, Msg: "PublishMetaData failed"}
	ErrRevokeMetadataMsg           = &RpcBizErr{Code: 11004, Msg: "RevokemetaData failed"}
	ErrQueryMetadataUsedTaskIdList = &RpcBizErr{Code: 11005, Msg: "Query metadata used taskId list failed"}
)

// power
var (
	ErrRevokePowerMsg       = &RpcBizErr{Code: 12001, Msg: "RevokePower failed"}
	ErrQueryGlobalPowerList = &RpcBizErr{Code: 12002, Msg: "Query global power list failed"}
	ErrQueryLocalPowerList  = &RpcBizErr{Code: 12003, Msg: "Query local power list failed"}
	ErrPublishPowerMsg      = &RpcBizErr{Code: 12004, Msg: "PublishPower failed"}
)

// task
var (
	ErrQueryNodeTaskList      = &RpcBizErr{Code: 13001, Msg: "Query all task of current node failed"}
	ErrQueryNodeTaskEventList = &RpcBizErr{Code: 13002, Msg: "Query all event of current node's task failed"}
	ErrPublishTaskMsg         = &RpcBizErr{Code: 13003, Msg: "PublishTask failed"}
	ErrTerminateTaskMsg       = &RpcBizErr{Code: 13004, Msg: "TerminateTask failed, send task terminate msg failed"}
)

// yarn
var (
	ErrQueryRegisteredPeers           = &RpcBizErr{Code: 14001, Msg: "Query all registeredNodes failed"}
	ErrSetSeedNode                    = &RpcBizErr{Code: 14002, Msg: "Set seed node failed"}
	ErrDeleteSeedNode                 = &RpcBizErr{Code: 14003, Msg: "Delete seed node failed"}
	ErrQuerySeedNodeList              = &RpcBizErr{Code: 14004, Msg: "Query seed nodes failed"}
	ErrSetDataNode                    = &RpcBizErr{Code: 14005, Msg: "Set data node failed"}
	ErrDeleteDataNode                 = &RpcBizErr{Code: 14006, Msg: "Delete data node failed"}
	ErrQueryDataNodeList              = &RpcBizErr{Code: 14007, Msg: "Query data nodes failed"}
	ErrSetJobNode                     = &RpcBizErr{Code: 14009, Msg: "Set job node failed"}
	ErrDeleteJobNode                  = &RpcBizErr{Code: 14011, Msg: "Delete job node failed"}
	ErrQueryJobNodeList               = &RpcBizErr{Code: 14010, Msg: "Query job nodes failed"}
	ErrReportTaskEvent                = &RpcBizErr{Code: 14012, Msg: "Report taskEvent failed"}
	ErrReportUpFileSummary            = &RpcBizErr{Code: 14013, Msg: "Report upFileSummary failed"}
	ErrQueryTaskResultFileSummary     = &RpcBizErr{Code: 14016, Msg: "Query taskResultFileSummary failed"}
	ErrQueryNodeInfo                  = &RpcBizErr{Code: 14017, Msg: "Query yarn node information failed"}
	ErrReportTaskResultFileSummary    = &RpcBizErr{Code: 14024, Msg: "Report task result file summary failed"}
	ErrQueryAvailableDataNode         = &RpcBizErr{Code: 14031, Msg: "Query availableDataNode failed"}
	ErrQueryFilePosition              = &RpcBizErr{Code: 14032, Msg: "Query filePosition failed"}
	ErrQueryTaskResultFileSummaryList = &RpcBizErr{Code: 14035, Msg: "Query taskResultFileSummary list failed"}
	ErrReportTaskResourceExpense      = &RpcBizErr{Code: 14036, Msg: "FReport taskResourceExpense failed"}
)
