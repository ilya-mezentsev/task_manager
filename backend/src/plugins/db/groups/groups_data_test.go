package groups

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
  dbFile string
  database *sql.DB
  groupsData DataPlugin
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

  groupsData = NewDataPlugin(database)
  db.ExecQuery(database, mock.TurnOnForeignKeys)
}

func dropGroupsTable() {
  db.ExecQuery(database, mock.DropGroupsTable)
}

func initGroupsTable() {
  dropGroupsTable()
  db.ExecQuery(database, mock.CreateGroupsTable)
  for _, q := range mock.TestingGroupsQueries {
    db.ExecQuery(database, q)
  }
}

func TestGetAllTasksSuccess(t *testing.T) {
  initGroupsTable()
  defer dropGroupsTable()

  groups, err := groupsData.GetAllGroups()
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mock.GroupListEqual(groups, mock.TestingGroups), func() {
    t.Log(GetExpectationString(mock.TestingGroups, groups))
    t.Fail()
  })
}

func TestGetAllGroupsErrorTableNotExists(t *testing.T) {
  dropGroupsTable()

  groups, err := groupsData.GetAllGroups()
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
  Assert(groups == nil, func() {
    t.Log("should not be groups:", groups)
    t.Fail()
  })
}

func TestCreateWorkGroupSuccess(t *testing.T) {
  initGroupsTable()
  defer dropGroupsTable()

  err := groupsData.CreateWorkGroup(mock.TestingGroup.Name)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })

  groups, _ := groupsData.GetAllGroups()
  expectedGroups := append(mock.TestingGroups, mock.TestingGroup)
  Assert(mock.GroupListEqual(expectedGroups, groups), func() {
    t.Log(GetExpectationString(expectedGroups, groups))
    t.Fail()
  })
}

func TestCreateWorkGroupErrorTableNotExists(t *testing.T) {
  dropGroupsTable()

  err := groupsData.CreateWorkGroup("")
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestCreateWorkGroupErrorGroupNameAlreadyExists(t *testing.T) {
  initGroupsTable()
  defer dropGroupsTable()

  err := groupsData.CreateWorkGroup(mock.ExistsGroupName)
  AssertErrorsEqual(err, db.WorkGroupAlreadyExists, func() {
    t.Log(GetExpectationString(db.WorkGroupAlreadyExists, err))
    t.Fail()
  })
  groups, _ := groupsData.GetAllGroups()
  Assert(mock.GroupListEqual(mock.TestingGroups, groups), func() {
    t.Log(GetExpectationString(mock.TestingGroups, groups))
    t.Fail()
  })
}

func TestDeleteWorkGroupSuccess(t *testing.T) {
  initGroupsTable()
  defer dropGroupsTable()

  err := groupsData.DeleteWorkGroup(1)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  groups, _ := groupsData.GetAllGroups()
  expectedGroups := mock.TestingGroups[1:]
  Assert(mock.GroupListEqual(expectedGroups, groups), func() {
    t.Log(GetExpectationString(expectedGroups, groups))
    t.Fail()
  })
}

func TestDeleteWorkGroupErrorTableNotExists(t *testing.T) {
  dropGroupsTable()

  err := groupsData.DeleteWorkGroup(1)
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestDeleteWorkGroupErrorIdNotExists(t *testing.T) {
  initGroupsTable()
  defer dropGroupsTable()

  err := groupsData.DeleteWorkGroup(11)
  AssertErrorsEqual(err, db.WorkGroupNotExists, func() {
    t.Log(GetExpectationString(db.WorkGroupNotExists, err))
    t.Fail()
  })
}
