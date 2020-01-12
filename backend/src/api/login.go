package api

import (
  "api/helpers"
  "interfaces"
  "net/http"
)

var loginRequestHandler LoginRequestHandler

type LoginRequestHandler struct {
  loginDataProvider interfaces.LoginData
  checker helpers.InputChecker
}

func InitLoginRequestHandler(loginDataPlugin interfaces.LoginData) {
  loginRequestHandler = LoginRequestHandler{
    loginDataProvider: loginDataPlugin,
    checker: helpers.NewInputChecker(),
  }
  bindLoginRoutesToHandlers()
}

func bindLoginRoutesToHandlers() {
  api := router.PathPrefix("/api/session").Subrouter()

  api.HandleFunc("/login", loginRequestHandler.Login).Methods(http.MethodPost)
}

func (handler LoginRequestHandler) Login(w http.ResponseWriter, r *http.Request) {}
