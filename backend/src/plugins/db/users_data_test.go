package db

import (
  "database/sql"
  "fmt"
  mock "mock/plugins"
  "os"
  "testing"
  . "utils"
)

const (
  deleteFromUsers = "delete from users;"
)

var (
  database  *sql.DB
  adminData UsersDataPlugin
)

func init() {
  var err error
  database, err = sql.Open(
    "sqlite3", "/home/ilya/prog/projects/task_manager/backend/data/test_data.db")
  if err != nil {
    fmt.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  adminData = NewUsersDataPlugin(database)
  deleteAllFromUsers()
}

func execQuery(q string, args ...interface{}) sql.Result {
  statement, err := database.Prepare(q)
  if err != nil {
    fmt.Println("An error while preparing db statement:", err)
    os.Exit(1)
  }

  result, err := statement.Exec(args...)
  if err != nil {
    fmt.Println("An error while creating db structure:", err)
    os.Exit(1)
  }

  return result
}

func createTestUsers() {
  for _, q := range mock.TestingUsersQueries {
    execQuery(q)
  }
}

func deleteAllFromUsers() {
  execQuery(deleteFromUsers)
}

func TestGetAllUsersSuccess(t *testing.T) {
  createTestUsers()
  defer deleteAllFromUsers()

  users, err := adminData.GetAllUsers()

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mock.UserListEqual(users, mock.TestingUsers), func() {
    t.Log("unexpected:", users)
    t.Log("wanted:", mock.TestingUsers)
  })
}
