package api

import (
  "encoding/json"
  "math"
  "mock"
  mock2 "mock/plugins"
  "net/http"
  "plugins"
  "plugins/db"
  "testing"
  . "utils"
)

func init() {
  InitGroupLeadRequestHandler(plugins.NewDBProxy(testingHelper.Database))
  db.ExecQuery(testingHelper.Database, mock2.TurnOnForeignKeys)
}

func TestGroupLeadRequestHandler_GetGroupTasksSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.TasksResponse
  responseBody := makeRequest(t, http.MethodGet, "group/lead/tasks", mock.GroupTasksRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  Assert(mock2.TasksListEqual(response.Data, mock2.TestingTasksByGroupId), func() {
    t.Log(GetExpectationString(mock2.TestingTasksByGroupId, response.Data))
    t.Fail()
  })
}

func TestGroupLeadRequestHandler_GetTasksByIncorrectGroupIdError(t *testing.T) {
  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodGet, "group/lead/tasks", mock.GroupTasksIncorrectIdRequestData)
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

func TestGroupLeadRequestHandler_GetTasksByGroupIdErrorTableNotExists(t *testing.T) {
  dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodGet, "group/lead/tasks", mock.GroupTasksRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.GetTasksByGroupIdInternalError.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestGroupLeadRequestHandler_AssignTaskToWorkerSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.EmptyDataResponse
  responseBody := makeRequest(t, http.MethodPost, "group/lead/task", mock.AssignTaskRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  tasks, _ := testingHelper.TasksData.GetAllTasks()
  Assert(tasks[0].UserId == 1, func() {
    t.Log(GetExpectationString(1, tasks[0].UserId))
    t.Fail()
  })
}

func TestGroupLeadRequestHandler_AssignTaskToWorkerErrorTaskIdNotExists(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "group/lead/task", mock.AssignTaskIdNotExistsRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToAssingTaskIdNotExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestGroupLeadRequestHandler_AssignTaskToWorkerErrorIncorrectWorkerId(t *testing.T) {
  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "group/lead/task", mock.AssignTaskIncorrectWorkerIdRequestData)
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

func TestGroupLeadRequestHandler_AssignTaskToWorkerErrorIncorrectTaskId(t *testing.T) {
  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "group/lead/task", mock.AssignTaskIncorrectTaskIdRequestData)
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
