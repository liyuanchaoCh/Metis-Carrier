package rawdb

import (
	"github.com/datumtechs/datum-network-carrier/db"
	libtypes "github.com/datumtechs/datum-network-carrier/lib/types"
	"github.com/datumtechs/datum-network-carrier/types"
	"gotest.tools/assert"
	"strings"
	"testing"
)

func TestRunningTask(t *testing.T) {
	database := db.NewMemoryDatabase()
	task := types.NewTask(&libtypes.TaskPB{
		Sender: &libtypes.TaskOrganization{
			PartyId:    "p0",
			IdentityId: "identity-task",
			NodeId:     "nodeId-task",
			NodeName:   "nodeName",
		},
		DataId:     "",
		DataStatus: libtypes.DataStatus_DataStatus_Valid,
		TaskId:     "taskID-01",
		TaskName:   "taskName-01",
		State:      libtypes.TaskState_TaskState_Succeed,
		Reason:     "",
		Desc:       "",
		CreateAt:   0,
		EndAt:      0,
	})
	WriteRunningTask(database, task)

	rtask := ReadRunningTask(database, "taskID-01")
	t.Logf("running task info : %v", rtask)
	assert.Assert(t, strings.EqualFold("taskID-01", rtask.GetTaskId()))

	// read all
	taskList := ReadAllRunningTask(database)
	assert.Assert(t, len(taskList) == 1)

	// delete
	DeleteRunningTask(database, "taskID-01")

	taskList = ReadAllRunningTask(database)
	assert.Assert(t, len(taskList) == 0)

}
