package api

import (
  "database/sql"
  "fmt"
  "mock"
  mock2 "mock/plugins"
  "os"
  "plugins"
  "plugins/code"
  "plugins/db"
  "plugins/db/groups"
  "plugins/db/tasks"
  "plugins/db/users"
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
