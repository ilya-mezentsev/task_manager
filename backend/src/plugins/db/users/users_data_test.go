package users

import (
  "database/sql"
  "fmt"
  mock "mock/plugins"
  "models"
  "os"
  "plugins/db"
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

func TestUsersDataPlugin_GetAllUsersSuccess(t *testing.T) {
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

func TestUsersDataPlugin_GetAllUsersErrorTableNotExists(t *testing.T) {
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

func TestUsersDataPlugin_GetUserByCredentialsSuccess(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

  user, err := usersData.GetUserByCredentials(mock.TestingCredentials[0], mock.TestingCredentials[1])
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(user == mock.TestingUsers[0], func() {
    t.Log(GetExpectationString(mock.TestingUsers[0], user))
    t.Fail()
  })
}

func TestUsersDataPlugin_GetUserByCredentialsErrorTableNotExists(t *testing.T) {
  dropUsersTable()

  user, err := usersData.GetUserByCredentials(mock.TestingCredentials[0], mock.TestingCredentials[1])
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
  Assert(user == models.User{}, func() {
    t.Log(GetExpectationString(models.User{}, user))
    t.Fail()
  })
}

func TestUsersDataPlugin_GetUserByCredentialsErrorUserNotExists(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

  user, err := usersData.GetUserByCredentials("", "")
  AssertErrorsEqual(err, db.UserNotFoundByCredentials, func() {
    t.Log(GetExpectationString(db.UserNotFoundByCredentials, err))
    t.Fail()
  })
  Assert(user == models.User{}, func() {
    t.Log(GetExpectationString(models.User{}, user))
    t.Fail()
  })
}

func TestUsersDataPlugin_GetUsersByGroupIdSuccess(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

  users, err := usersData.GetUsersByGroupId(2)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mock.UserListEqual(users, mock.TestingUsersByGroupId), func() {
    t.Log(GetExpectationString(mock.TestingUsersByGroupId, users))
    t.Fail()
  })
}

func TestUsersDataPlugin_GetUsersByNotExistsGroupId(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

  users, err := usersData.GetUsersByGroupId(11)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mock.UserListEqual(users, nil), func() {
    t.Log(GetExpectationString(nil, users))
    t.Fail()
  })
}

func TestUsersDataPlugin_GetUserSuccess(t *testing.T) {
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

func TestUsersDataPlugin_GetUserErrorIdNotExists(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

  user, err := usersData.GetUser(11)
  Assert(err == db.WorkerIdNotExists, func() {
    t.Log(GetExpectationString(db.WorkerIdNotExists, err))
    t.Fail()
  })
  Assert(user == mock.EmptyUser, func() {
    t.Log(GetExpectationString(mock.EmptyUser, user))
    t.Fail()
  })
}

func TestUsersDataPlugin_GetUserErrorTableNotExists(t *testing.T) {
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

func TestUsersDataPlugin_CreateUserSuccess(t *testing.T) {
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

func TestUsersDataPlugin_CreateUserErrorTableNotExists(t *testing.T) {
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

func TestUsersDataPlugin_CreateUserErrorNameAlreadyExists(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

  createdUserId, err := usersData.CreateUser(mock.TestingUserWithExistsName)
  Assert(err == db.UserNameAlreadyExists, func() {
    t.Log("should be error")
    t.Fail()
  })
  Assert(createdUserId == 0, func() {
    t.Log(GetExpectationString(0, createdUserId))
    t.Fail()
  })
}

func TestUsersDataPlugin_CreateUserErrorGroupNotExists(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

  createdUserId, err := usersData.CreateUser(mock.TestingUserWithNotExistsGroupId)
  Assert(err == db.WorkGroupNotExists, func() {
    t.Log("should be error")
    t.Fail()
  })
  Assert(createdUserId == 0, func() {
    t.Log("unexpected created user id:", createdUserId)
    t.Log("wanted:", 0)
    t.Fail()
  })
}

func TestUsersDataPlugin_DeleteUserSuccess(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

  err := usersData.DeleteUser(1)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
}

func TestUsersDataPlugin_DeleteUserErrorTableNotExists(t *testing.T) {
  dropUsersTable()

  err := usersData.DeleteUser(1)
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestUsersDataPlugin_DeleteUserErrorNotExists(t *testing.T) {
  initUsersTable()
  defer dropUsersTable()

  err := usersData.DeleteUser(11)
  AssertErrorsEqual(err, db.UserIdNotExists, func() {
    t.Log("should be error")
    t.Fail()
  })
}
