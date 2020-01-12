package api

import (
  "api/helpers"
  "api/middleware"
  "interfaces"
  "models"
  "net/http"
  "services/login"
)

var loginRequestHandler LoginRequestHandler

type LoginRequestHandler struct {
  loginService login.Service
  checker helpers.InputChecker
}

func InitLoginRequestHandler(loginDataPlugin interfaces.LoginData) {
  loginRequestHandler = LoginRequestHandler{
    loginService: login.NewLoginService(loginDataPlugin),
    checker: helpers.NewInputChecker(),
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

  userRole, err := handler.loginService.GetUserRole(loginReq.UserName, loginReq.UserPassword)
  if err != nil {
    panic(err)
  }

  http.SetCookie(w, middleware.CreatAuthCookie(userRole))
  encodeAndSendResponse(w, nil)
}

func (handler LoginRequestHandler) Logout(w http.ResponseWriter, _ *http.Request) {
  http.SetCookie(w, middleware.GetExpiredAuthCookie())
  encodeAndSendResponse(w, nil)
}
