package mock

import (
  "errors"
  "models"
  "plugins/db"
)

const (
  NotExistsTaskId uint = iota
  CommentingErrorTaskId
  CompletingErrorTaskId
  GettingTasksErrorUserId
)

var (
  commentingError = errors.New("commenting error")
  markingError = errors.New("marking error")
  gettingTasksError = errors.New("getting tasks error")
  MockTasks []models.Task
  UnableToCommentTaskIdNotExists = errors.New("unable to comment task: id not exists")
  UnableToCommentInternalError = errors.New("unable to comment task: internal error")
  UnableToCompleteTaskTaskIdNotExists = errors.New("unable to complete task: id not exists")
  UnableToCompleteInternalError = errors.New("unable to complete task: internal error")
  UnableToGetTasksByUserIdInternalError = errors.New("unable to get tasks by user id: internal error")
)

type GroupWorkerDataMock struct {
  TaskComments map[uint]string
  CompletedTask map[uint]bool
}

func (gwd GroupWorkerDataMock) AddCommentToTask(taskId uint, comment string) error {
  if taskId == NotExistsTaskId {
    return db.TaskIdNotExists
  } else if taskId == CommentingErrorTaskId {
    return commentingError
  }

  gwd.TaskComments[taskId] = comment
  return nil
}

func (gwd GroupWorkerDataMock) IsTaskCommented(taskId uint) bool {
  _, ok := gwd.TaskComments[taskId]
  return ok
}

func (gwd GroupWorkerDataMock) MarkTaskAsCompleted(taskId uint) error {
  if taskId == NotExistsTaskId {
    return db.TaskIdNotExists
  } else if taskId == CompletingErrorTaskId {
    return markingError
  }

  gwd.CompletedTask[taskId] = true
  return nil
}

func (gwd GroupWorkerDataMock) TaskCompleted(taskId uint) bool {
  completed, ok := gwd.CompletedTask[taskId]
  return ok && completed
}

func (gwd GroupWorkerDataMock) GetTasksByUserId(userId uint) ([]models.Task, error) {
  if userId == GettingTasksErrorUserId {
    return nil, gettingTasksError
  }

  return MockTasks, nil
}
