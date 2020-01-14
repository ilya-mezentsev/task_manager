package login

import (
  "api/middleware"
  "mock/services"
  "testing"
  . "utils"
)

var (
  mockLoginData = services.LoginDataMock{}
  ls = NewLoginService(mockLoginData)
)

func TestLoginService_GetUserRoleAdmin(t *testing.T) {
  userSession, err := ls.GetSessionUserData(services.AdminLogin, services.AdminPassword)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(userSession.Role == middleware.RoleAdmin, func() {
    t.Log(GetExpectationString(middleware.RoleAdmin, userSession.Role))
    t.Fail()
  })
}

func TestLoginService_GetUserRoleWorkLead(t *testing.T) {
  userSession, err := ls.GetSessionUserData(services.GroupLeadLogin, services.DefaultPassword)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(userSession.Role == middleware.RoleGroupLead, func() {
    t.Log(GetExpectationString(middleware.RoleGroupLead, userSession.Role))
    t.Fail()
  })
}

func TestLoginService_GetUserRoleGroupWorker(t *testing.T) {
  userSession, err := ls.GetSessionUserData(services.GroupWorkerLogin, services.DefaultPassword)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(userSession.Role == middleware.RoleGroupWorker, func() {
    t.Log(GetExpectationString(middleware.RoleGroupWorker, userSession.Role))
    t.Fail()
  })
}

func TestLoginService_GetUserRoleErrorUserNotFound(t *testing.T) {
  _, err := ls.GetSessionUserData(services.NotExistsLogin, services.DefaultPassword)

  AssertErrorsEqual(err, UnableToLoginUserNotFound, func() {
    t.Log(GetExpectationString(UnableToLoginUserNotFound, err))
    t.Fail()
  })
}

func TestLoginService_GetUserRoleInternalError(t *testing.T) {
  _, err := ls.GetSessionUserData(services.ErroredLogin, services.DefaultPassword)

  AssertErrorsEqual(err, UnableToLoginUserInternalError, func() {
    t.Log(GetExpectationString(UnableToLoginUserInternalError, err))
    t.Fail()
  })
}
