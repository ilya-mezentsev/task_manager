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
)

func getIncorrectGroupNameError(groupName string) error {
  return fmt.Errorf(IncorrectGroupNameTemplate, groupName)
}
