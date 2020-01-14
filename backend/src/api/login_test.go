package api

import (
  "api/middleware"
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
  "mock"
  mock2 "mock/plugins"
  "net/http"
  "net/http/httptest"
  "os"
  "plugins"
  "plugins/code"
  "plugins/db"
  "services/login"
  "testing"
  . "utils"
)

var coder code.Coder

func init() {
  coderKey := os.Getenv("CODER_KEY")
  if coderKey == "" {
    panic("CODER_KEY is not set")
  }
  coder = code.NewCoder(coderKey)

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
    if cookie.Name == "TM-Session-Token" {
      return cookie, nil
    }
  }

  return nil, errors.New("auth cookie not found")
}

func TestLoginRequestHandler_LoginAdminSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  resp := makeLoginRequest(t, http.MethodPost, "session/login", mock.AdminLoginRequestData)
  authCookie, err := getAuthCookie(resp)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  decryptedToken, _ := coder.Decrypt(authCookie.Value)
  Assert(decryptedToken["session"] == mock.AdminSessionData, func() {
    t.Log(GetExpectationString(mock.AdminSessionData, decryptedToken["session"]))
    t.Fail()
  })
}

func TestLoginRequestHandler_LoginGroupLeadSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  resp := makeLoginRequest(t, http.MethodPost, "session/login", mock.GroupLeadLoginRequestData)
  authCookie, err := getAuthCookie(resp)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  decryptedToken, _ := coder.Decrypt(authCookie.Value)
  Assert(decryptedToken["session"] == mock.GroupLeadSessionData, func() {
    t.Log(GetExpectationString(mock.GroupLeadSessionData, decryptedToken["session"]))
    t.Fail()
  })
}

func TestLoginRequestHandler_LoginGroupWorkerSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  resp := makeLoginRequest(t, http.MethodPost, "session/login", mock.GroupWorkerLoginRequestData)
  authCookie, err := getAuthCookie(resp)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  decryptedToken, _ := coder.Decrypt(authCookie.Value)
  Assert(decryptedToken["session"] == mock.GroupWorkerSessionData, func() {
    t.Log(GetExpectationString(mock.GroupWorkerSessionData, decryptedToken["session"]))
    t.Fail()
  })
}

func TestLoginRequestHandler_LoginErrorIncorrectUserName(t *testing.T) {
  var response mock.ErroredResponse
  resp := makeLoginRequest(t, http.MethodPost, "session/login", mock.IncorrectUserNameLoginRequestData)
  err := json.NewDecoder(resp.Body).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getIncorrectUserNameError("").Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestLoginRequestHandler_LoginErrorIncorrectUserPassword(t *testing.T) {
  var response mock.ErroredResponse
  resp := makeLoginRequest(t, http.MethodPost, "session/login", mock.IncorrectUserPasswordLoginRequestData)
  err := json.NewDecoder(resp.Body).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getIncorrectUserPasswordError("").Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestLoginRequestHandler_LoginGroupWorkerErrorTableNotExists(t *testing.T) {
  dropTestTables()

  var response mock.ErroredResponse
  resp := makeLoginRequest(t, http.MethodPost, "session/login", mock.GroupWorkerLoginRequestData)
  err := json.NewDecoder(resp.Body).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := login.UnableToLoginUserInternalError.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestLoginRequestHandler_LogoutSuccess(t *testing.T) {
  resp := makeLoginRequest(t, http.MethodPost, "session/logout", mock.AdminLoginRequestData)
  authCookie, err := getAuthCookie(resp)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  expectedCookie := middleware.GetExpiredAuthCookie()
  Assert(authCookie.Value == expectedCookie.Value, func() {
    t.Log(GetExpectationString(expectedCookie, authCookie))
    t.Fail()
  })
}
