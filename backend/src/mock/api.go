package mock

import (
  "database/sql"
  mock "mock/plugins"
  "models"
  "plugins/db/groups"
  "plugins/db/tasks"
  "plugins/db/users"
  "utils"
)

type (
  TestingHelpers struct {
    Database *sql.DB
    GroupsData groups.DataPlugin
    UsersData users.DataPlugin
    TasksData tasks.DataPlugin
  }
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
  TasksResponse struct {
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

  AssignTasksToWGRequestData = `{"group_id": 1, "tasks": [{"title": "some_title", "description": "hello world"}]}`
  AssignTasksToWGRequestDataIncorrectGroupId = `{"group_id": 18446744073709551615, "tasks": []}`
  AssignTasksToWGRequestDataIncorrectTaskTitle = `{"group_id": 1, "tasks": [{"title": "", "description": "hello world"}]}`
  AssignTasksToWGRequestDataIncorrectTaskDescription = `{"group_id": 1, "tasks": [{"title": "some_title", "description": ""}]}`
  AssignTasksToWGRequestDataGroupNotExists = `{"group_id": 4, "tasks": [{"title": "some_title", "description": "hello world"}]}`

  DeleteTaskRequestData = `{"task_id": 1}`
  DeleteTaskRequestDataIncorrectId = `{"task_id": 18446744073709551615}`
  DeleteTaskRequestDataNotExists = `{"task_id": 4}`

  GroupTasksRequestData = `{"group_id": 1}`
  GroupUsersRequestData = GroupTasksRequestData
  GroupTasksIncorrectIdRequestData = DeleteGroupRequestDataIncorrectId
  GroupUsersIncorrectIdRequestData = DeleteGroupRequestDataIncorrectId
  AssignTaskRequestData = `{"user_id": 1, "task": {"id": 1}}`
  AssignTaskIdNotExistsRequestData = `{"user_id": 1, "task": {"id": 11}}`
  AssignTaskIncorrectWorkerIdRequestData = `{"user_id": 18446744073709551615, "task": {"id": 1}}`
  AssignTaskIncorrectTaskIdRequestData = `{"user_id": 1, "task": {"id": 18446744073709551615}}`

  GetTasksByUserIdRequestData = `{"user_id": 1}`
  GetTasksByIncorrectUserIdRequestData = `{"user_id": 18446744073709551615}`
  CommentTaskRequestData = `{"task_id": 1, "comment": "hello world"}`
  CommentTaskIncorrectIdRequestData = `{"task_id": 18446744073709551615, "comment": "hello world"}`
  IncorrectCommentTaskRequestData = `{"task_id": 1, "comment": ""}`
  CommentTaskIdNotExistsRequestData = `{"task_id": 11, "comment": "hello world"}`
  CompleteTaskRequestData = `{"task_id": 1}`
  CompleteTaskIncorrectIdRequestData = `{"task_id": 18446744073709551615}`
  CompleteTaskIdNotExistsRequestData = `{"task_id": 11}`
)

var (
  AdminAuthCookieValue =
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.6ZyYQatKgUHGaWvfT2-aJ4uYhCRi9PHtT374eMKIC7w"
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
  CreatedTask = models.Task{
    ID: 4,
    Title: "some_title",
    Description: "hello world",
    GroupId: 1,
    UserId: 0,
    IsComplete: false,
    Comment: "",
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
