package twopc

import (
	"fmt"
	"github.com/datumtechs/datum-network-carrier/common"
	ctypes "github.com/datumtechs/datum-network-carrier/consensus/twopc/types"
	carriertwopcpb "github.com/datumtechs/datum-network-carrier/pb/carrier/netmsg/consensus/twopc"
	carriertypespb "github.com/datumtechs/datum-network-carrier/pb/carrier/types"
	commonconstantpb "github.com/datumtechs/datum-network-carrier/pb/common/constant"
	"github.com/datumtechs/datum-network-carrier/types"
	"sync"
	"time"
)

type state struct {
	proposalTaskCache map[string]map[string]*ctypes.ProposalTask // (taskId -> partyId -> task)
	// Proposal being processed (proposalId -> partyId -> proposalState)
	proposalSet map[common.Hash]map[string]*ctypes.OrgProposalState
	// About the voting state of prepareMsg for proposal
	prepareVotes map[common.Hash]*prepareVoteState
	// About the voting state of confirmMsg for proposal
	confirmVotes map[common.Hash]*confirmVoteState
	// cache
	proposalPeerInfoCache map[common.Hash]*carriertwopcpb.ConfirmTaskPeerInfo
	// wal
	wal *walDB
	// v 0.3.0 monitors
	syncProposalStateMonitors *ctypes.SyncProposalStateMonitorQueue

	// v 0.4.0 LRC replica check cache for proposalMsg/votingMsg/confirmMsg/commitMsg
	msgCache *TwopcMsgCache

	proposalTaskLock    sync.RWMutex
	proposalsLock       sync.RWMutex
	prepareVotesLock    sync.RWMutex
	confirmVotesLock    sync.RWMutex
	confirmPeerInfoLock sync.RWMutex
}

func newState(ldb *walDB) (*state, error) {
	cache, err := NewTwopcMsgCache(default2pcMsgCacheSize)
	if nil != err {
		return nil, err
	}
	st := &state{
		proposalTaskCache:         make(map[string]map[string]*ctypes.ProposalTask),
		proposalSet:               make(map[common.Hash]map[string]*ctypes.OrgProposalState, 0),
		prepareVotes:              make(map[common.Hash]*prepareVoteState, 0),
		confirmVotes:              make(map[common.Hash]*confirmVoteState, 0),
		proposalPeerInfoCache:     make(map[common.Hash]*carriertwopcpb.ConfirmTaskPeerInfo, 0),
		syncProposalStateMonitors: ctypes.NewSyncProposalStateMonitorQueue(0),
		wal:                       ldb,
		msgCache:                  cache,
	}
	return st, nil
}
func (s *state) IsEmpty() bool    { return nil == s }
func (s *state) IsNotEmpty() bool { return !s.IsEmpty() }

func (s *state) AddMsg(msg interface{}) bool {

	var key []byte
	switch msg.(type) {
	case *carriertwopcpb.PrepareMsg:
		pure := msg.(*carriertwopcpb.PrepareMsg)
		key = append(twopcPrepareMsgCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)

	case *carriertwopcpb.PrepareVote:
		pure := msg.(*carriertwopcpb.PrepareVote)
		key = append(twopcPrepareVoteCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)

	case *carriertwopcpb.ConfirmMsg:
		pure := msg.(*carriertwopcpb.ConfirmMsg)
		key = append(twopcConfirmMsgCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)

	case *carriertwopcpb.ConfirmVote:
		pure := msg.(*carriertwopcpb.ConfirmVote)
		key = append(twopcConfirmVoteCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)

	case *carriertwopcpb.CommitMsg:
		pure := msg.(*carriertwopcpb.CommitMsg)
		key = append(twopcCommitMsgCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)

	//case *TerminateConsensusMsgWrap:
	//default:
	//	return false
	}

	if len(key) != 0 {
		s.msgCache.Add(string(key), struct {}{})
		return true
	}
	return false
}

func (s *state) ContainMsg(msg interface{}) bool {
	switch msg.(type) {
	case *carriertwopcpb.PrepareMsg:
		pure := msg.(*carriertwopcpb.PrepareMsg)
		key := append(twopcPrepareMsgCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)
		return s.msgCache.Contains(string(key))
	case *carriertwopcpb.PrepareVote:
		pure := msg.(*carriertwopcpb.PrepareVote)
		key := append(twopcPrepareVoteCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)
		return s.msgCache.Contains(string(key))
	case *carriertwopcpb.ConfirmMsg:
		pure := msg.(*carriertwopcpb.ConfirmMsg)
		key := append(twopcConfirmMsgCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)
		return s.msgCache.Contains(string(key))
	case *carriertwopcpb.ConfirmVote:
		pure := msg.(*carriertwopcpb.ConfirmVote)
		key := append(twopcConfirmVoteCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)
		return s.msgCache.Contains(string(key))
	case *carriertwopcpb.CommitMsg:
		pure := msg.(*carriertwopcpb.CommitMsg)
		key := append(twopcCommitMsgCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)
		return s.msgCache.Contains(string(key))
	//case *TerminateConsensusMsgWrap:
	default:
		return false
	}
}

// return: ok, evict
func (s *state) ContainsOrAddMsg(msg interface{}) error {

	var key []byte
	switch msg.(type) {
	case *carriertwopcpb.PrepareMsg:
		pure := msg.(*carriertwopcpb.PrepareMsg)
		key = append(twopcPrepareMsgCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)

	case *carriertwopcpb.PrepareVote:
		pure := msg.(*carriertwopcpb.PrepareVote)
		key = append(twopcPrepareVoteCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)

	case *carriertwopcpb.ConfirmMsg:
		pure := msg.(*carriertwopcpb.ConfirmMsg)
		key = append(twopcConfirmMsgCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)

	case *carriertwopcpb.ConfirmVote:
		pure := msg.(*carriertwopcpb.ConfirmVote)
		key = append(twopcConfirmVoteCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)

	case *carriertwopcpb.CommitMsg:
		pure := msg.(*carriertwopcpb.CommitMsg)
		key = append(twopcCommitMsgCacheKeyPrefix, append(pure.GetMsgOption().GetProposalId(), pure.GetMsgOption().GetSenderPartyId()...)...)

	//case *TerminateConsensusMsgWrap:
	//default:
	//	has, evict = false, false
	}

	if len(key) == 0 {
		return fmt.Errorf("not match msg type")
	}
	if has, _ := s.msgCache.ContainsOrAdd(string(key), struct {}{}); has {
		return fmt.Errorf("key value already exists in lru cache")
	}
	return nil
}

// about proposalTask
func (s *state) HasProposalTaskWithTaskId(taskId string) bool {
	s.proposalTaskLock.RLock()
	defer s.proposalTaskLock.RUnlock()
	_, ok := s.proposalTaskCache[taskId]
	return ok
}
func (s *state) HasNotProposalTaskWithTaskId(taskId string) bool {
	return !s.HasProposalTaskWithTaskId(taskId)
}
func (s *state) HasProposalTaskWithTaskIdAndPartyId(taskId, partyId string) bool {
	s.proposalTaskLock.RLock()
	defer s.proposalTaskLock.RUnlock()
	partyCache, ok := s.proposalTaskCache[taskId]
	if !ok {
		return false
	}
	_, ok = partyCache[partyId]
	if !ok {
		return false
	}
	return true
}
func (s *state) HasNotProposalTaskWithTaskIdAndPartyId(taskId, partyId string) bool {
	return !s.HasProposalTaskWithTaskIdAndPartyId(taskId, partyId)
}
func (s *state) StoreProposalTaskWithPartyId(partyId string, task *ctypes.ProposalTask) {
	s.proposalTaskLock.Lock()
	partyCache, ok := s.proposalTaskCache[task.GetTaskId()]
	if !ok {
		partyCache = make(map[string]*ctypes.ProposalTask, 0)
	}
	partyCache[partyId] = task
	s.proposalTaskCache[task.GetTaskId()] = partyCache
	s.proposalTaskLock.Unlock()
}
func (s *state) StoreProposalTaskWithPartyIdUnsafe(partyId string, task *ctypes.ProposalTask) {

	partyCache, ok := s.proposalTaskCache[task.GetTaskId()]
	if !ok {
		partyCache = make(map[string]*ctypes.ProposalTask, 0)
	}
	partyCache[partyId] = task
	s.proposalTaskCache[task.GetTaskId()] = partyCache

}
func (s *state) RemoveProposalTaskWithTaskId(taskId string) {
	s.proposalTaskLock.Lock()
	delete(s.proposalTaskCache, taskId)
	s.proposalTaskLock.Unlock()
}
func (s *state) RemoveProposalTaskWithTaskIdUnsafe(taskId string) {
	delete(s.proposalTaskCache, taskId)
}
func (s *state) RemoveProposalTaskWithTaskIdAndPartyId(taskId, partyId string) {
	s.proposalTaskLock.Lock()
	partyCache, ok := s.proposalTaskCache[taskId]
	if ok {
		delete(partyCache, partyId)
		if len(partyCache) == 0 {
			delete(s.proposalTaskCache, taskId)
		} else {
			s.proposalTaskCache[taskId] = partyCache
		}
	}
	s.proposalTaskLock.Unlock()
}
func (s *state) RemoveProposalTaskWithTaskIdAndPartyIdUnsafe(taskId, partyId string) {

	partyCache, ok := s.proposalTaskCache[taskId]
	if ok {
		delete(partyCache, partyId)
		if len(partyCache) == 0 {
			delete(s.proposalTaskCache, taskId)
		} else {
			s.proposalTaskCache[taskId] = partyCache
		}
	}

}

func (s *state) QueryProposalTaskWithTaskIdAndPartyId(taskId, partyId string) (*ctypes.ProposalTask, bool) {
	s.proposalTaskLock.RLock()
	defer s.proposalTaskLock.RUnlock()
	partyCache, ok := s.proposalTaskCache[taskId]
	if !ok {
		return nil, false
	}
	task, ok := partyCache[partyId]
	return task, ok
}
func (s *state) MustQueryProposalTaskWithTaskIdAndPartyId(taskId, partyId string) *ctypes.ProposalTask {
	task, _ := s.QueryProposalTaskWithTaskIdAndPartyId(taskId, partyId)
	return task
}

// about orgProposalState
func (s *state) HasOrgProposalWithProposalId(proposalId common.Hash) bool {
	s.proposalsLock.RLock()
	defer s.proposalsLock.RUnlock()
	_, ok := s.proposalSet[proposalId]
	return ok
}
func (s *state) HasNotOrgProposalWithProposalId(proposalId common.Hash) bool {
	return !s.HasOrgProposalWithProposalId(proposalId)
}
func (s *state) HasOrgProposalWithProposalIdAndPartyId(proposalId common.Hash, partyId string) bool {
	s.proposalsLock.RLock()
	defer s.proposalsLock.RUnlock()

	if cache, ok := s.proposalSet[proposalId]; ok {
		if _, has := cache[partyId]; has {
			return true
		}
	}
	return false
}
func (s *state) HasNotOrgProposalWithProposalIdAndPartyId(proposalId common.Hash, partyId string) bool {
	return !s.HasOrgProposalWithProposalIdAndPartyId(proposalId, partyId)
}
func (s *state) StoreOrgProposalState(orgState *ctypes.OrgProposalState) {
	s.proposalsLock.Lock()

	cache, ok := s.proposalSet[orgState.GetProposalId()]
	if !ok {
		cache = make(map[string]*ctypes.OrgProposalState, 0)
	}
	cache[orgState.GetTaskOrg().GetPartyId()] = orgState
	s.proposalSet[orgState.GetProposalId()] = cache

	s.proposalsLock.Unlock()
}
func (s *state) StoreOrgProposalStateUnsafe(orgState *ctypes.OrgProposalState) {

	cache, ok := s.proposalSet[orgState.GetProposalId()]
	if !ok {
		cache = make(map[string]*ctypes.OrgProposalState, 0)
	}
	cache[orgState.GetTaskOrg().GetPartyId()] = orgState
	s.proposalSet[orgState.GetProposalId()] = cache

}
func (s *state) RemoveProposalStateWithProposalId(proposalId common.Hash) {
	s.proposalsLock.Lock()
	delete(s.proposalSet, proposalId)
	s.proposalsLock.Unlock()
}
func (s *state) RemoveProposalStateWithProposalIdUnsafe(proposalId common.Hash) {
	delete(s.proposalSet, proposalId)
}
func (s *state) RemoveOrgProposalStateWithProposalIdAndPartyId(proposalId common.Hash, partyId string) {
	s.proposalsLock.Lock()
	if cache, ok := s.proposalSet[proposalId]; ok {

		delete(cache, partyId)

		if len(cache) == 0 {
			delete(s.proposalSet, proposalId)
		} else {
			s.proposalSet[proposalId] = cache
		}
	}
	s.proposalsLock.Unlock()
}
func (s *state) RemoveOrgProposalStateWithProposalIdAndPartyIdUnsafe(proposalId common.Hash, partyId string) {

	if cache, ok := s.proposalSet[proposalId]; ok {

		delete(cache, partyId)

		if len(cache) == 0 {
			delete(s.proposalSet, proposalId)
		} else {
			s.proposalSet[proposalId] = cache
		}
	}

}
func (s *state) RandomOrgProposalStateWithProposalId(proposalId common.Hash) (*ctypes.OrgProposalState, bool) {
	s.proposalsLock.Lock()
	defer s.proposalsLock.Unlock()

	cache, ok := s.proposalSet[proposalId]
	if !ok {
		return nil, false
	}

	for _, orgState := range cache {
		if nil != orgState {
			return orgState, true
		}
	}
	return nil, false
}
func (s *state) QueryOrgProposalStateWithProposalIdAndPartyId(proposalId common.Hash, partyId string) (*ctypes.OrgProposalState, bool) {
	s.proposalsLock.Lock()
	defer s.proposalsLock.Unlock()

	cache, ok := s.proposalSet[proposalId]
	if !ok {
		return nil, false
	}
	orgState, ok := cache[partyId]
	return orgState, ok
}
func (s *state) MustQueryOrgProposalStateWithProposalIdAndPartyId(proposalId common.Hash, partyId string) *ctypes.OrgProposalState {
	orgState, _ := s.QueryOrgProposalStateWithProposalIdAndPartyId(proposalId, partyId)
	return orgState
}

func (s *state) ChangeToConfirm(proposalId common.Hash, partyId string, startTime uint64) {
	s.proposalsLock.Lock()
	defer s.proposalsLock.Unlock()

	cache, ok := s.proposalSet[proposalId]
	if !ok {
		return
	}
	orgProposalState, ok := cache[partyId]
	if !ok {
		return
	}
	orgProposalState.ChangeToConfirm()
	cache[partyId] = orgProposalState

	s.proposalSet[proposalId] = cache
}

func (s *state) ChangeToCommit(proposalId common.Hash, partyId string, startTime uint64) {
	s.proposalsLock.Lock()
	defer s.proposalsLock.Unlock()

	cache, ok := s.proposalSet[proposalId]
	if !ok {
		return
	}
	orgProposalState, ok := cache[partyId]
	if !ok {
		return
	}
	orgProposalState.ChangeToCommit()
	cache[partyId] = orgProposalState

	s.proposalSet[proposalId] = cache
}

func (s *state) RemoveOrgProposalStateAnyCache(proposalId common.Hash, taskId, partyId string) {
	s.RemoveProposalTaskWithTaskIdAndPartyId(taskId, partyId)
	s.RemoveOrgProposalStateWithProposalIdAndPartyId(proposalId, partyId)
	s.RemoveOrgPrepareVoteState(proposalId, partyId)
	s.RemoveOrgConfirmVoteState(proposalId, partyId)
	// wal
	s.wal.DeleteState(s.wal.GetProposalTaskCacheKey(taskId, partyId))
	s.wal.DeleteState(s.wal.GetProposalSetKey(proposalId, partyId))
	s.wal.DeleteState(s.wal.GetPrepareVotesKey(proposalId, partyId))
	s.wal.DeleteState(s.wal.GetConfirmVotesKey(proposalId, partyId))
	// v 0.3.0 delete proposal state monitor
	s.DelMonitor(proposalId, partyId)
}

// ---------------- PrepareVote ----------------
func (s *state) HasPrepareVoting(proposalId common.Hash, org *carriertypespb.TaskOrganization) bool {
	s.prepareVotesLock.RLock()
	pvs, ok := s.prepareVotes[proposalId]
	s.prepareVotesLock.RUnlock()
	if !ok {
		return false
	}
	return pvs.hasPrepareVoting(org.GetPartyId(), org.GetIdentityId())
}

func (s *state) StorePrepareVote(vote *types.PrepareVote) {
	s.prepareVotesLock.Lock()
	pvs, ok := s.prepareVotes[vote.GetMsgOption().GetProposalId()]
	if !ok {
		pvs = newPrepareVoteState()
	}
	pvs.addVote(vote)
	s.wal.StorePrepareVote(vote)
	s.prepareVotes[vote.GetMsgOption().GetProposalId()] = pvs
	s.prepareVotesLock.Unlock()
}

func (s *state) GetPrepareVoteArr(proposalId common.Hash) []*types.PrepareVote {
	s.prepareVotesLock.RLock()
	pvs, ok := s.prepareVotes[proposalId]
	s.prepareVotesLock.RUnlock()
	if !ok {
		return nil
	}
	return pvs.getVotes()
}

func (s *state) GetPrepareVote(proposalId common.Hash, partyId string) *types.PrepareVote {
	s.prepareVotesLock.RLock()
	pvs, ok := s.prepareVotes[proposalId]
	s.prepareVotesLock.RUnlock()
	if !ok {
		return nil
	}
	return pvs.getVote(partyId)
}

func (s *state) HasPrepareVoteState(proposalId common.Hash, partyId, identityId string) bool {
	s.prepareVotesLock.RLock()
	pvs, ok := s.prepareVotes[proposalId]
	s.prepareVotesLock.RUnlock()
	if ok {
		return pvs.hasPrepareVoting(partyId, identityId)
	}
	return false
}

func (s *state) RemovePrepareVoteState(proposalId common.Hash) {
	s.prepareVotesLock.Lock()
	delete(s.prepareVotes, proposalId)
	s.prepareVotesLock.Unlock()
}

func (s *state) RemoveOrgPrepareVoteState(proposalId common.Hash, partyId string) {
	s.prepareVotesLock.Lock()
	if pvotes, ok := s.prepareVotes[proposalId]; ok {
		if vote := pvotes.getVote(partyId); nil != vote {
			pvotes.removeVote(partyId, vote.GetMsgOption().GetSenderRole())
		}
		if pvotes.isEmpty() {
			delete(s.prepareVotes, proposalId)
		} else {
			s.prepareVotes[proposalId] = pvotes
		}
	}
	s.prepareVotesLock.Unlock()
}

func (s *state) GetTaskPrepareYesVoteCount(proposalId common.Hash) uint32 {
	return s.GetTaskDataSupplierPrepareYesVoteCount(proposalId) +
		s.GetTaskPowerSupplierPrepareYesVoteCount(proposalId) +
		s.GetTaskReceiverPrepareYesVoteCount(proposalId)
}
func (s *state) GetTaskDataSupplierPrepareYesVoteCount(proposalId common.Hash) uint32 {
	s.prepareVotesLock.RLock()
	pvs, ok := s.prepareVotes[proposalId]
	s.prepareVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return pvs.voteYesCount(commonconstantpb.TaskRole_TaskRole_DataSupplier)
}
func (s *state) GetTaskPowerSupplierPrepareYesVoteCount(proposalId common.Hash) uint32 {
	s.prepareVotesLock.RLock()
	pvs, ok := s.prepareVotes[proposalId]
	s.prepareVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return pvs.voteYesCount(commonconstantpb.TaskRole_TaskRole_PowerSupplier)
}
func (s *state) GetTaskReceiverPrepareYesVoteCount(proposalId common.Hash) uint32 {
	s.prepareVotesLock.RLock()
	pvs, ok := s.prepareVotes[proposalId]
	s.prepareVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return pvs.voteYesCount(commonconstantpb.TaskRole_TaskRole_Receiver)
}

func (s *state) GetTaskPrepareTotalVoteCount(proposalId common.Hash) uint32 {
	return s.GetTaskDataSupplierPrepareTotalVoteCount(proposalId) +
		s.GetTaskPowerSupplierPrepareTotalVoteCount(proposalId) +
		s.GetTaskReceiverPrepareTotalVoteCount(proposalId)
}
func (s *state) GetTaskDataSupplierPrepareTotalVoteCount(proposalId common.Hash) uint32 {
	s.prepareVotesLock.RLock()
	pvs, ok := s.prepareVotes[proposalId]
	s.prepareVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return pvs.voteTotalCount(commonconstantpb.TaskRole_TaskRole_DataSupplier)
}
func (s *state) GetTaskPowerSupplierPrepareTotalVoteCount(proposalId common.Hash) uint32 {
	s.prepareVotesLock.RLock()
	pvs, ok := s.prepareVotes[proposalId]
	s.prepareVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return pvs.voteTotalCount(commonconstantpb.TaskRole_TaskRole_PowerSupplier)
}
func (s *state) GetTaskReceiverPrepareTotalVoteCount(proposalId common.Hash) uint32 {
	s.prepareVotesLock.RLock()
	pvs, ok := s.prepareVotes[proposalId]
	s.prepareVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return pvs.voteTotalCount(commonconstantpb.TaskRole_TaskRole_Receiver)
}

// ---------------- ConfirmVote ----------------
func (s *state) HasConfirmVoting(proposalId common.Hash, org *carriertypespb.TaskOrganization) bool {
	s.confirmVotesLock.RLock()
	cvs, ok := s.confirmVotes[proposalId]
	s.confirmVotesLock.RUnlock()
	if !ok {
		return false
	}
	return cvs.hasConfirmVoting(org.GetPartyId(), org.GetIdentityId())
}
func (s *state) StoreConfirmVote(vote *types.ConfirmVote) {
	s.confirmVotesLock.Lock()
	cvs, ok := s.confirmVotes[vote.GetMsgOption().GetProposalId()]
	if !ok {
		cvs = newConfirmVoteState()
	}
	cvs.addVote(vote)
	s.confirmVotes[vote.GetMsgOption().GetProposalId()] = cvs
	s.wal.StoreConfirmVote(vote)
	s.confirmVotesLock.Unlock()
}

func (s *state) GetConfirmVoteArr(proposalId common.Hash) []*types.ConfirmVote {
	s.confirmVotesLock.RLock()
	cvs, ok := s.confirmVotes[proposalId]
	s.confirmVotesLock.RUnlock()
	if !ok {
		return nil
	}
	return cvs.getVotes()
}

func (s *state) GetConfirmVote(proposalId common.Hash, partyId string) *types.ConfirmVote {
	s.confirmVotesLock.RLock()
	cvs, ok := s.confirmVotes[proposalId]
	s.confirmVotesLock.RUnlock()
	if !ok {
		return nil
	}
	return cvs.getVote(partyId)
}

func (s *state) HasConfirmVoteState(proposalId common.Hash, partyId, identityId string) bool {
	s.confirmVotesLock.RLock()
	cvs, ok := s.confirmVotes[proposalId]
	s.confirmVotesLock.RUnlock()
	if ok {
		return cvs.hasConfirmVoting(partyId, identityId)
	}
	return false
}

func (s *state) RemoveConfirmVoteState(proposalId common.Hash) {
	s.confirmVotesLock.Lock()
	delete(s.confirmVotes, proposalId)
	s.confirmVotesLock.Unlock()
}

func (s *state) RemoveOrgConfirmVoteState(proposalId common.Hash, partyId string) {
	s.confirmVotesLock.Lock()
	if cvotes, ok := s.confirmVotes[proposalId]; ok {
		if vote := cvotes.getVote(partyId); nil != vote {
			cvotes.removeVote(partyId, vote.GetMsgOption().GetSenderRole())
		}
		if cvotes.isEmpty() {
			delete(s.confirmVotes, proposalId)
		} else {
			s.confirmVotes[proposalId] = cvotes
		}
	}
	s.confirmVotesLock.Unlock()
}

func (s *state) GetTaskConfirmYesVoteCount(proposalId common.Hash) uint32 {
	return s.GetTaskDataSupplierConfirmYesVoteCount(proposalId) +
		s.GetTaskPowerSupplierConfirmYesVoteCount(proposalId) +
		s.GetTaskReceiverConfirmYesVoteCount(proposalId)
}
func (s *state) GetTaskDataSupplierConfirmYesVoteCount(proposalId common.Hash) uint32 {
	s.confirmVotesLock.RLock()
	cvs, ok := s.confirmVotes[proposalId]
	s.confirmVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return cvs.voteYesCount(commonconstantpb.TaskRole_TaskRole_DataSupplier)
}
func (s *state) GetTaskPowerSupplierConfirmYesVoteCount(proposalId common.Hash) uint32 {
	s.confirmVotesLock.RLock()
	cvs, ok := s.confirmVotes[proposalId]
	s.confirmVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return cvs.voteYesCount(commonconstantpb.TaskRole_TaskRole_PowerSupplier)
}
func (s *state) GetTaskReceiverConfirmYesVoteCount(proposalId common.Hash) uint32 {
	s.confirmVotesLock.RLock()
	cvs, ok := s.confirmVotes[proposalId]
	s.confirmVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return cvs.voteYesCount(commonconstantpb.TaskRole_TaskRole_Receiver)
}
func (s *state) GetTaskConfirmTotalVoteCount(proposalId common.Hash) uint32 {
	return s.GetTaskDataSupplierConfirmTotalVoteCount(proposalId) +
		s.GetTaskPowerSupplierConfirmTotalVoteCount(proposalId) +
		s.GetTaskReceiverConfirmTotalVoteCount(proposalId)
}
func (s *state) GetTaskDataSupplierConfirmTotalVoteCount(proposalId common.Hash) uint32 {
	s.confirmVotesLock.RLock()
	cvs, ok := s.confirmVotes[proposalId]
	s.confirmVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return cvs.voteTotalCount(commonconstantpb.TaskRole_TaskRole_DataSupplier)
}
func (s *state) GetTaskPowerSupplierConfirmTotalVoteCount(proposalId common.Hash) uint32 {
	s.confirmVotesLock.RLock()
	cvs, ok := s.confirmVotes[proposalId]
	s.confirmVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return cvs.voteTotalCount(commonconstantpb.TaskRole_TaskRole_PowerSupplier)
}
func (s *state) GetTaskReceiverConfirmTotalVoteCount(proposalId common.Hash) uint32 {
	s.confirmVotesLock.RLock()
	cvs, ok := s.confirmVotes[proposalId]
	s.confirmVotesLock.RUnlock()
	if !ok {
		return 0
	}
	return cvs.voteTotalCount(commonconstantpb.TaskRole_TaskRole_Receiver)
}

// about prepareVote
type prepareVoteState struct {
	votes      map[string]*types.PrepareVote // partyId -> vote
	yesVotes   map[commonconstantpb.TaskRole]uint32
	voteStatus map[commonconstantpb.TaskRole]uint32 // total vote count
	lock       sync.Mutex
}

func newPrepareVoteState() *prepareVoteState {
	return &prepareVoteState{
		votes:      make(map[string]*types.PrepareVote, 0),
		yesVotes:   make(map[commonconstantpb.TaskRole]uint32, 0),
		voteStatus: make(map[commonconstantpb.TaskRole]uint32, 0),
	}
}

func (st *prepareVoteState) isEmpty() bool {
	if nil == st {
		return true
	}
	return len(st.votes) == 0
}

func (st *prepareVoteState) isNotEmpty() bool {
	return !st.isEmpty()
}

func (st *prepareVoteState) addVote(vote *types.PrepareVote) {
	st.lock.Lock()
	defer st.lock.Unlock()

	if _, ok := st.votes[vote.GetMsgOption().GetSenderPartyId()]; ok {
		return
	}
	st.votes[vote.GetMsgOption().GetSenderPartyId()] = vote
	if count, ok := st.yesVotes[vote.GetMsgOption().GetSenderRole()]; ok {
		if vote.GetVoteOption() == types.YES {
			st.yesVotes[vote.GetMsgOption().GetSenderRole()] = count + 1
		}
	} else {
		if vote.GetVoteOption() == types.YES {
			st.yesVotes[vote.GetMsgOption().GetSenderRole()] = 1
		}
	}

	if count, ok := st.voteStatus[vote.GetMsgOption().GetSenderRole()]; ok {
		st.voteStatus[vote.GetMsgOption().GetSenderRole()] = count + 1
	} else {
		st.voteStatus[vote.GetMsgOption().GetSenderRole()] = 1
	}
}
func (st *prepareVoteState) removeVote(partyId string, role commonconstantpb.TaskRole) {
	st.lock.Lock()
	defer st.lock.Unlock()

	vote, ok := st.votes[partyId]
	if !ok {
		return
	}

	delete(st.votes, partyId)

	if count, ok := st.yesVotes[role]; ok {
		if vote.GetVoteOption() == types.YES && count != 0 {
			st.yesVotes[role] = count - 1
		}
	}

	if count, ok := st.voteStatus[role]; ok {
		if count != 0 {
			st.voteStatus[role] = count - 1
		}
	}
}
func (st *prepareVoteState) isEmptyVote() bool                        { return len(st.votes) == 0 }
func (st *prepareVoteState) isNotEmptyVote() bool                     { return !st.isEmptyVote() }
func (st *prepareVoteState) getVote(partId string) *types.PrepareVote { return st.votes[partId] }
func (st *prepareVoteState) getVotes() []*types.PrepareVote {
	arr := make([]*types.PrepareVote, 0, len(st.votes))
	for _, vote := range st.votes {
		arr = append(arr, vote)
	}
	return arr
}
func (st *prepareVoteState) voteTotalCount(role commonconstantpb.TaskRole) uint32 {
	st.lock.Lock()
	defer st.lock.Unlock()

	if count, ok := st.voteStatus[role]; ok {
		return count
	} else {
		return 0
	}
}
func (st *prepareVoteState) voteYesCount(role commonconstantpb.TaskRole) uint32 {
	st.lock.Lock()
	defer st.lock.Unlock()

	if count, ok := st.yesVotes[role]; ok {
		return count
	} else {
		return 0
	}
}
func (st *prepareVoteState) hasPrepareVoting(partyId, identityId string) bool {
	if vote, ok := st.votes[partyId]; ok {
		if vote.GetMsgOption().GetSenderPartyId() == partyId && vote.GetMsgOption().GetOwner().GetIdentityId() == identityId {
			return true
		}
	}
	return false
}

// about confirmVote
type confirmVoteState struct {
	votes      map[string]*types.ConfirmVote // partyId -> vote
	yesVotes   map[commonconstantpb.TaskRole]uint32
	voteStatus map[commonconstantpb.TaskRole]uint32
	lock       sync.Mutex
}

func newConfirmVoteState() *confirmVoteState {
	return &confirmVoteState{
		votes:      make(map[string]*types.ConfirmVote, 0),
		yesVotes:   make(map[commonconstantpb.TaskRole]uint32, 0),
		voteStatus: make(map[commonconstantpb.TaskRole]uint32, 0),
	}
}

func (st *confirmVoteState) isEmpty() bool {
	if nil == st {
		return true
	}
	return len(st.votes) == 0
}

func (st *confirmVoteState) isNotEmpty() bool {
	return !st.isEmpty()
}

func (st *confirmVoteState) addVote(vote *types.ConfirmVote) {
	st.lock.Lock()
	defer st.lock.Unlock()

	if _, ok := st.votes[vote.GetMsgOption().GetSenderPartyId()]; ok {
		return
	}

	st.votes[vote.GetMsgOption().GetSenderPartyId()] = vote
	if count, ok := st.yesVotes[vote.GetMsgOption().GetSenderRole()]; ok {
		if vote.GetVoteOption() == types.YES {
			st.yesVotes[vote.GetMsgOption().GetSenderRole()] = count + 1
		}
	} else {
		if vote.GetVoteOption() == types.YES {
			st.yesVotes[vote.GetMsgOption().GetSenderRole()] = 1
		}
	}

	if count, ok := st.voteStatus[vote.GetMsgOption().GetSenderRole()]; ok {
		st.voteStatus[vote.GetMsgOption().GetSenderRole()] = count + 1
	} else {
		st.voteStatus[vote.GetMsgOption().GetSenderRole()] = 1
	}
}

func (st *confirmVoteState) removeVote(partyId string, role commonconstantpb.TaskRole) {
	st.lock.Lock()
	defer st.lock.Unlock()

	vote, ok := st.votes[partyId]
	if !ok {
		return
	}

	delete(st.votes, partyId)

	if count, ok := st.yesVotes[role]; ok {
		if vote.GetVoteOption() == types.YES && count != 0 {
			st.yesVotes[role] = count - 1
		}
	}

	if count, ok := st.voteStatus[role]; ok {
		if count != 0 {
			st.voteStatus[role] = count - 1
		}
	}
}
func (st *confirmVoteState) isEmptyVote() bool                        { return len(st.votes) == 0 }
func (st *confirmVoteState) isNotEmptyVote() bool                     { return !st.isEmptyVote() }
func (st *confirmVoteState) getVote(partId string) *types.ConfirmVote { return st.votes[partId] }
func (st *confirmVoteState) getVotes() []*types.ConfirmVote {
	arr := make([]*types.ConfirmVote, 0, len(st.votes))
	for _, vote := range st.votes {
		arr = append(arr, vote)
	}
	return arr
}

func (st *confirmVoteState) voteYesCount(role commonconstantpb.TaskRole) uint32 {
	st.lock.Lock()
	defer st.lock.Unlock()

	if count, ok := st.yesVotes[role]; ok {
		return count
	} else {
		return 0
	}
}
func (st *confirmVoteState) voteTotalCount(role commonconstantpb.TaskRole) uint32 {
	st.lock.Lock()
	defer st.lock.Unlock()

	if count, ok := st.voteStatus[role]; ok {
		return count
	} else {
		return 0
	}
}
func (st *confirmVoteState) hasConfirmVoting(partyId, identityId string) bool {
	if vote, ok := st.votes[partyId]; ok {
		if vote.GetMsgOption().GetSenderPartyId() == partyId && vote.GetMsgOption().GetOwner().GetIdentityId() == identityId {
			return true
		}
	}
	return false
}

func (s *state) StoreConfirmTaskPeerInfo(proposalId common.Hash, peerDesc *carriertwopcpb.ConfirmTaskPeerInfo) {
	s.confirmPeerInfoLock.Lock()
	_, ok := s.proposalPeerInfoCache[proposalId]
	if !ok {
		s.proposalPeerInfoCache[proposalId] = peerDesc
		s.wal.StoreConfirmTaskPeerInfo(proposalId, peerDesc)
	}
	s.confirmPeerInfoLock.Unlock()
}

func (s *state) HasConfirmTaskPeerInfo(proposalId common.Hash) bool {
	s.confirmPeerInfoLock.RLock()
	_, ok := s.proposalPeerInfoCache[proposalId]
	s.confirmPeerInfoLock.RUnlock()
	return ok
}

func (s *state) QueryConfirmTaskPeerInfo(proposalId common.Hash) (*carriertwopcpb.ConfirmTaskPeerInfo, bool) {
	s.confirmPeerInfoLock.RLock()
	peers, ok := s.proposalPeerInfoCache[proposalId]
	s.confirmPeerInfoLock.RUnlock()
	return peers, ok
}

func (s *state) MustQueryConfirmTaskPeerInfo(proposalId common.Hash) *carriertwopcpb.ConfirmTaskPeerInfo {
	peers, _ := s.QueryConfirmTaskPeerInfo(proposalId)
	return peers
}

func (s *state) RemoveConfirmTaskPeerInfo(proposalId common.Hash) {
	s.confirmPeerInfoLock.Lock()
	delete(s.proposalPeerInfoCache, proposalId)
	s.confirmPeerInfoLock.Unlock()
	s.wal.DeleteState(s.wal.GetProposalPeerInfoCacheKey(proposalId))
}

// v 0.3.0 proposal state monitor
func (s *state) CheckProposalStateMonitors(now int64, asyncCall bool) int64 {
	return s.syncProposalStateMonitors.CheckMonitors(now, asyncCall)
}
func (s *state) TimeSleepUntil() int64 {
	return s.syncProposalStateMonitors.TimeSleepUntil()
}
func (s *state) AddMonitor(m *ctypes.ProposalStateMonitor) {
	s.syncProposalStateMonitors.AddMonitor(m)
}
func (s *state) DelMonitor(proposalId common.Hash, partyId string) {
	s.syncProposalStateMonitors.DelMonitor(proposalId, partyId)
}
func (s *state) Timer() *time.Timer {
	return s.syncProposalStateMonitors.Timer()
}
func (s *state) MonitorsLen() int {
	return s.syncProposalStateMonitors.Len()
}
