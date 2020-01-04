package db

import (
  "database/sql"
  "fmt"
  "os"
  "testing"
  . "utils"
)

const (
  deleteFromUsers = "delete from users;"
)

var (
  testingUsers = []string{
    "insert into users values(1, 'name1', 1, 'some_pass', 0);",
    "insert into users values(2, 'name2', 2, 'some_pass', 0);",
    "insert into users values(3, 'name3', 1, 'some_pass', 1);",
  }
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

func execQuery(q string) {
  statement, err := database.Prepare(q)
  if err != nil {
    fmt.Println("An error while preparing db statement:", err)
    os.Exit(1)
  }

  _, err = statement.Exec()
  if err != nil {
    fmt.Println("An error while creating db structure:", err)
    os.Exit(1)
  }
}

func createTestUsers() {
  for _, q := range testingUsers {
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
  t.Log(users)
}
