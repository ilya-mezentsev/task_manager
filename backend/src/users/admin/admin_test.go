package admin

import (
  "errors"
  mock "mock/admin"
  "models"
  "testing"
  . "utils"
)

var (
  mockAdminData = mock.AdminDataMock{
    Users: []models.User{},
    WorkGroups: make(map[uint]string),
    Tasks: make(map[uint][]models.Task),
  }
  admin = NewAdmin(&mockAdminData)
)

func TestAdminCreateUserSuccessfully(t *testing.T) {
  testUser := models.User{}
  err := admin.CreateUser(testUser)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mockAdminData.HasUser(testUser), func() {
    t.Log("failed to create user")
    t.Fail()
  })
}

func TestAdminCreateUserErrorByUserNameExists(t *testing.T) {
  testUser := models.User{Name: mock.UserNameAlreadyExists}
  err := admin.CreateUser(testUser)

  Assert(errors.Is(err, UnableToCreateUser), func() {
    t.Log("wrong error:", err)
    t.Fail()
  })
  Assert(!mockAdminData.HasUser(testUser), func() {
    t.Log("user should not be created")
    t.Fail()
  })
}

func TestAdminCreateWorkGroupSuccess(t *testing.T) {
  err := admin.CreateWorkGroup(mock.TestWgName)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mockAdminData.HasWorkGroup(mock.TestWgName), func() {
    t.Log("work group should be created")
    t.Fail()
  })
}

func TestAdminCreateWorkGroupErrorGroupAlreadyExists(t *testing.T) {
  err := admin.CreateWorkGroup(mock.WorkGroupAlreadyExists)

  Assert(errors.Is(err, UnableToCreateWorkGroup), func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(!mockAdminData.HasWorkGroup(mock.WorkGroupAlreadyExists), func() {
    t.Log("work group should not be created")
    t.Fail()
  })
}

func TestAdminAssignTasksToWorkGroup(t *testing.T) {
  tasks := []models.Task{
    { Title: "", Description: "" },
  }
  err := admin.AssignTasksToWorkGroup(mock.WorkGroupId, tasks)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mockAdminData.TasksAssigned(mock.WorkGroupId, tasks), func() {
    t.Log("tasks should be assigned")
    t.Fail()
  })
}

func TestAdminAssignTasksErrorByNotExistsWrkGroup(t *testing.T) {
  tasks := []models.Task{
    { Title: "", Description: "" },
  }
  err := admin.AssignTasksToWorkGroup(mock.WgIdNotExists, tasks)

  Assert(errors.Is(err, UnableToAssignTasks), func() {
    t.Log("wrong error:", err)
    t.Fail()
  })
  Assert(!mockAdminData.TasksAssigned(mock.WgIdNotExists, tasks), func() {
    t.Log("tasks should not be assigned")
    t.Fail()
  })
}
