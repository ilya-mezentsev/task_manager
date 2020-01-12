package mock

import (
  "errors"
  "models"
  "plugins/db"
)

const (
  WorkerIdNotExists = iota
  WorkerIdAssigningError
  GroupIdError
)

var (
  assigningError = errors.New("assigning error")
  groupIdError = errors.New("getting all tasks error")
  WorkerIdNotExistsError = errors.New("unable to assign task: worker id not exists")
  UnableToAssignTaskIdNotExists = errors.New("unable to assign task: task id not exists")
  AssignTaskInternalError = errors.New("unable to assign task: internal error")
  GetTasksByGroupIdInternalError = errors.New("unable to get tasks by group id: internal error")
  GetUsersByGroupIdInternalError = errors.New("unable to get users by group id: internal error")
)

type GroupLeadDataMock struct {
  WorkersTasks map[uint][]models.Task
}

func (gld GroupLeadDataMock) AssignTaskToWorker(workerId uint, task models.Task) error {
  if workerId == WorkerIdNotExists {
    return db.WorkerIdNotExists
  } else if workerId == WorkerIdAssigningError {
    return assigningError
  }

  _, ok := gld.WorkersTasks[workerId]
  if !ok {
    gld.WorkersTasks[workerId] = []models.Task{}
  }

  gld.WorkersTasks[workerId] = append(gld.WorkersTasks[workerId], task)
  return nil
}

func (gld GroupLeadDataMock) TaskAssigned(workerId uint, task models.Task) bool {
  tasks, ok := gld.WorkersTasks[workerId]
  if !ok {
    return false
  }

  for _, t := range tasks {
    if t != task {
      return false
    }
  }

  return true
}

func (gld GroupLeadDataMock) GetTasksByGroupId(groupId uint) ([]models.Task, error) {
  if groupId == GroupIdError {
    return nil, groupIdError
  }

  tasks := gld.WorkersTasks[groupId]
  return tasks, nil
}

func (gld GroupLeadDataMock) GetUsersByGroupId(groupId uint) ([]models.User, error) {
  if groupId == GroupIdError {
    return nil, groupIdError
  }

  return []models.User{}, nil
}
