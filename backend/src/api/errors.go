package api

import (
  "errors"
  "fmt"
)

var (
  ReadRequestBodyError = errors.New("read request body error")
  CannotDecodeRequestBody = errors.New("unable to decode request body")
  CannotWriteResponse = errors.New("unable to write response")

  IncorrectGroupNameTemplate = "incorrect group name: '%s'"
  IncorrectGroupIdTemplate = "incorrect group id: %v"

  IncorrectUserNameTemplate = "incorrect user name: '%s'"
  IncorrectUserPasswordTemplate = "incorrect user password: '%s'"
  IncorrectUserGroupIdTemplate = "incorrect user group id: %v"
  IncorrectUserIdTemplate = "incorrect user id: %v"

  IncorrectTaskTitleTemplate = "incorrect task title: '%s'"
  IncorrectTaskDescriptionTemplate = "incorrect task description: '%s'"
  IncorrectTaskIdTemplate = "incorrect task id: %v"
  IncorrectTaskComment = "incorrect task comment: %s"
)

func getIncorrectGroupNameError(groupName string) error {
  return fmt.Errorf(IncorrectGroupNameTemplate, groupName)
}

func getIncorrectGroupIdError(groupId uint) error {
  return fmt.Errorf(IncorrectGroupIdTemplate, groupId)
}

func getIncorrectUserNameError(userName string) error {
  return fmt.Errorf(IncorrectUserNameTemplate, userName)
}

func getIncorrectUserPasswordError(userPassword string) error {
  return fmt.Errorf(IncorrectUserPasswordTemplate, userPassword)
}

func getIncorrectUserGroupIdError(groupId uint) error {
  return fmt.Errorf(IncorrectUserGroupIdTemplate, groupId)
}

func getIncorrectUserIdError(userId uint) error {
  return fmt.Errorf(IncorrectUserIdTemplate, userId)
}

func getIncorrectTaskTitleError(title string) error {
  return fmt.Errorf(IncorrectTaskTitleTemplate, title)
}

func getIncorrectTaskDescriptionError(description string) error {
  return fmt.Errorf(IncorrectTaskDescriptionTemplate, description)
}

func getIncorrectTaskIdError(taskId uint) error {
  return fmt.Errorf(IncorrectTaskIdTemplate, taskId)
}

func getIncorrectTaskCommentError(comment string) error {
  return fmt.Errorf(IncorrectTaskComment, comment)
}
