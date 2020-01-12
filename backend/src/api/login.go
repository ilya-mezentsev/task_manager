package api

import (
  "api/helpers"
  "api/middleware"
  "interfaces"
  "models"
  "net/http"
  "utils"
)

var loginRequestHandler LoginRequestHandler

type LoginRequestHandler struct {
  loginDataProvider interfaces.LoginData
  checker helpers.InputChecker
  adminLogin, adminPassword string
}

func InitLoginRequestHandler(loginDataPlugin interfaces.LoginData) {
  loginRequestHandler = LoginRequestHandler{
    loginDataProvider: loginDataPlugin,
    checker: helpers.NewInputChecker(),
    adminLogin: "tm_admin", adminPassword: "tm_password",
  }
  bindLoginRoutesToHandlers()
}

func bindLoginRoutesToHandlers() {
  api := router.PathPrefix("/api/session").Subrouter()

  api.HandleFunc("/login", loginRequestHandler.Login).Methods(http.MethodPost)
  api.HandleFunc("/logout", loginRequestHandler.Logout).Methods(http.MethodPost)
}

func (handler LoginRequestHandler) Login(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var loginReq models.LoginRequest
  decodeRequestBody(r, &loginReq)

  if !handler.checker.IsStringCorrect(loginReq.UserName) {
    panic(getIncorrectUserNameError(loginReq.UserName))
  } else if !handler.checker.IsStringCorrect(loginReq.UserPassword) {
    panic(getIncorrectUserPasswordError(loginReq.UserPassword))
  }

  http.SetCookie(w, middleware.CreatAuthCookie(handler.getUserRole(loginReq)))
  encodeAndSendResponse(w, nil)
}

func (handler LoginRequestHandler) getUserRole(loginReq models.LoginRequest) string {
  var userRole string
  if handler.isAdmin(loginReq) {
    userRole = middleware.RoleAdmin
  } else {
    userRole = handler.getRoleFromDB(loginReq)
  }

  return userRole
}

func (handler LoginRequestHandler) getRoleFromDB(loginReq models.LoginRequest) string {
  user, err := handler.loginDataProvider.GetUserByCredentials(
    loginReq.UserName, utils.GetHash(loginReq.UserPassword))
  if err != nil {
    panic(err)
  }

  if user.IsGroupLead {
    return middleware.RoleGroupLead
  } else {
    return middleware.RoleGroupWorker
  }
}

func (handler LoginRequestHandler) isAdmin(loginReq models.LoginRequest) bool {
  return loginReq.UserName == handler.adminLogin && loginReq.UserPassword == handler.adminPassword
}

func (handler LoginRequestHandler) Logout(w http.ResponseWriter, _ *http.Request) {
  http.SetCookie(w, middleware.GetExpiredAuthCookie())
  encodeAndSendResponse(w, nil)
}
