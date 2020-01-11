package group_lead

import (
  "mock"
  "models"
  "testing"
  . "utils"
)

var (
  safeWorkerId uint = 3
  mockGroupLeadData = mock.GroupLeadDataMock{WorkersTasks: map[uint][]models.Task{}}
  groupLead = NewGroupLead(mockGroupLeadData)
)

func TestGroupLead_AssignTaskToWorkerSuccess(t *testing.T) {
  testTask := models.Task{
    Title: "title",
    Description: "desc",
  }
  err := groupLead.AssignTaskToWorker(safeWorkerId, testTask)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mockGroupLeadData.TaskAssigned(safeWorkerId, testTask), func() {
    t.Log("task should be assigned")
    t.Fail()
  })
}

func TestGroupLead_AssignTaskToWorkerErrorWorkerIdNotExists(t *testing.T) {
  testTask := models.Task{
    Title: "title",
    Description: "desc",
  }
  err := groupLead.AssignTaskToWorker(mock.WorkerIdNotExists, testTask)

  AssertErrorsEqual(err, mock.WorkerIdNotExistsError, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.WorkerIdNotExistsError)
    t.Fail()
  })
  Assert(!mockGroupLeadData.TaskAssigned(mock.WorkerIdNotExists, testTask), func() {
    t.Log("task should not be assigned")
    t.Fail()
  })
}

func TestGroupLead_AssignTaskToWorkerInternalError(t *testing.T) {
  testTask := models.Task{
    Title: "title",
    Description: "desc",
  }
  err := groupLead.AssignTaskToWorker(mock.WorkerIdAssigningError, testTask)

  AssertErrorsEqual(err, mock.AssignTaskInternalError, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.AssignTaskInternalError)
    t.Fail()
  })
  Assert(!mockGroupLeadData.TaskAssigned(mock.WorkerIdNotExists, testTask), func() {
    t.Log("task should not be assigned")
    t.Fail()
  })
}

func TestGroupLead_GetTasksByGroupIdSuccess(t *testing.T) {
  var testTask = models.Task{
    Title: "title",
    Description: "desc",
  }
  mockGroupLeadData.WorkersTasks[safeWorkerId] = []models.Task{testTask}

  tasks, err := groupLead.GetTasksByGroupId(safeWorkerId)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  lastTask := tasks[len(tasks)-1]
  Assert(lastTask == testTask, func() {
    t.Log(GetExpectationString(testTask, lastTask))
    t.Fail()
  })
}

func TestGroupLead_GetTasksByGroupIdInternalError(t *testing.T) {
  tasks, err := groupLead.GetTasksByGroupId(mock.GroupIdError)

  AssertErrorsEqual(err, mock.GetTasksByGroupIdInternalError, func() {
    t.Log(GetExpectationString(mock.GetTasksByGroupIdInternalError, err))
    t.Fail()
  })
  Assert(tasks == nil, func() {
    t.Log(GetExpectationString(nil, tasks))
    t.Fail()
  })
}
