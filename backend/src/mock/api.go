package mock

import (
  mock "mock/plugins"
  "models"
  "utils"
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
  AllUsersResponse struct {
    Status string `json:"status"`
    Data []models.User `json:"data"`
  }
  AllTasksResponse struct {
    Status string `json:"status"`
    Data []models.Task `json:"data"`
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
  DeleteGroupRequestDataIncorrectId = `{"group_id": 18446744073709551615}`

  CreateUserRequestData = `{"user": {"name": "name4", "group_id": 1, "is_group_lead": false}}`
  CreateUserRequestDataIncorrectName = `{"user": {"name": ""}}`
  CreateUserRequestDataIncorrectGroupId = `{"user": {"group_id": 18446744073709551615}}`
  CreateUserRequestDataAlreadyExists = `{"user": {"name": "name1"}}`

  DeleteUserRequestData = `{"user_id": 1}`
  DeleteUserRequestDataIncorrectId = `{"user_id": 18446744073709551615}`
  DeleteUserRequestDataNotExists = `{"user_id": 4}`
)

var (
  DropTestTablesQueries = []string{
    mock.DropGroupsTable, mock.DropUsersTable, mock.DropTasksTable,
  }
  CreateTestTablesQueries = []string{
    mock.CreateGroupsTable, mock.CreateUsersTable, mock.CreateTasksTable,
  }
  AddDataToTestTablesQueries = getTestingDataQueries()
  CreatedUser = models.User{
    ID: 4,
    Name: "name4",
    GroupId: 1,
    Password: utils.GetHash("name4_1"),
    IsGroupLead: false,
  }
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
