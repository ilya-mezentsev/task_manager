package processing

import "errors"

var (
  UserIdNotExists = errors.New("user id not exists")
  UserNameAlreadyExists = errors.New("user name already exists")
  WorkGroupAlreadyExists = errors.New("work group already exists")
  WorkGroupNotExists = errors.New("work group not exists")

  WorkerIdNotExists = errors.New("worker id not exists")

  TaskIdNotExists = errors.New("task id not exists")
)
