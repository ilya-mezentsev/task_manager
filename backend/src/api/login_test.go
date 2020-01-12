package api

import (
  "bytes"
  "errors"
  "fmt"
  "mock"
  mock2 "mock/plugins"
  "net/http"
  "net/http/httptest"
  "plugins"
  "plugins/db"
  "testing"
  . "utils"
)

func init() {
  InitLoginRequestHandler(plugins.NewDBProxy(testingHelper.Database))
  db.ExecQuery(testingHelper.Database, mock2.TurnOnForeignKeys)
}

func makeLoginRequest(t *testing.T, method, endpoint, data string) *http.Response {
  srv := httptest.NewServer(router)
  defer srv.Close()

  client := &http.Client{}
  req, err := http.NewRequest(
    method,
    fmt.Sprintf("%s/api/%s", srv.URL, endpoint),
    bytes.NewBuffer([]byte(data)),
  )
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })

  req.Header.Set("Content-Type", "application/json; charset=utf-8")
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  return resp
}

func getAuthCookie(response *http.Response) (*http.Cookie, error) {
  for _, cookie := range response.Cookies() {
    if cookie.Name == "TM-Auth-Token" {
      return cookie, nil
    }
  }

  return nil, errors.New("auth cookie not found")
}

func TestLoginRequestHandler_LoginSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  resp := makeLoginRequest(
    t, http.MethodPost, "session/login", `{"user_name": "tm_admin", "user_password": "tm_password"}`)
  authCookie, err := getAuthCookie(resp)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(authCookie.Value == mock.AdminAuthCookieValue, func() {
    t.Log(GetExpectationString(mock.AdminAuthCookieValue, authCookie.Value))
    t.Fail()
  })
}
