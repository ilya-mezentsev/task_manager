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

func prepareRequest(method, url, data string) *http.Request {
  req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
  Assert(err == nil, func() {
    panic(err)
  })

  req.Header.Set("Content-Type", "application/json; charset=utf-8")
  return req
}

func makeLoginRequest(method, endpoint, data string) *http.Response {
  srv := httptest.NewServer(router)
  defer srv.Close()

  client := &http.Client{}
  req := prepareRequest(method, fmt.Sprintf("%s/%s", srv.URL, endpoint), data)
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  return resp
}

func makeLoginRequestWithCookie(method, endpoint, data string, cookie *http.Cookie) *http.Response {
  srv := httptest.NewServer(router)
  defer srv.Close()

  client := &http.Client{}
  req := prepareRequest(method, fmt.Sprintf("%s/%s", srv.URL, endpoint), data)
  req.AddCookie(cookie)
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

  resp := makeLoginRequest(http.MethodPost, "session/login", mock.AdminLoginRequestData)
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

  resp := makeLoginRequest(http.MethodPost, "session/login", mock.GroupLeadLoginRequestData)
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

  resp := makeLoginRequest(http.MethodPost, "session/login", mock.GroupWorkerLoginRequestData)
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
  resp := makeLoginRequest(http.MethodPost, "session/login", mock.IncorrectUserNameLoginRequestData)
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
  resp := makeLoginRequest(http.MethodPost, "session/login", mock.IncorrectUserPasswordLoginRequestData)
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
  resp := makeLoginRequest(http.MethodPost, "session/login", mock.GroupWorkerLoginRequestData)
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
  resp := makeLoginRequest(http.MethodPost, "session/logout", mock.AdminLoginRequestData)
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

func TestLoginRequestHandler_GetSessionSuccess(t *testing.T) {
  var response mock.SessionDataResponse
  resp := makeLoginRequestWithCookie(
    http.MethodGet, "session/", mock.GroupWorkerLoginRequestData,
    middleware.CreatAuthCookie(mock.AdminSessionData))
  err := json.NewDecoder(resp.Body).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  Assert(response.Data == mock.AdminSession, func() {
    t.Log(GetExpectationString(mock.AdminSession, response.Data))
    t.Fail()
  })
}

func TestLoginRequestHandler_GetSessionErroredJSON(t *testing.T) {
  var response mock.ErroredResponse
  resp := makeLoginRequestWithCookie(
    http.MethodGet, "session/", mock.GroupWorkerLoginRequestData,
    middleware.CreatAuthCookie(""))
  err := json.NewDecoder(resp.Body).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getSessionError(UnableToDecodeSession).Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestLoginRequestHandler_GetSessionErrorEmptyCookie(t *testing.T) {
  var response mock.ErroredResponse
  resp := makeLoginRequest(
    http.MethodGet, "session/", mock.GroupWorkerLoginRequestData)
  err := json.NewDecoder(resp.Body).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getSessionError(NoAuthTokenInCookie).Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestLoginRequestHandler_GetSessionErrorNoSessionInToken(t *testing.T) {
  badToken, _ := coder.Encrypt(map[string]interface{}{
    "hello": "world",
  })
  var response mock.ErroredResponse
  resp := makeLoginRequestWithCookie(
    http.MethodGet, "session/", mock.GroupWorkerLoginRequestData,
    &http.Cookie{
      Name: "TM-Session-Token",
      Value: badToken,
      Path: "/",
      HttpOnly: true,
      MaxAge: 3600,
    })
  err := json.NewDecoder(resp.Body).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getSessionError(NoSessionInToken).Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}
