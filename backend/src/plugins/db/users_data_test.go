package db

import (
  "database/sql"
  "fmt"
  mock "mock/plugins"
  "os"
  "processing"
  "testing"
  . "utils"
)

var (
  database  *sql.DB
  usersData UsersDataPlugin
)

func init() {
  var err error
  database, err = sql.Open(
    "sqlite3", "/home/ilya/prog/projects/task_manager/backend/data/test_data.db")
  if err != nil {
    fmt.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  usersData = NewUsersDataPlugin(database)
  createUsersTable()
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

func createUsersTable() {
  execQuery(mock.CreateUsersTable)
}

func deleteUsersTable() {
  execQuery(mock.DropUsersTable)
}

func createTestUsers() {
  for _, q := range mock.TestingUsersQueries {
    execQuery(q)
  }
}

func deleteAllFromUsers() {
  execQuery(mock.ClearUsersTable)
}

func TestGetAllUsersSuccess(t *testing.T) {
  createTestUsers()
  defer deleteAllFromUsers()

  users, err := usersData.GetAllUsers()

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mock.UserListEqual(users, mock.TestingUsers), func() {
    t.Log("unexpected:", users)
    t.Log("wanted:", mock.TestingUsers)
    t.Fail()
  })
}

func TestGetAllUsersErrorTableNotExists(t *testing.T) {
  deleteUsersTable()
  defer createUsersTable()

  users, err := usersData.GetAllUsers()

  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
  Assert(users == nil, func() {
    t.Log("should not be users:", users)
    t.Fail()
  })
}

func TestGetUserSuccess(t *testing.T) {
  createTestUsers()
  defer deleteAllFromUsers()

  user, err := usersData.GetUser(1)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(user == mock.TestingUsers[0], func() {
    t.Log("unexpected user:", user)
    t.Log("wanted:", mock.TestingUsers[0])
    t.Fail()
  })
}

func TestGetUserErrorIdNotExists(t *testing.T) {
  createTestUsers()
  defer deleteAllFromUsers()

  user, err := usersData.GetUser(11)
  Assert(err == processing.WorkerIdNotExists, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", processing.WorkerIdNotExists)
    t.Fail()
  })
  Assert(user == mock.EmptyUser, func() {
    t.Log("unexpected user:", user)
    t.Log("wanted:", mock.EmptyUser)
    t.Fail()
  })
}

func TestGetUserErrorTableNotExists(t *testing.T) {
  deleteUsersTable()
  defer createUsersTable()

  user, err := usersData.GetUser(11)
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
  Assert(user == mock.EmptyUser, func() {
    t.Log("unexpected user:", user)
    t.Log("wanted:", mock.EmptyUser)
    t.Fail()
  })
}

func TestCreateUserSuccess(t *testing.T) {
  deleteUsersTable()
  createUsersTable()
  defer deleteAllFromUsers()

  createdUserId, err := usersData.CreateUser(mock.TestingUser)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(createdUserId == 1, func() {
    t.Log("unexpected created user id:", createdUserId)
    t.Log("wanted:", 1)
    t.Fail()
  })
}

func TestCreateUserErrorTableNotExists(t *testing.T) {
  deleteUsersTable()
  defer createUsersTable()

  createdUserId, err := usersData.CreateUser(mock.TestingUser)
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
  Assert(createdUserId == 0, func() {
    t.Log("unexpected created user id:", createdUserId)
    t.Log("wanted:", 0)
    t.Fail()
  })
}

func TestCreateUserErrorNameAlreadyExists(t *testing.T) {
  deleteUsersTable()
  createUsersTable()
  createTestUsers()
  defer deleteAllFromUsers()

  createdUserId, err := usersData.CreateUser(mock.TestingUserWithExistsName)
  Assert(err == processing.UserNameAlreadyExists, func() {
    t.Log("should be error")
    t.Fail()
  })
  Assert(createdUserId == 0, func() {
    t.Log("unexpected created user id:", createdUserId)
    t.Log("wanted:", 0)
    t.Fail()
  })
}

func TestDeleteUserSuccess(t *testing.T) {
  deleteUsersTable()
  createUsersTable()
  createTestUsers()
  defer deleteAllFromUsers()

  err := usersData.DeleteUser(1)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
}

func TestDeleteUserErrorTableNotExists(t *testing.T) {
  deleteUsersTable()
  defer createUsersTable()

  err := usersData.DeleteUser(1)
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}
