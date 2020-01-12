package services

import (
  "errors"
  "models"
  "plugins/db"
)

const (
  AdminLogin = "tm_login"
  AdminPassword = "tm_password"
  GroupLeadLogin = "name3"
  GroupWorkerLogin = "name2"
  NotExistsLogin = "name1"
  ErroredLogin = "errored"
  DefaultPassword = ""
)

var (
  nameToUser = map[string]models.User{
    "name3": {
      ID: 3,
      Name: "name3",
      GroupId: 1,
      Password: "some_pass",
      IsGroupLead: true,
    },
    "name2": {
      ID: 2,
      Name: "name2",
      GroupId: 2,
      Password: "some_pass",
      IsGroupLead: false,
    },
  }
)

type LoginDataMock struct {}

func (l LoginDataMock) GetUserByCredentials(name, password string) (models.User, error) {
  if name == ErroredLogin {
    return models.User{}, errors.New(ErroredLogin)
  }

  user, found := nameToUser[name]
  if !found {
    return models.User{}, db.UserNotFoundByCredentials
  }

  return user, nil
}
