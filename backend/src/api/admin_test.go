package api

import (
  "api/middleware"
  "bytes"
  "database/sql"
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "log"
  "math"
  "mock"
  mock2 "mock/plugins"
  "net/http"
  "net/http/httptest"
  "os"
  "plugins"
  "plugins/db"
  "plugins/db/groups"
  "plugins/db/tasks"
  "plugins/db/users"
  "strings"
  "testing"
  . "utils"
)

var testingHelper mock.TestingHelpers

func init() {
  dbFile := os.Getenv("TEST_DB_FILE")
  if dbFile == "" {
    fmt.Println("TEST_DB_FILE env var is not set")
    os.Exit(1)
  }

  var err error
  testingHelper.Database, err = sql.Open("sqlite3", dbFile)
  if err != nil {
    fmt.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  InitAdminRequestHandler(plugins.NewDBProxy(testingHelper.Database))
  testingHelper.GroupsData = groups.NewDataPlugin(testingHelper.Database)
  testingHelper.UsersData = users.NewDataPlugin(testingHelper.Database)
  testingHelper.TasksData = tasks.NewDataPlugin(testingHelper.Database)
  db.ExecQuery(testingHelper.Database, mock2.TurnOnForeignKeys)
}

func dropTestTables() {
  for _, q := range mock.DropTestTablesQueries {
    db.ExecQuery(testingHelper.Database, q)
  }
}

func initTestTables() {
  dropTestTables()
  for _, q := range mock.CreateTestTablesQueries {
    db.ExecQuery(testingHelper.Database, q)
  }
  for _, q := range mock.AddDataToTestTablesQueries {
    db.ExecQuery(testingHelper.Database, q)
  }
}

func makeRequest(t *testing.T, method, endpoint, data string) io.ReadCloser {
  srv := httptest.NewServer(router)
  defer srv.Close()

  client := &http.Client{}
  req, err := http.NewRequest(
    method,
    fmt.Sprintf("%s/%s", srv.URL, endpoint),
    bytes.NewBuffer([]byte(data)),
  )
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })

  var setTokenFn func(r *http.Request)
  switch {
  case strings.Contains(endpoint, "admin"):
    setTokenFn = func(r *http.Request) {
      middleware.SetAuthCookie(r, mock.AdminSessionData)
    }
  case strings.Contains(endpoint, "lead"):
    setTokenFn = func(r *http.Request) {
      middleware.SetAuthCookie(r, mock.GroupLeadSessionData)
    }
  default:
    setTokenFn = func(r *http.Request) {
      middleware.SetAuthCookie(r, mock.GroupWorkerSessionData)
    }
  }
  req.Header.Set("Content-Type", "application/json; charset=utf-8")
  setTokenFn(req)
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

func TestAdminRequestHandler_RequestWithBadData(t *testing.T) {
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

func TestAdminRequestHandler_GetAllGroupsSuccess(t *testing.T) {
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

func TestAdminRequestHandler_GetAllGroupsErrorTableNotExists(t *testing.T) {
  dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodGet, "admin/groups", "")
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToGetAllGroupsInternal.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_CreateGroupSuccess(t *testing.T) {
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

func TestAdminRequestHandler_CreateGroupErrorIncorrectName(t *testing.T) {
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
  expectedError := getIncorrectGroupNameError("").Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_CreateGroupErrorGroupAlreadyExists(t *testing.T) {
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
  expectedError := mock.UnableToCreateWgAlreadyExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_DeleteGroupSuccess(t *testing.T) {
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
  actualGroups, _ := testingHelper.GroupsData.GetAllGroups()
  expectedGroups := mock2.TestingGroups[1:]
  Assert(mock2.GroupListEqual(expectedGroups, actualGroups), func() {
    t.Log(GetExpectationString(expectedGroups, actualGroups))
    t.Fail()
  })
}

func TestAdminRequestHandler_DeleteGroupBadGroupId(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodDelete, "admin/group", mock.DeleteGroupRequestDataIncorrectId)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getIncorrectGroupIdError(math.MaxUint64).Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_DeleteGroupErrorIdNotExists(t *testing.T) {
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
  expectedError := mock.UnableToDeleteWgIdNotExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_GetAllUsers(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.AllUsersResponse
  responseBody := makeRequest(t, http.MethodGet, "admin/users", "")
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  Assert(mock2.UserListEqual(response.Data, mock2.TestingUsers), func() {
    t.Log(GetExpectationString(mock2.TestingUsers, response.Data))
    t.Fail()
  })
}

func TestAdminRequestHandler_GetAllUsersErrorTableNotExists(t *testing.T) {
  dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodGet, "admin/users", "")
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToGetAllUsersInternal.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_CreateUserSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.EmptyDataResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/user", mock.CreateUserRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  allUsers, _ := testingHelper.UsersData.GetAllUsers()
  actualUser := allUsers[len(allUsers)-1]
  Assert(actualUser == mock.CreatedUser, func() {
    t.Log(GetExpectationString(mock.CreatedUser, actualUser))
  })
}

func TestAdminRequestHandler_CreateUserErrorIncorrectName(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/user", mock.CreateUserRequestDataIncorrectName)
  err := json.NewDecoder(responseBody).Decode(&response)

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

func TestAdminRequestHandler_CreateUserErrorIncorrectGroupId(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(
    t, http.MethodPost, "admin/user", mock.CreateUserRequestDataIncorrectGroupId)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getIncorrectUserGroupIdError(math.MaxUint64).Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_CreateUserErrorAlreadyExists(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/user", mock.CreateUserRequestDataAlreadyExists)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToCreateUserNameAlreadyExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_DeleteUserSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.EmptyDataResponse
  responseBody := makeRequest(t, http.MethodDelete, "admin/user", mock.DeleteUserRequestData)
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
  actualUsers, _ := testingHelper.UsersData.GetAllUsers()
  expectedUsers := mock2.TestingUsers[1:]
  Assert(mock2.UserListEqual(expectedUsers, actualUsers), func() {
    t.Log(GetExpectationString(expectedUsers, actualUsers))
    t.Fail()
  })
}

func TestAdminRequestHandler_DeleteUserErrorIncorrectId(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodDelete, "admin/user", mock.DeleteUserRequestDataIncorrectId)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getIncorrectUserIdError(math.MaxUint64).Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_DeleteUserErrorNotExists(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodDelete, "admin/user", mock.DeleteUserRequestDataNotExists)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToDeleteUserIdNotExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_GetAllTasksSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.TasksResponse
  responseBody := makeRequest(t, http.MethodGet, "admin/tasks", "")
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  actualTasks := response.Data
  expectedTasks, _ := testingHelper.TasksData.GetAllTasks()
  Assert(mock2.TasksListEqual(expectedTasks, actualTasks), func() {
    t.Log(GetExpectationString(expectedTasks, actualTasks))
    t.Fail()
  })
}

func TestAdminRequestHandler_GetAllTasksErrorTableNotExists(t *testing.T) {
  dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodGet, "admin/tasks", "")
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToGetAllTasksInternal.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_AssignTasksToWorkGroupSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.EmptyDataResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/tasks", mock.AssignTasksToWGRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  allTasks, _ := testingHelper.TasksData.GetAllTasks()
  addedTask := allTasks[len(allTasks)-1]
  expectedTask := mock.CreatedTask
  Assert(addedTask == expectedTask, func() {
    t.Log(GetExpectationString(expectedTask, addedTask))
    t.Fail()
  })
}

func TestAdminRequestHandler_AssignTasksToWorkGroupErrorIncorrectGroupId(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/tasks", mock.AssignTasksToWGRequestDataIncorrectGroupId)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getIncorrectGroupIdError(math.MaxUint64).Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_AssignTasksToWorkGroupErrorIncorrectTaskTitle(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/tasks", mock.AssignTasksToWGRequestDataIncorrectTaskTitle)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getIncorrectTaskTitleError("").Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_AssignTasksToWorkGroupErrorIncorrectTaskDescription(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/tasks", mock.AssignTasksToWGRequestDataIncorrectTaskDescription)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getIncorrectTaskDescriptionError("").Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_AssignTasksToWorkGroupErrorGroupNotExists(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "admin/tasks", mock.AssignTasksToWGRequestDataGroupNotExists)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToAssignTasksNotExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_DeleteTaskSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.EmptyDataResponse
  responseBody := makeRequest(t, http.MethodDelete, "admin/task", mock.DeleteTaskRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  actualTasks, _ := testingHelper.TasksData.GetAllTasks()
  expectedTasks := mock2.TestingTasks[1:]
  Assert(mock2.TasksListEqual(actualTasks, expectedTasks), func() {
    t.Log(GetExpectationString(expectedTasks, actualTasks))
    t.Fail()
  })
}

func TestAdminRequestHandler_DeleteTaskErrorIncorrectId(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodDelete, "admin/task", mock.DeleteTaskRequestDataIncorrectId)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getIncorrectTaskIdError(math.MaxUint64).Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestAdminRequestHandler_DeleteTaskErrorNotExists(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodDelete, "admin/task", mock.DeleteTaskRequestDataNotExists)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToDeleteTaskIdNotExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}
