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

func TestAdmin_AdminCreateUserSuccessfully(t *testing.T) {
  testUser := models.User{Password: GetHash("_0")}
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

func TestAdmin_AdminGetAllUsers(t *testing.T) {
  users, err := admin.GetAllUsers()

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(len(users) == len(mockAdminData.Users), func() {
    t.Log(GetExpectationString(mockAdminData.Users, users))
    t.Fail()
  })
}

func TestAdmin_AdminGetAllUsersError(t *testing.T) {
  mockAdminData.TurnOnReturningErrorOnGettingAll()
  defer mockAdminData.TurnOffReturningErrorOnGettingAll()

  users, err := admin.GetAllUsers()

  AssertErrorsEqual(err, mock.UnableToGetAllUsersInternal, func() {
    t.Log(GetExpectationString(mock.UnableToGetAllUsersInternal, err))
    t.Fail()
  })
  Assert(users == nil, func() {
    t.Log(GetExpectationString(nil, users))
    t.Fail()
  })
}

func TestAdmin_AdminCreateUserErrorByUserNameExists(t *testing.T) {
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

func TestAdmin_AdminCreateUserInternalError(t *testing.T) {
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

func TestAdmin_AdminDeleteUserSuccess(t *testing.T) {
  err := admin.DeleteUser(2)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
}

func TestAdmin_AdminDeleteUserErrorIdNotExists(t *testing.T) {
  err := admin.DeleteUser(mock.UserIdNotExists)

  AssertErrorsEqual(err, mock.UnableToDeleteUserIdNotExists, func() {
    t.Log(GetExpectationString(mock.UnableToDeleteUserIdNotExists, err))
    t.Fail()
  })
}

func TestAdmin_AdminDeleteUserInternalError(t *testing.T) {
  err := admin.DeleteUser(mock.UserIdDeletingError)

  AssertErrorsEqual(err, mock.UnableToDeleteUserInternal, func() {
    t.Log(GetExpectationString(mock.UnableToDeleteUserInternal, err))
    t.Fail()
  })
}

func TestAdmin_AdminGetAllGroupsSuccess(t *testing.T) {
  groups, err := admin.GetAllGroups()

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(len(groups) == len(mockAdminData.WorkGroups), func() {
    t.Log(GetExpectationString(len(mockAdminData.WorkGroups), len(groups)))
    t.Fail()
  })
}

func TestAdmin_AdminGetAllGroupsError(t *testing.T) {
  mockAdminData.TurnOnReturningErrorOnGettingAll()
  defer mockAdminData.TurnOffReturningErrorOnGettingAll()

  groups, err := admin.GetAllGroups()

  AssertErrorsEqual(err, mock.UnableToGetAllGroupsInternal, func() {
    t.Log(GetExpectationString(mock.UnableToGetAllGroupsInternal, err))
    t.Fail()
  })
  Assert(groups == nil, func() {
    t.Log(GetExpectationString(nil, groups))
    t.Fail()
  })
}

func TestAdmin_AdminCreateWorkGroupSuccess(t *testing.T) {
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

func TestAdmin_AdminCreateWorkGroupErrorGroupAlreadyExists(t *testing.T) {
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

func TestAdmin_AdminCreateWorkGroupInternalError(t *testing.T) {
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

func TestAdmin_AdminDeleteWorkGroupSuccess(t *testing.T) {
  err := admin.DeleteWorkGroup(2)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
}

func TestAdmin_AdminDeleteWorkGroupErrorIdNotExists(t *testing.T) {
  err := admin.DeleteWorkGroup(mock.WgIdNotExists)

  AssertErrorsEqual(err, mock.UnableToDeleteWgIdNotExists, func() {
    t.Log(GetExpectationString(mock.UnableToDeleteWgIdNotExists, err))
    t.Fail()
  })
}

func TestAdmin_AdminGetAllTasksSuccess(t *testing.T) {
  tasks, err := admin.GetAllTasks()

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  expectedTasks, _ := mockAdminData.GetAllTasks()
  Assert(len(tasks) == len(expectedTasks), func() {
    t.Log(GetExpectationString(expectedTasks, tasks))
    t.Fail()
  })
}

func TestAdmin_AdminGetAllTasksError(t *testing.T) {
  mockAdminData.TurnOnReturningErrorOnGettingAll()
  defer mockAdminData.TurnOffReturningErrorOnGettingAll()

  tasks, err := admin.GetAllTasks()

  AssertErrorsEqual(err, mock.UnableToGetAllTasksInternal, func() {
    t.Log(GetExpectationString(mock.UnableToGetAllTasksInternal, err))
    t.Fail()
  })
  Assert(tasks == nil, func() {
    t.Log(GetExpectationString(nil, tasks))
    t.Fail()
  })
}

func TestAdmin_AdminAssignTasksToWorkGroup(t *testing.T) {
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

func TestAdmin_AdminAssignTasksErrorByNotExistsWrkGroup(t *testing.T) {
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

func TestAdmin_AdminAssignTasksInternalError(t *testing.T) {
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

func TestAdmin_AdminDeleteTaskSuccess(t *testing.T) {
  err := admin.DeleteTask(2)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
}

func TestAdmin_AdminDeleteTaskErrorIdNotExists(t *testing.T) {
  err := admin.DeleteTask(mock.TaskIdNotExists)

  AssertErrorsEqual(err, mock.UnableToDeleteTaskIdNotExists, func() {
    t.Log(GetExpectationString(mock.UnableToDeleteTaskIdNotExists, err))
    t.Fail()
  })
}

func TestAdmin_AdminDeleteTaskInternalError(t *testing.T) {
  err := admin.DeleteTask(mock.TaskIdDeletingError)

  AssertErrorsEqual(err, mock.UnableToDeleteTaskInternal, func() {
    t.Log(GetExpectationString(mock.UnableToDeleteTaskInternal, err))
    t.Fail()
  })
}
