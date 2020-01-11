package api

import (
  "database/sql"
  "encoding/json"
  "fmt"
  "mock"
  mock2 "mock/plugins"
  "net/http"
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
  groupLeadTestingHelpers mock.TestingHelpers
)

func init() {
  var coder = code.NewCoder("123456789012345678901234")
  groupLeadTestingHelpers.Token, _ = coder.Encrypt(map[string]interface{}{
    "role": "admin",
  })

  dbFile := os.Getenv("TEST_DB_FILE")
  if dbFile == "" {
    fmt.Println("TEST_DB_FILE env var is not set")
    os.Exit(1)
  }

  var err error
  groupLeadTestingHelpers.Database, err = sql.Open("sqlite3", dbFile)
  if err != nil {
    fmt.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  InitGroupLeadRequestHandler(plugins.NewDBProxy(groupLeadTestingHelpers.Database))
  groupLeadTestingHelpers.GroupsData = groups.NewDataPlugin(groupLeadTestingHelpers.Database)
  groupLeadTestingHelpers.UsersData = users.NewDataPlugin(groupLeadTestingHelpers.Database)
  groupLeadTestingHelpers.TasksData = tasks.NewDataPlugin(groupLeadTestingHelpers.Database)
  db.ExecQuery(groupLeadTestingHelpers.Database, mock2.TurnOnForeignKeys)
}

func TestGetGroupTasksSuccess(t *testing.T) {
  initTestTables()
  defer dropTestTables()

  var response mock.TasksResponse
  responseBody := makeRequest(t, http.MethodGet, "group/lead/tasks", mock.GroupTasksRequestData)
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  t.Log(response.Data)
}
