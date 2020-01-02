package group_worker

import (
  "mock"
  "testing"
  . "utils"
)

var (
  workerDataMock = mock.GroupWorkerDataMock{
    CompletedTask: make(map[uint]bool),
    TaskComments: make(map[uint]string),
  }
  groupWorker = NewGroupWorker(workerDataMock)
)

func TestAddCommentToTaskSuccess(t *testing.T) {
  err := groupWorker.AddCommentToTask(3, "comment")

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(workerDataMock.IsTaskCommented(3), func() {
    t.Log("task should be commented")
    t.Fail()
  })
}

func TestAddCommentToTaskErrorNotExistsTaskId(t *testing.T) {
  err := groupWorker.AddCommentToTask(mock.NotExistsTaskId, "")

  AssertErrorsEqual(err, mock.UnableToCommentTaskIdNotExists, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.UnableToCommentTaskIdNotExists)
    t.Fail()
  })
  Assert(!workerDataMock.IsTaskCommented(mock.NotExistsTaskId), func() {
    t.Log("task should not be commented")
    t.Fail()
  })
}

func TestAddCommentToTaskCommentingError(t *testing.T) {
  err := groupWorker.AddCommentToTask(mock.CommentingErrorTaskId, "")

  AssertErrorsEqual(err, mock.UnableToCommentInternalError, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.UnableToCommentInternalError)
    t.Fail()
  })
  Assert(!workerDataMock.IsTaskCommented(mock.CommentingErrorTaskId), func() {
    t.Log("task should not be commented")
    t.Fail()
  })
}

func TestMarkTaskAsCompletedSuccess(t *testing.T) {
  err := groupWorker.MarkTaskAsCompleted(3)

  Assert(err == nil, func() {
    t.Log("wrong error:", err)
    t.Fail()
  })
  Assert(workerDataMock.TaskCompleted(3), func() {
    t.Log("task should not be completed")
    t.Fail()
  })
}

func TestMarkTaskAsCompletedErrorNotExistsTaskId(t *testing.T) {
  err := groupWorker.MarkTaskAsCompleted(mock.NotExistsTaskId)

  AssertErrorsEqual(err, mock.UnableToCompleteTaskTaskIdNotExists, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.UnableToCompleteTaskTaskIdNotExists)
    t.Fail()
  })
  Assert(!workerDataMock.TaskCompleted(mock.NotExistsTaskId), func() {
    t.Log("task should not be completed")
    t.Fail()
  })
}

func TestMarkTaskAsCompletedInternalError(t *testing.T) {
  err := groupWorker.MarkTaskAsCompleted(mock.CompletingErrorTaskId)

  AssertErrorsEqual(err, mock.UnableToCompleteInternalError, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.UnableToCompleteInternalError)
    t.Fail()
  })
  Assert(!workerDataMock.TaskCompleted(mock.NotExistsTaskId), func() {
    t.Log("task should not be commented")
    t.Fail()
  })
}
