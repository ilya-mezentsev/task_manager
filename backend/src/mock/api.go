package mock

import (
  mock "mock/plugins"
  "models"
)

type (
  ErroredResponse struct {
    Status string `json:"status"`
    ErrorDetail string `json:"error_detail"`
  }
  AllGroupsResponse struct {
    Status string `json:"status"`
    Data []models.Group `json:"data"`
    ErrorDetail string `json:"error_detail"`
  }
  CreateGroupResponse struct {
    Status string `json:"status"`
    Data interface{} `json:"data"`
    ErrorDetail string `json:"error_detail"`
  }
  EmptyDataResponse struct {
    Status string `json:"status"`
    Data interface{} `json:"data"`
  }
)

const (
  BadRequestData = ``
  CreateGroupRequestData = `{"group_name": "group4"}`
  CreateGroupRequestDataAlreadyExists = `{"group_name": "group1"}`
  CreateGroupRequestDataEmptyName = `{"group_name": ""}`

  DeleteGroupRequestData = `{"group_id": 1}`
  DeleteGroupRequestDataIdNotExists = `{"group_id": 4}`
)

var (
  DropTestTablesQueries = []string{
    mock.DropGroupsTable, mock.DropUsersTable, mock.DropTasksTable,
  }
  CreateTestTablesQueries = []string{
    mock.CreateGroupsTable, mock.CreateUsersTable, mock.CreateTasksTable,
  }
  AddDataToTestTablesQueries = getTestingDataQueries()
)

func getTestingDataQueries() []string {
  var allTestingDataQueries []string
  for _, testingDataQueries := range [][]string{
    mock.TestingGroupsQueries, mock.TestingUsersQueries, mock.TestingTasksQueries,
  } {
    allTestingDataQueries = append(allTestingDataQueries, testingDataQueries...)
  }
  return allTestingDataQueries
}
