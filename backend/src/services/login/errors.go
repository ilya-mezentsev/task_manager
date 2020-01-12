package login

import "errors"

var (
  UnableToLoginUserNotFound = errors.New("unable to login user: not found by credentials")
  UnableToLoginUserInternalError = errors.New("unable to login user: internal error")
)
