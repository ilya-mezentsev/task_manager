package admin

import (
  "mock"
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

  AssertErrorsEqual(err, mock.UnableToCreateUserNameAlreadyExists, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.UnableToCreateUserNameAlreadyExists)
    t.Fail()
  })
  Assert(!mockAdminData.HasUser(testUser), func() {
    t.Log("user should not be created")
    t.Fail()
  })
}

func TestAdminCreateUserInternalError(t *testing.T) {
  testUser := models.User{Name: mock.UserNameCreationError}
  err := admin.CreateUser(testUser)

  AssertErrorsEqual(err, mock.UnableToCreateUserInternal, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.UnableToCreateUserInternal)
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

  AssertErrorsEqual(err, mock.UnableToCreateWgAlreadyExists, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.UnableToCreateWgAlreadyExists)
    t.Fail()
  })
  Assert(!mockAdminData.HasWorkGroup(mock.WorkGroupAlreadyExists), func() {
    t.Log("work group should not be created")
    t.Fail()
  })
}

func TestAdminCreateWorkGroupInternalError(t *testing.T) {
  err := admin.CreateWorkGroup(mock.WorkGroupCreationError)

  AssertErrorsEqual(err, mock.UnableToCreateWgInternal, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.UnableToCreateWgAlreadyExists)
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

  AssertErrorsEqual(err, mock.UnableToAssignTasksNotExists, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.UnableToAssignTasksNotExists)
    t.Fail()
  })
  Assert(!mockAdminData.TasksAssigned(mock.WgIdNotExists, tasks), func() {
    t.Log("tasks should not be assigned")
    t.Fail()
  })
}

func TestAdminAssignTasksInternalError(t *testing.T) {
  tasks := []models.Task{
    { Title: "", Description: "" },
  }
  err := admin.AssignTasksToWorkGroup(mock.WgIdAssigningError, tasks)

  AssertErrorsEqual(err, mock.UnableToAssignTasksInternal, func() {
    t.Log("wrong error:", err)
    t.Log("should be:", mock.UnableToAssignTasksInternal)
    t.Fail()
  })
  Assert(!mockAdminData.TasksAssigned(mock.WgIdAssigningError, tasks), func() {
    t.Log("tasks should not be assigned")
    t.Fail()
  })
}
