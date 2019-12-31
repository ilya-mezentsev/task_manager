package admin

import "errors"

var (
  UnableToCreateUser = errors.New("unable to create user")
  UnableToCreateWorkGroup = errors.New("unable to create work group")
  UnableToAssignTasks = errors.New("unable to assign tasks")
)
