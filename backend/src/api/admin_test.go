package api

import (
  "bytes"
  "database/sql"
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "log"
  "mock"
  mock2 "mock/plugins"
  "net/http"
  "net/http/httptest"
  "os"
  "plugins"
  "plugins/db"
  "plugins/db/groups"
  "testing"
  . "utils"
)

var (
  database *sql.DB
  groupsData groups.DataPlugin
)

func init() {
  dbFile := os.Getenv("TEST_DB_FILE")
  if dbFile == "" {
    fmt.Println("TEST_DB_FILE env var is not set")
    os.Exit(1)
  }

  var err error
  database, err = sql.Open("sqlite3", dbFile)
  if err != nil {
    fmt.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  InitAdminRequestHandler(plugins.NewDBProxy(database))
  groupsData = groups.NewDataPlugin(database)
}

func dropTestTables() {
  for _, q := range mock.DropTestTablesQueries {
    db.ExecQuery(database, q)
  }
}

func initTestTables() {
  dropTestTables()
  for _, q := range mock.CreateTestTablesQueries {
    db.ExecQuery(database, q)
  }
  for _, q := range mock.AddDataToTestTablesQueries {
    db.ExecQuery(database, q)
  }
}

func makeRequest(t *testing.T, method, endpoint, data string) io.ReadCloser {
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

  return resp.Body
}

func assertStatusesEqual(t *testing.T, actualStatus, expectedStatus string) {
  Assert(actualStatus == expectedStatus, func() {
    t.Log(GetExpectationString(expectedStatus, actualStatus))
    t.Fail()
  })
}

func assertStatusIsOk(t *testing.T, responseStatus string) {
  assertStatusesEqual(t, responseStatus, "ok")
}

func assertStatusIsError(t *testing.T, responseStatus string) {
  assertStatusesEqual(t, responseStatus, "error")
}

func TestMain(m *testing.M) {
  log.SetOutput(ioutil.Discard)
  os.Exit(m.Run())
}

func TestRequestWithBadData(t *testing.T) {
  var response mock.AllGroupsResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/group", mock.BadRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  Assert(response.ErrorDetail == CannotDecodeRequestBody.Error(), func() {
    t.Log(GetExpectationString(CannotDecodeRequestBody.Error(), response.ErrorDetail))
    t.Fail()
  })
}

func TestGetAllGroupsSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.AllGroupsResponse
  responseBody := makeRequest(t, http.MethodGet, "admin/groups", "")
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  Assert(mock2.GroupListEqual(response.Data, mock2.TestingGroups), func() {
    t.Log(GetExpectationString(mock2.TestingGroups, response.Data))
    t.Fail()
  })
}

func TestCreateGroupSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.CreateGroupResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/group", mock.CreateGroupRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  Assert(response.Data == nil, func() {
    t.Log(GetExpectationString(nil, response.Data))
    t.Fail()
  })
}

func TestCreateGroupErrorIncorrectName(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.CreateGroupResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/group", mock.CreateGroupRequestDataEmptyName)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  Assert(response.ErrorDetail == getIncorrectGroupNameError("").Error(), func() {
    t.Log(GetExpectationString(getIncorrectGroupNameError("").Error(), response.ErrorDetail))
    t.Fail()
  })
}

func TestCreateGroupErrorGroupAlreadyExists(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.CreateGroupResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/group", mock.CreateGroupRequestDataAlreadyExists)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  Assert(response.ErrorDetail == mock.UnableToCreateWgAlreadyExists.Error(), func() {
    t.Log(GetExpectationString(mock.UnableToCreateWgAlreadyExists.Error(), response.ErrorDetail))
    t.Fail()
  })
}

func TestDeleteGroupSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.EmptyDataResponse
  responseBody := makeRequest(t, http.MethodDelete, "admin/group", mock.DeleteGroupRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  Assert(response.Data == nil, func() {
    t.Log(GetExpectationString(nil, response.Data))
    t.Fail()
  })
  actualGroups, _ := groupsData.GetAllGroups()
  expectedGroups := mock2.TestingGroups[1:]
  Assert(len(actualGroups) == len(expectedGroups), func() {
    t.Log(GetExpectationString(expectedGroups, actualGroups))
    t.Fail()
  })
}

func TestDeleteGroupErrorIdNotExists(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodDelete, "admin/group", mock.DeleteGroupRequestDataIdNotExists)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  Assert(response.ErrorDetail == mock.UnableToDeleteWgIdNotExists.Error(), func() {
    t.Log(GetExpectationString(mock.UnableToDeleteWgIdNotExists.Error(), response.ErrorDetail))
    t.Fail()
  })
}
