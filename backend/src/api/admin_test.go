package api

import (
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
  "plugins/code"
  "plugins/db"
  "plugins/db/groups"
  "plugins/db/tasks"
  "plugins/db/users"
  "testing"
  . "utils"
)

var (
  adminTestingHelper mock.TestingHelpers
)

func init() {
  var coder = code.NewCoder("123456789012345678901234")
  adminTestingHelper.Token, _ = coder.Encrypt(map[string]interface{}{
    "role": "admin",
  })

  dbFile := os.Getenv("TEST_DB_FILE")
  if dbFile == "" {
    fmt.Println("TEST_DB_FILE env var is not set")
    os.Exit(1)
  }

  var err error
  adminTestingHelper.Database, err = sql.Open("sqlite3", dbFile)
  if err != nil {
    fmt.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  InitAdminRequestHandler(plugins.NewDBProxy(adminTestingHelper.Database))
  adminTestingHelper.GroupsData = groups.NewDataPlugin(adminTestingHelper.Database)
  adminTestingHelper.UsersData = users.NewDataPlugin(adminTestingHelper.Database)
  adminTestingHelper.TasksData = tasks.NewDataPlugin(adminTestingHelper.Database)
  db.ExecQuery(adminTestingHelper.Database, mock2.TurnOnForeignKeys)
}

func dropTestTables() {
  for _, q := range mock.DropTestTablesQueries {
    db.ExecQuery(adminTestingHelper.Database, q)
  }
}

func initTestTables() {
  dropTestTables()
  for _, q := range mock.CreateTestTablesQueries {
    db.ExecQuery(adminTestingHelper.Database, q)
  }
  for _, q := range mock.AddDataToTestTablesQueries {
    db.ExecQuery(adminTestingHelper.Database, q)
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
  req.Header.Set("TM-Session-Token", adminTestingHelper.Token)
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

func TestGetAllGroupsErrorTableNotExists(t *testing.T) {
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
  expectedError := getIncorrectGroupNameError("").Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
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
  expectedError := mock.UnableToCreateWgAlreadyExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
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
  actualGroups, _ := adminTestingHelper.GroupsData.GetAllGroups()
  expectedGroups := mock2.TestingGroups[1:]
  Assert(mock2.GroupListEqual(expectedGroups, actualGroups), func() {
    t.Log(GetExpectationString(expectedGroups, actualGroups))
    t.Fail()
  })
}

func TestDeleteGroupBadGroupId(t *testing.T) {
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
  expectedError := mock.UnableToDeleteWgIdNotExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestGetAllUsers(t *testing.T) {
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

func TestGetAllUsersErrorTableNotExists(t *testing.T) {
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

func TestCreateUserSuccess(t *testing.T) {
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
  allUsers, _ := adminTestingHelper.UsersData.GetAllUsers()
  actualUser := allUsers[len(allUsers)-1]
  Assert(actualUser == mock.CreatedUser, func() {
    t.Log(GetExpectationString(mock.CreatedUser, actualUser))
  })
}

func TestCreateUserErrorIncorrectName(t *testing.T) {
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

func TestCreateUserErrorIncorrectGroupId(t *testing.T) {
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

func TestCreateUserErrorAlreadyExists(t *testing.T) {
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

func TestDeleteUserSuccess(t *testing.T) {
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
  actualUsers, _ := adminTestingHelper.UsersData.GetAllUsers()
  expectedUsers := mock2.TestingUsers[1:]
  Assert(mock2.UserListEqual(expectedUsers, actualUsers), func() {
    t.Log(GetExpectationString(expectedUsers, actualUsers))
    t.Fail()
  })
}

func TestDeleteUserErrorIncorrectId(t *testing.T) {
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

func TestDeleteUserErrorNotExists(t *testing.T) {
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

func TestGetAllTasksSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.AllTasksResponse
  responseBody := makeRequest(t, http.MethodGet, "admin/tasks", "")
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  actualTasks := response.Data
  expectedTasks, _ := adminTestingHelper.TasksData.GetAllTasks()
  Assert(mock2.TasksListEqual(expectedTasks, actualTasks), func() {
    t.Log(GetExpectationString(expectedTasks, actualTasks))
    t.Fail()
  })
}

func TestGetAllTasksErrorTableNotExists(t *testing.T) {
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

func TestAssignTasksToWorkGroupSuccess(t *testing.T) {
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
  allTasks, _ := adminTestingHelper.TasksData.GetAllTasks()
  addedTask := allTasks[len(allTasks)-1]
  expectedTask := mock.CreatedTask
  Assert(addedTask == expectedTask, func() {
    t.Log(GetExpectationString(expectedTask, addedTask))
    t.Fail()
  })
}

func TestAssignTasksToWorkGroupErrorIncorrectGroupId(t *testing.T) {
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

func TestAssignTasksToWorkGroupErrorIncorrectTaskTitle(t *testing.T) {
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

func TestAssignTasksToWorkGroupErrorIncorrectTaskDescription(t *testing.T) {
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

func TestAssignTasksToWorkGroupErrorGroupNotExists(t *testing.T) {
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

func TestDeleteTaskSuccess(t *testing.T) {
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
  actualTasks, _ := adminTestingHelper.TasksData.GetAllTasks()
  expectedTasks := mock2.TestingTasks[1:]
  Assert(mock2.TasksListEqual(actualTasks, expectedTasks), func() {
    t.Log(GetExpectationString(expectedTasks, actualTasks))
    t.Fail()
  })
}

func TestDeleteTaskErrorIncorrectId(t *testing.T) {
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

func TestDeleteTaskErrorNotExists(t *testing.T) {
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
