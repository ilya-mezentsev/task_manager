package users

import (
  "database/sql"
  "fmt"
  mock "mock/plugins"
  "os"
  "plugins/db"
  "processing"
  "testing"
  . "utils"
)

var (
  usersDatabase *sql.DB
  usersData UsersDataPlugin
)

func init() {
  var err error
  usersDatabase, err = sql.Open(
    "sqlite3", os.Getenv("TEST_DB_FILE"))
  if err != nil {
    fmt.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  usersData = NewUsersDataPlugin(usersDatabase)
  execUsersQuery(mock.TurnOnForeignKeys)
  createGroupsForUsers()
  createUsersTable()
  deleteAllFromUsers()
}

func execUsersQuery(q string, args ...interface{}) {
  db.ExecQuery(usersDatabase, q, args...)
}

func createGroupsForUsers() {
  db.CreateGroups(usersDatabase)
}

func createUsersTable() {
  execUsersQuery(mock.CreateUsersTable)
}

func deleteUsersTable() {
  execUsersQuery(mock.DropUsersTable)
}

func createTestUsers() {
  for _, q := range mock.TestingUsersQueries {
    execUsersQuery(q)
  }
}

func deleteAllFromUsers() {
  execUsersQuery(mock.ClearUsersTable)
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
    t.Log(GetExpectationString(mock.TestingUsers, users))
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
    t.Log(GetExpectationString(mock.TestingUsers[0], user))
    t.Fail()
  })
}

func TestGetUserErrorIdNotExists(t *testing.T) {
  createTestUsers()
  defer deleteAllFromUsers()

  user, err := usersData.GetUser(11)
  Assert(err == processing.WorkerIdNotExists, func() {
    t.Log(GetExpectationString(processing.WorkerIdNotExists, err))
    t.Fail()
  })
  Assert(user == mock.EmptyUser, func() {
    t.Log(GetExpectationString(mock.EmptyUser, user))
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
    t.Log(GetExpectationString(mock.EmptyUser, user))
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
    t.Log(GetExpectationString(1, createdUserId))
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
    t.Log(GetExpectationString(0, createdUserId))
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
    t.Log(GetExpectationString(0, createdUserId))
    t.Fail()
  })
}

func TestCreateUserErrorGroupNotExists(t *testing.T) {
  deleteUsersTable()
  createUsersTable()
  defer deleteAllFromUsers()

  createdUserId, err := usersData.CreateUser(mock.TestingUserWithNotExistsGroupId)
  Assert(err == processing.WorkGroupNotExists, func() {
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
