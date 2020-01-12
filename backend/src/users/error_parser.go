package users

import (
  "fmt"
  "plugins/db"
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
      db.UserNameAlreadyExists: "user name already exists",
    },
  },
  "DeleteUser": {
    errorTemplate: "unable to delete user: %s",
    errorsDetail: map[error]string{
      db.UserIdNotExists: "user id not exists",
    },
  },
  "GetAllGroups": {
    errorTemplate: "unable to get all groups: %s",
  },
  "CreateWorkGroup": {
    errorTemplate: "unable to create work group: %s",
    errorsDetail: map[error]string{
      db.WorkGroupAlreadyExists: "work group already exists",
    },
  },
  "DeleteWorkGroup": {
    errorTemplate: "unable to delete work group: %s",
    errorsDetail: map[error]string{
      db.WorkGroupNotExists: "work group not exists",
    },
  },
  "GetAllTasks": {
    errorTemplate: "unable to get all tasks: %s",
  },
  "AssignTasksToWorkGroup": {
    errorTemplate: "unable to assign tasks: %s",
    errorsDetail: map[error]string{
      db.WorkGroupNotExists: "work group not exists",
    },
  },
  "DeleteTask": {
    errorTemplate: "unable to delete task: %s",
    errorsDetail: map[error]string{
      db.TaskIdNotExists: "task id not exists",
    },
  },

  // group lead errors
  "AssignTaskToWorker": {
    errorTemplate: "unable to assign task: %s",
    errorsDetail: map[error]string{
      db.WorkerIdNotExists: "worker id not exists",
      db.TaskIdNotExists: "task id not exists",
    },
  },
  "GetTasksByGroupId": {
    errorTemplate: "unable to get tasks by group id: %s",
  },

  // group worker errors
  "GetTasksByUserId": {
    errorTemplate: "unable to get tasks by user id: %s",
  },
  "AddCommentToTask": {
    errorTemplate: "unable to comment task: %s",
    errorsDetail: map[error]string{
      db.TaskIdNotExists: "id not exists",
    },
  },
  "MarkTaskAsCompleted": {
    errorTemplate: "unable to complete task: %s",
    errorsDetail: map[error]string{
      db.TaskIdNotExists: "id not exists",
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
