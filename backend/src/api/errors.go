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
  IncorrectUserGroupIdTemplate = "incorrect user group id: %v"
  IncorrectUserIdTemplate = "incorrect user id: %v"
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

func getIncorrectUserGroupIdError(groupId uint) error {
  return fmt.Errorf(IncorrectUserGroupIdTemplate, groupId)
}

func getIncorrectUserIdError(userId uint) error {
  return fmt.Errorf(IncorrectUserIdTemplate, userId)
}
