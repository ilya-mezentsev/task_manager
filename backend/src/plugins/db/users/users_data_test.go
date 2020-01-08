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
  dbFile string
  database *sql.DB
  usersData DataPlugin
)

func init() {
  dbFile = os.Getenv("TEST_DB_FILE")
  if dbFile == "" {
    fmt.Println("TEST_DB_FILE env var is not set")
    os.Exit(1)
  }

  var err error
  database, err = sql.Open("sqlite3", dbFile)
  if err != nil {
    fmt.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  usersData = NewDataPlugin(database)
  db.ExecQuery(database, mock.TurnOnForeignKeys)
  db.CreateGroups(database)
  initUsersTable()
}

func dropUsersTable() {
  db.ExecQuery(database, mock.DropUsersTable)
}

func initUsersTable() {
  dropUsersTable()
  db.ExecQuery(database, mock.CreateUsersTable)
  for _, q := range mock.TestingUsersQueries {
    db.ExecQuery(database, q)
  }
}

func TestGetAllUsersSuccess(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

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
  dropUsersTable()

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
  initUsersTable()
  defer dropUsersTable()

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
  initUsersTable()
  defer dropUsersTable()

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
  initUsersTable()
  defer dropUsersTable()

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
  initUsersTable()
  defer dropUsersTable()

  createdUserId, err := usersData.CreateUser(mock.TestingUser)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(createdUserId == uint(len(mock.TestingUsers)+1), func() {
    t.Log(GetExpectationString(uint(len(mock.TestingUsers)+1), createdUserId))
    t.Fail()
  })
}

func TestCreateUserErrorTableNotExists(t *testing.T) {
  dropUsersTable()

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
  initUsersTable()
  defer dropUsersTable()

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
  initUsersTable()
  defer dropUsersTable()

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
  initUsersTable()
  defer dropUsersTable()

  err := usersData.DeleteUser(1)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
}

func TestDeleteUserErrorTableNotExists(t *testing.T) {
  dropUsersTable()

  err := usersData.DeleteUser(1)
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestDeleteUserErrorNotExists(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

  err := usersData.DeleteUser(11)
  AssertErrorsEqual(err, processing.UserIdNotExists, func() {
    t.Log("should be error")
    t.Fail()
  })
}
