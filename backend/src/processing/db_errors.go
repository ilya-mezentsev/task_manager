package processing

import "errors"

var (
  UserNameAlreadyExists = errors.New("user name already exists")
  WorkGroupAlreadyExists = errors.New("work group already exists")
  WorkGroupNotExists = errors.New("work group not exists")

  WorkerIdNotExists = errors.New("worker id not exists")

  TaskIdNotExists = errors.New("task id not exists")
)
