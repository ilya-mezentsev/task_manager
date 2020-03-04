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
  InitGroupWorkerRequestHandler(plugins.NewDBProxy(testingHelper.Database))
  db.ExecQuery(testingHelper.Database, mock2.TurnOnForeignKeys)
}

func TestGroupWorkerRequestHandler_GetTasksByWorkerIdSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.TasksResponse
  responseBody := makeRequest(t, http.MethodPost, "group/worker/tasks", mock.GetTasksByUserIdRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  expectedTasks, _ := testingHelper.TasksData.GetTasksByUserId(1)
  Assert(mock2.TasksListEqual(response.Data, expectedTasks), func() {
    t.Log(GetExpectationString(expectedTasks, response.Data))
    t.Fail()
  })
}

func TestGroupWorkerRequestHandler_GetTasksByIncorrectWorkerIdError(t *testing.T) {
  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "group/worker/tasks", mock.GetTasksByIncorrectUserIdRequestData)
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

func TestGroupWorkerRequestHandler_GetTasksByWorkerIdInternalError(t *testing.T) {
  dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPost, "group/worker/tasks", mock.GetTasksByUserIdRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToGetTasksByUserIdInternalError.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestGroupWorkerRequestHandler_CommentTaskSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.EmptyDataResponse
  responseBody := makeRequest(t, http.MethodPatch, "group/worker/task/comment", mock.CommentTaskRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  task, _ := testingHelper.TasksData.GetAllTasks()
  Assert(task[0].Comment == "hello world", func() {
    t.Log(GetExpectationString("hello world", task[0].Comment))
    t.Fail()
  })
}

func TestGroupWorkerRequestHandler_CommentTaskErrorIncorrectTaskId(t *testing.T) {
  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPatch, "group/worker/task/comment", mock.CommentTaskIncorrectIdRequestData)
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

func TestGroupWorkerRequestHandler_CommentTaskErrorIncorrectComment(t *testing.T) {
  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPatch, "group/worker/task/comment", mock.IncorrectCommentTaskRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := getIncorrectTaskCommentError("").Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestGroupWorkerRequestHandler_CommentTaskErrorIdNotExists(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPatch, "group/worker/task/comment", mock.CommentTaskIdNotExistsRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToCommentTaskIdNotExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}

func TestGroupWorkerRequestHandler_CompleteTaskSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.EmptyDataResponse
  responseBody := makeRequest(t, http.MethodPatch, "group/worker/task/complete", mock.CompleteTaskRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsOk(t, response.Status)
  tasks, _ := testingHelper.TasksData.GetAllTasks()
  Assert(tasks[0].IsComplete, func() {
    t.Log(GetExpectationString(true, tasks[0].IsComplete))
    t.Fail()
  })
}

func TestGroupWorkerRequestHandler_CompleteTaskErrorIncorrectId(t *testing.T) {
  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPatch, "group/worker/task/complete", mock.CompleteTaskIncorrectIdRequestData)
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

func TestGroupWorkerRequestHandler_CompleteTaskErrorIdNotExists(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.ErroredResponse
  responseBody := makeRequest(t, http.MethodPatch, "group/worker/task/complete", mock.CompleteTaskIdNotExistsRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  assertStatusIsError(t, response.Status)
  expectedError := mock.UnableToCompleteTaskTaskIdNotExists.Error()
  Assert(response.ErrorDetail == expectedError, func() {
    t.Log(GetExpectationString(expectedError, response.ErrorDetail))
    t.Fail()
  })
}
