package users

import (
  "fmt"
  "processing"
)

type ErrorParsingData struct {
  errorTemplate string
  errorsDetail map[error]string
}

var parsingData = map[string]ErrorParsingData{
  // admin errors
  "GetAllUsers": {
    errorTemplate: "unable to get all users: %s",
  },
  "CreateUser": {
    errorTemplate: "unable to create user: %s",
    errorsDetail: map[error]string{
      processing.UserNameAlreadyExists: "user name already exists",
    },
  },
  "DeleteUser": {
    errorTemplate: "unable to delete user: %s",
    errorsDetail: map[error]string{
      processing.UserIdNotExists: "user id not exists",
    },
  },
  "GetAllGroups": {
    errorTemplate: "unable to get all groups: %s",
  },
  "CreateWorkGroup": {
    errorTemplate: "unable to create work group: %s",
    errorsDetail: map[error]string{
      processing.WorkGroupAlreadyExists: "work group already exists",
    },
  },
  "DeleteWorkGroup": {
    errorTemplate: "unable to delete work group: %s",
    errorsDetail: map[error]string{
      processing.WorkGroupNotExists: "work group not exists",
    },
  },
  "GetAllTasks": {
    errorTemplate: "unable to get all tasks: %s",
  },
  "AssignTasksToWorkGroup": {
    errorTemplate: "unable to assign tasks: %s",
    errorsDetail: map[error]string{
      processing.WorkGroupNotExists: "work group not exists",
    },
  },
  "DeleteTask": {
    errorTemplate: "unable to delete task: %s",
    errorsDetail: map[error]string{
      processing.TaskIdNotExists: "task id not exists",
    },
  },

  // group lead errors
  "AssignTaskToWorker": {
    errorTemplate: "unable to assign task: %s",
    errorsDetail: map[error]string{
      processing.WorkerIdNotExists: "worker id not exists",
    },
  },

  // group worker errors
  "AddCommentToTask": {
    errorTemplate: "unable to comment task: %s",
    errorsDetail: map[error]string{
      processing.TaskIdNotExists: "id not exists",
    },
  },
  "MarkTaskAsCompleted": {
    errorTemplate: "unable to complete task: %s",
    errorsDetail: map[error]string{
      processing.TaskIdNotExists: "id not exists",
    },
  },
}

func ParseError(methodName string, err error) error {
  data := parsingData[methodName]

  for handledError, errorDescription := range data.errorsDetail {
    if err == handledError {
      return fmt.Errorf(data.errorTemplate, errorDescription)
    }
  }

  return fmt.Errorf(data.errorTemplate, "internal error")
}
