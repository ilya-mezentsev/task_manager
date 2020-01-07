package mock

import (
  mock "mock/plugins"
  "models"
)

type (
  AllGroupsResponse struct {
    Status string `json:"status"`
    Data []models.Group `json:"data"`
    ErrorDetail string `json:"error_detail"`
  }
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
