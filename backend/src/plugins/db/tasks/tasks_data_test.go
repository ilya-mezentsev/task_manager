package tasks

import (
  "database/sql"
  "fmt"
  mock "mock/plugins"
  "os"
  "plugins/db"
  "testing"
  . "utils"
)

var (
  tasksDatabase *sql.DB
  tasksData TasksDataPlugin
)

func init() {
  var err error
  tasksDatabase, err = sql.Open(
    "sqlite3", os.Getenv("TEST_DB_FILE"))
  if err != nil {
    fmt.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  tasksData = NewTasksDataPlugin(tasksDatabase)
  execTasksQuery(mock.TurnOnForeignKeys)
  db.CreateGroups(tasksDatabase)
}

func execTasksQuery(q string, args ...interface{}) {
  db.ExecQuery(tasksDatabase, q, args...)
}

func dropTasksTable() {
  execTasksQuery(mock.DropTasksTable)
}

func initTasksTable() {
  dropTasksTable()
  execTasksQuery(mock.CreateTasksTable)
  for _, q := range mock.TestingTasksQueries {
    execTasksQuery(q)
  }
}

func TestGetAllTasksSuccess(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  tasks, err := tasksData.GetAllTasks()

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mock.TasksListEqual(tasks, mock.TestingTasks), func() {
    t.Log(GetExpectationString(mock.TestingTasks, tasks))
    t.Fail()
  })
}
