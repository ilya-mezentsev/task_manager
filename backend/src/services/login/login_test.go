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
  role, err := ls.GetUserRole(services.AdminLogin, services.AdminPassword)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(role == middleware.RoleAdmin, func() {
    t.Log(GetExpectationString(middleware.RoleAdmin, role))
    t.Fail()
  })
}

func TestLoginService_GetUserRoleWorkLead(t *testing.T) {
  role, err := ls.GetUserRole(services.GroupLeadLogin, services.DefaultPassword)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(role == middleware.RoleGroupLead, func() {
    t.Log(GetExpectationString(middleware.RoleGroupLead, role))
    t.Fail()
  })
}

func TestLoginService_GetUserRoleGroupWorker(t *testing.T) {
  role, err := ls.GetUserRole(services.GroupWorkerLogin, services.DefaultPassword)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(role == middleware.RoleGroupWorker, func() {
    t.Log(GetExpectationString(middleware.RoleGroupWorker, role))
    t.Fail()
  })
}

func TestLoginService_GetUserRoleErrorUserNotFound(t *testing.T) {
  _, err := ls.GetUserRole(services.NotExistsLogin, services.DefaultPassword)

  AssertErrorsEqual(err, UnableToLoginUserNotFound, func() {
    t.Log(GetExpectationString(UnableToLoginUserNotFound, err))
    t.Fail()
  })
}

func TestLoginService_GetUserRoleInternalError(t *testing.T) {
  _, err := ls.GetUserRole(services.ErroredLogin, services.DefaultPassword)

  AssertErrorsEqual(err, UnableToLoginUserInternalError, func() {
    t.Log(GetExpectationString(UnableToLoginUserInternalError, err))
    t.Fail()
  })
}
