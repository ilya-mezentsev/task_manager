package api

import (
  "api/helpers"
  "api/middleware"
  "encoding/json"
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
  api := router.PathPrefix("/session").Subrouter()

  api.HandleFunc("/", loginRequestHandler.GetSession).Methods(http.MethodGet)
  api.HandleFunc("/login", loginRequestHandler.Login).Methods(http.MethodPost)
  api.HandleFunc("/logout", loginRequestHandler.Logout).Methods(http.MethodPost)
}

func (handler LoginRequestHandler) GetSession(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  tokenData, err := middleware.GetAuthTokenData(r)
  if err != nil {
    panic(getSessionError(NoAuthTokenInCookie))
  }
  sessionData, found := tokenData["session"]
  if !found {
    panic(getSessionError(NoSessionInToken))
  }

  var userSession models.UserSession
  err = json.Unmarshal([]byte(sessionData.(string)), &userSession)
  if err != nil {
    panic(getSessionError(UnableToDecodeSession))
  }

  encodeAndSendResponse(w, userSession)
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

  userSession, err := handler.loginService.GetSessionUserData(loginReq.UserName, loginReq.UserPassword)
  if err != nil {
    panic(err)
  }

  jsonUserSession, _ := json.Marshal(userSession)
  http.SetCookie(w, middleware.CreatAuthCookie(string(jsonUserSession)))
  encodeAndSendResponse(w, userSession)
}

func (handler LoginRequestHandler) Logout(w http.ResponseWriter, _ *http.Request) {
  http.SetCookie(w, middleware.GetExpiredAuthCookie())
  encodeAndSendResponse(w, nil)
}
