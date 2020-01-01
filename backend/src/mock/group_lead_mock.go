package mock

import (
  "errors"
  "models"
  "processing"
)

const (
  WorkerIdNotExists = iota
  WorkerIdAssigningError
)

var (
  assigningError = errors.New("assigning error")
  WorkerIdNotExistsError = errors.New("unable to assign task: worker id not exists")
  AssignTaskInternalError = errors.New("unable to assign task: internal error")
)

type GroupLeadDataMock struct {
  WorkersTasks map[uint][]models.Task
}

func (gld GroupLeadDataMock) AssignTaskToWorker(workerId uint, task models.Task) error {
  if workerId == WorkerIdNotExists {
    return processing.WorkerIdNotExists
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
