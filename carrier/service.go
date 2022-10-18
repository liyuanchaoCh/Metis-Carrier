package carrier

import (
	"context"
	"fmt"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/datum-network-carrier/ach/auth"
	"github.com/datumtechs/datum-network-carrier/ach/tk"
	"github.com/datumtechs/datum-network-carrier/ach/tk/kms"
	"github.com/datumtechs/datum-network-carrier/blacklist"
	"github.com/datumtechs/datum-network-carrier/carrierdb"
	"github.com/datumtechs/datum-network-carrier/common"
	"github.com/datumtechs/datum-network-carrier/common/flags"
	"github.com/datumtechs/datum-network-carrier/consensus/chaincons"
	"github.com/datumtechs/datum-network-carrier/consensus/twopc"
	"github.com/datumtechs/datum-network-carrier/core/election"
	"github.com/datumtechs/datum-network-carrier/core/evengine"
	"github.com/datumtechs/datum-network-carrier/core/message"
	"github.com/datumtechs/datum-network-carrier/core/policy"
	"github.com/datumtechs/datum-network-carrier/core/resource"
	"github.com/datumtechs/datum-network-carrier/core/schedule"
	"github.com/datumtechs/datum-network-carrier/core/task"
	"github.com/datumtechs/datum-network-carrier/core/workflow"
	"github.com/datumtechs/datum-network-carrier/db"
	"github.com/datumtechs/datum-network-carrier/grpclient"
	"github.com/datumtechs/datum-network-carrier/handler"
	"github.com/datumtechs/datum-network-carrier/p2p"
	carrierapipb "github.com/datumtechs/datum-network-carrier/pb/carrier/api"
	"github.com/datumtechs/datum-network-carrier/service/discovery"
	"github.com/datumtechs/datum-network-carrier/types"
	"github.com/datumtechs/did-sdk-go/did"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
	"sync"
)

type Service struct {
	isRunning      bool
	processingLock sync.RWMutex
	config         *Config
	carrierDB      carrierdb.CarrierDB
	ctx            context.Context
	cancel         context.CancelFunc
	mempool        *message.Mempool
	Engines        map[types.ConsensusEngineType]handler.Engine

	// DB interfaces
	dataDb          db.Database
	APIBackend      *CarrierAPIBackend
	DebugAPIBackend *CarrierDebugAPIBackend
	BlackListAPI    *blacklist.IdentityBackListCache
	resourceManager *resource.Manager
	messageManager  *message.MessageHandler
	TaskManager     handler.TaskManager
	WorkflowManager *workflow.Manager
	authManager     *auth.AuthorityManager
	scheduler       schedule.Scheduler
	consulManager   *discovery.ConnectConsul
	runError        error
	// add by v0.4.0
	payAgent     *tk.PayAgent
	didService   *did.DIDService
	policyEngine *policy.PolicyEngine
	// add by v0.5.1
	privateIP      *common.CarrierPrivateIP
	adminIPAddress string
	quit           chan struct{}
}

// NewService creates a new CarrierServer object (including the
// initialisation of the common Carrier object)
func NewService(ctx context.Context, cliCtx *cli.Context, config *Config, mockIdentityIdsFile, consensusStateFile string, privateIP *common.CarrierPrivateIP) (*Service, error) {
	ctx, cancel := context.WithCancel(ctx)
	_ = cancel // govet fix for lost cancel. Cancel is handled in service.Stop()

	nodeIdStr := config.P2P.NodeId()
	// read config from p2p config.
	nodeId, _ := p2p.HexID(nodeIdStr)

	pool := message.NewMempool(&message.MempoolConfig{NodeId: nodeIdStr})
	eventEngine := evengine.NewEventEngine(config.CarrierDB)

	needReplayScheduleTaskCh, needExecuteTaskCh :=
		make(chan *types.NeedReplayScheduleTask, config.TaskManagerConfig.NeedReplayScheduleTaskChanSize),
		make(chan *types.NeedExecuteTask, config.TaskManagerConfig.NeedExecuteTaskChanSize)

	log.Debugf("Get some chan size value from config when carrier NewService, NeedReplayScheduleTaskChanSize: %d, NeedExecuteTaskChanSize: %d",
		config.TaskManagerConfig.NeedReplayScheduleTaskChanSize, config.TaskManagerConfig.NeedExecuteTaskChanSize)

	identityBlackListCache, blackListError := blacklist.NewIdentityBackListCache()
	if blackListError != nil {
		return nil, blackListError
	}
	policyEngine := policy.NewPolicyEngine(config.CarrierDB)
	resourceClientSet := grpclient.NewInternalResourceNodeSet()
	resourceMng := resource.NewResourceManager(config.CarrierDB, resourceClientSet, mockIdentityIdsFile)
	authManager := auth.NewAuthorityManager(config.CarrierDB, policyEngine)
	scheduler := schedule.NewSchedulerStarveFIFO(election.NewVrfElector(config.P2P.PirKey(), resourceMng),
		eventEngine, resourceMng, authManager, policyEngine, identityBlackListCache)

	twopcEngine, err := twopc.New(
		&twopc.Config{
			Option: &twopc.OptionConfig{
				NodePriKey: config.P2P.PirKey(),
				NodeID:     nodeId,
			},
			PeerMsgQueueSize:    1024,
			ConsensusStateFile:  consensusStateFile,
			DefaultConsensusWal: config.DefaultConsensusWal,
			DatabaseCache:       config.DatabaseCache,
			DatabaseHandles:     config.DatabaseHandles,
		},
		resourceMng,
		config.P2P,
		needReplayScheduleTaskCh,
		needExecuteTaskCh,
		identityBlackListCache,
	)

	if nil != err {
		return nil, err
	}

	var kmsConfig *kms.Config
	if cliCtx.IsSet(flags.KMSKeyId.Name) && cliCtx.IsSet(flags.KMSRegionId.Name) && cliCtx.IsSet(flags.KMSAccessKeyId.Name) && cliCtx.IsSet(flags.KMSAccessKeySecret.Name) {
		kmsConfig = &kms.Config{
			KeyId:           cliCtx.String(flags.KMSKeyId.Name),
			RegionId:        cliCtx.String(flags.KMSRegionId.Name),
			AccessKeyId:     cliCtx.String(flags.KMSAccessKeyId.Name),
			AccessKeySecret: cliCtx.String(flags.KMSAccessKeySecret.Name),
		}
	}
	//初始化钱包管理器
	tk.InitWalletManager(config.CarrierDB, kmsConfig)

	log.Infof("success to init organization wallet:%s", tk.WalletManagerInstance().GetAddress())
	var ethContext *chainclient.EthContext

	if cliCtx.IsSet(flags.BlockChain.Name) {
		chainUrl := cliCtx.String(flags.BlockChain.Name)
		chainHrp := ""
		if cliCtx.IsSet(flags.HRP.Name) {
			chainHrp = cliCtx.String(flags.HRP.Name)
		}
		ethContext = chainclient.NewEthClientContext(chainUrl, chainHrp, tk.WalletManagerInstance())
	}
	var payAgentContractProxy = ethcommon.HexToAddress(cliCtx.String(flags.PayAgentContractProxy.Name))
	payAgent := tk.NewPayAgent(ethContext, payAgentContractProxy)

	var didConfig *did.Config
	if cliCtx.IsSet(flags.DidDocumentContractProxy.Name) && cliCtx.IsSet(flags.DidPctContractProxy.Name) && cliCtx.IsSet(flags.DidProposalContractProxy.Name) && cliCtx.IsSet(flags.DidCredentialContractProxy.Name) {
		didConfig = &did.Config{
			DocumentContractProxy:   ethcommon.HexToAddress(cliCtx.String(flags.DidDocumentContractProxy.Name)),
			PctContractProxy:        ethcommon.HexToAddress(cliCtx.String(flags.DidPctContractProxy.Name)),
			ProposalContractProxy:   ethcommon.HexToAddress(cliCtx.String(flags.DidProposalContractProxy.Name)),
			CredentialContractProxy: ethcommon.HexToAddress(cliCtx.String(flags.DidCredentialContractProxy.Name)),
		}
	}
	didService := did.NewDIDService(ethContext, didConfig)
	taskExecuteResultCh := make(chan *carrierapipb.WorkFlowTaskStatus, 30)
	taskToMessageHandlerCh := make(chan *types.TaskMsg, 30)
	workflowManager := workflow.NewWorkflowService(config.CarrierDB, taskExecuteResultCh, taskToMessageHandlerCh)
	taskManager, err := task.NewTaskManager(
		config.P2P.PirKey(),
		config.P2P,
		scheduler,
		twopcEngine,
		eventEngine,
		resourceMng,
		authManager,
		payAgent,
		needReplayScheduleTaskCh,
		needExecuteTaskCh,
		config.TaskManagerConfig,
		policyEngine,
		taskExecuteResultCh,
	)
	if nil != err {
		return nil, err
	}

	s := &Service{
		ctx:             ctx,
		cancel:          cancel,
		config:          config,
		carrierDB:       config.CarrierDB,
		mempool:         pool,
		resourceManager: resourceMng,
		messageManager:  message.NewHandler(pool, resourceMng, taskManager, authManager, workflowManager),
		TaskManager:     taskManager,
		WorkflowManager: workflowManager,
		authManager:     authManager,
		payAgent:        payAgent,
		didService:      didService,
		policyEngine:    policyEngine,
		scheduler:       scheduler,
		consulManager: discovery.NewConsulClient(&discovery.ConsulService{
			ServiceIP:   p2p.IpAddr().String(),
			ServicePort: strconv.Itoa(cliCtx.Int(flags.RPCPort.Name)),
			Tags:        config.DiscoverServiceConfig.DiscoveryServerTags,
			Name:        config.DiscoverServiceConfig.DiscoveryServiceName,
			Id:          config.DiscoverServiceConfig.DiscoveryServiceId,
			Interval:    config.DiscoverServiceConfig.DiscoveryServiceHealthCheckInterval,
			Deregister:  config.DiscoverServiceConfig.DiscoveryServiceHealthCheckDeregister,
		},
			config.DiscoverServiceConfig.DiscoveryServerIP,
			config.DiscoverServiceConfig.DiscoveryServerPort,
		),
		privateIP:      privateIP,
		adminIPAddress: strings.TrimSpace(cliCtx.String(flags.AdminIpAddress.Name)),
		//enableGrpcGateWayPrivateCheck: cliCtx.Bool(flags.EnableGrpcGateWayPrivateCheck.Name),
		quit: make(chan struct{}),
	}

	//s.APIBackend = &CarrierAPIBackend{carrier: s}
	s.APIBackend = NewCarrierAPIBackend(s)
	s.DebugAPIBackend = NewCarrierDebugAPIBackend(twopcEngine, identityBlackListCache)
	s.Engines = make(map[types.ConsensusEngineType]handler.Engine, 0)
	s.Engines[types.TwopcTyp] = twopcEngine
	s.Engines[types.ChainconsTyp] = chaincons.New()
	s.BlackListAPI = identityBlackListCache

	// load stored jobNode and dataNode
	jobNodeList, err := s.carrierDB.QueryRegisterNodeList(carrierapipb.PrefixTypeJobNode)
	if err == nil {
		for _, node := range jobNodeList {
			client, err := grpclient.NewJobNodeClient(ctx, fmt.Sprintf("%s:%s", node.GetInternalIp(), node.GetInternalPort()), node.GetId())
			if err == nil {
				s.resourceManager.StoreJobNodeClient(node.GetId(), client)
			}
		}
	}
	dataNodeList, err := s.carrierDB.QueryRegisterNodeList(carrierapipb.PrefixTypeDataNode)
	if err == nil {
		for _, node := range dataNodeList {
			client, err := grpclient.NewDataNodeClient(ctx, fmt.Sprintf("%s:%s", node.GetInternalIp(), node.GetInternalPort()), node.GetId())
			if err == nil {
				s.resourceManager.StoreDataNodeClient(node.GetId(), client)
			}
		}
	}
	return s, nil
}

func (s *Service) Start() error {

	if nil != s.authManager {
		if err := s.authManager.Start(); nil != err {
			log.WithError(err).Errorf("Failed to start the authManager")
		}
	}

	for typ, engine := range s.Engines {
		if err := engine.Start(); nil != err {
			log.WithError(err).Errorf("Cound not start the consensus engine: %s", typ.String())
		}
	}
	if nil != s.resourceManager {
		if err := s.resourceManager.Start(); nil != err {
			log.WithError(err).Errorf("Failed to start the resourceManager")
		}
	}
	if nil != s.messageManager {
		if err := s.messageManager.Start(); nil != err {
			log.WithError(err).Errorf("Failed to start the messageManager")
		}
	}
	if nil != s.TaskManager {
		if err := s.TaskManager.Start(); nil != err {
			log.WithError(err).Errorf("Failed to start the TaskManager")
		}
	}
	if nil != s.scheduler {
		if err := s.scheduler.Start(); nil != err {
			log.WithError(err).Errorf("Failed to start the schedule")
		}
	}
	if nil != s.WorkflowManager {
		if err := s.WorkflowManager.Start(); nil != err {
			log.WithError(err).Errorf("Failed to start the WorkflowManager")
		}
	}
	if err := s.initServicesWithDiscoveryCenter(); nil != err {
		log.Fatal(err)
	}
	go s.loop()
	return nil
}

func (s *Service) Stop() error {
	if s.cancel != nil {
		defer s.cancel()
	}
	s.carrierDB.Stop()

	for typ, engine := range s.Engines {
		if err := engine.Stop(); nil != err {
			log.WithError(err).Errorf("Cound not close the consensus engine: %s", typ.String())
		}
	}
	if nil != s.resourceManager {
		if err := s.resourceManager.Stop(); nil != err {
			log.WithError(err).Errorf("Failed to stop the resourceManager")
		}
	}
	if nil != s.messageManager {
		if err := s.messageManager.Stop(); nil != err {
			log.WithError(err).Errorf("Failed to stop the messageManager")
		}
	}
	if nil != s.TaskManager {
		if err := s.TaskManager.Stop(); nil != err {
			log.WithError(err).Errorf("Failed to stop the TaskManager")
		}
	}
	if nil != s.scheduler {
		if err := s.scheduler.Stop(); nil != err {
			log.WithError(err).Errorf("Failed to stop the schedule")
		}
	}

	if nil != s.authManager {
		if err := s.authManager.Stop(); nil != err {
			log.WithError(err).Errorf("Failed to stop the authManager")
		}
	}
	if nil != s.consulManager {
		if err := s.consulManager.DeregisterService2DiscoveryCenter(); nil != err {
			log.WithError(err).Errorf("Failed to deregister discover service")
		}
	}
	if nil != s.WorkflowManager {
		if err := s.WorkflowManager.Stop(); nil != err {
			log.WithError(err).Errorf("Failed to stop the WorkflowManager")
		}
	}
	// stop service loop gorutine
	close(s.quit)
	return nil
}

// Status is service health checks. Return nil or error.
func (s *Service) Status() error {
	// Service don't start
	if !s.isRunning {
		return nil
	}
	// get error from run function
	if s.runError != nil {
		return s.runError
	}
	return nil
}
