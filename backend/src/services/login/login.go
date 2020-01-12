package login

import (
  "api/middleware"
  "interfaces"
  "plugins/db"
)

var (
  adminLogin, adminPassword string
)

type Service struct {
  dataProvider interfaces.LoginData
}

func init() {
  adminLogin = "tm_login"
  adminPassword = "tm_password"
}

func NewLoginService(provider interfaces.LoginData) Service {
  return Service{dataProvider: provider}
}

func (l Service) GetUserRole(name, password string) (string, error) {
  switch {
  case isAdmin(name, password):
    return middleware.RoleAdmin, nil
  default:
    return l.getRoleFromStorage(name, password)
  }
}

func isAdmin(name, password string) bool {
  return name == adminLogin && password == adminPassword
}

func (l Service) getRoleFromStorage(name, password string) (string, error) {
  user, err := l.dataProvider.GetUserByCredentials(name, password)
  switch err {
  case nil:
    break
  case db.UserNotFoundByCredentials:
    return "", UnableToLoginUserNotFound
  default:
    return "", UnableToLoginUserInternalError
  }

  if user.IsGroupLead {
    return middleware.RoleGroupLead, nil
  } else {
    return middleware.RoleGroupWorker, nil
  }
}
