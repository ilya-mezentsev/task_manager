package login

import (
  "api/middleware"
  "interfaces"
  "models"
  "plugins/db"
  "utils"
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

func (l Service) GetSessionUserData(name, password string) (models.UserSession, error) {
  switch {
  case isAdmin(name, password):
    return models.UserSession{
      Name: adminLogin,
      Role: middleware.RoleAdmin,
    }, nil
  default:
    return l.createSessionFromStorage(name, password)
  }
}

func isAdmin(name, password string) bool {
  return name == adminLogin && password == adminPassword
}

func (l Service) createSessionFromStorage(name, password string) (models.UserSession, error) {
  user, err := l.dataProvider.GetUserByCredentials(name, utils.GetHash(password))
  switch err {
  case nil:
    break
  case db.UserNotFoundByCredentials:
    return models.UserSession{}, UnableToLoginUserNotFound
  default:
    return models.UserSession{}, UnableToLoginUserInternalError
  }

  userSession := models.UserSession{ID: user.ID, Name: name, GroupId: user.GroupId}
  if user.IsGroupLead {
    userSession.Role = middleware.RoleGroupLead
  } else {
    userSession.Role = middleware.RoleGroupWorker
  }
  return userSession, nil
}
