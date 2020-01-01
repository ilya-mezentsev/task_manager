package group_lead

import (
  "mock"
  "models"
  "testing"
  . "utils"
)

var (
  mockGroupLeadData = mock.GroupLeadDataMock{WorkersTasks: map[uint][]models.Task{}}
  groupLead = NewGroupLead(mockGroupLeadData)
)

func TestAssignTaskToWorkerSuccess(t *testing.T) {
  testTask := models.Task{
    Title: "title",
    Description: "desc",
  }
  err := groupLead.AssignTaskToWorker(2, testTask)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mockGroupLeadData.TaskAssigned(2, testTask), func() {
    t.Log("task should be assigned")
    t.Fail()
  })
}

func TestAssignTaskToWorkerErrorWorkerIdNotExists(t *testing.T) {
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

func TestAssignTaskToWorkerInternalError(t *testing.T) {
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
