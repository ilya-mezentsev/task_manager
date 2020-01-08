package mock

import (
  "errors"
  "models"
  "plugins/db"
)

const (
  TaskIdNotExists uint = 0
  TaskIdDeletingError uint = 1
  TaskIdDeletingErrorMessage  = "t_id_deleting_error"
  UserIdNotExists uint = 0
  UserIdDeletingError uint = 1
  UserIdDeletingErrorMessage = "u_id_deleting_error"
  UserNameAlreadyExists = "u_already_exists"
  UserNameCreationError = "u_creation_error"
  WorkGroupAlreadyExists = "wg_already_exists"
  WorkGroupCreationError = "wg_creation_error"
  WorkGroupAssigningError = "wg_assigning_error"
  WgIdNotExists uint = 0
  WgIdAssigningError uint = 1
)

var (
  TestWgName = "test_wg"
  WorkGroupId uint = 1

  gettingAllError = errors.New("error")

  UnableToGetAllUsersInternal = errors.New("unable to get all users: internal error")
  UnableToCreateUserInternal = errors.New("unable to create user: internal error")
  UnableToCreateUserNameAlreadyExists = errors.New("unable to create user: user name already exists")
  UnableToDeleteUserIdNotExists = errors.New("unable to delete user: user id not exists")
  UnableToDeleteUserInternal = errors.New("unable to delete user: internal error")

  UnableToGetAllGroupsInternal = errors.New("unable to get all groups: internal error")
  UnableToCreateWgInternal = errors.New("unable to create work group: internal error")
  UnableToCreateWgAlreadyExists = errors.New("unable to create work group: work group already exists")
  UnableToDeleteWgIdNotExists = errors.New("unable to delete work group: work group not exists")

  UnableToGetAllTasksInternal = errors.New("unable to get all tasks: internal error")
  UnableToAssignTasksInternal = errors.New("unable to assign tasks: internal error")
  UnableToAssignTasksNotExists = errors.New("unable to assign tasks: work group not exists")
  UnableToDeleteTaskIdNotExists = errors.New("unable to delete task: task id not exists")
  UnableToDeleteTaskInternal = errors.New("unable to delete task: internal error")
)

type AdminDataMock struct {
  Users []models.User
  WorkGroups map[uint]string
  Tasks map[uint][]models.Task
  gettingAllReturnsError bool
}

func (m *AdminDataMock) TurnOnReturningErrorOnGettingAll() {
  m.gettingAllReturnsError = true
}

func (m *AdminDataMock) TurnOffReturningErrorOnGettingAll() {
  m.gettingAllReturnsError = false
}

func (m *AdminDataMock) GetAllGroups() ([]models.Group, error) {
  if m.gettingAllReturnsError {
    return nil, gettingAllError
  }

  var allGroups []models.Group
  for groupId, groupName := range m.WorkGroups {
    allGroups = append(allGroups, models.Group{
      ID: groupId,
      Name: groupName,
    })
  }

  return allGroups, nil
}

func (m *AdminDataMock) CreateWorkGroup(groupName string) error {
  if groupName == WorkGroupAlreadyExists {
    return db.WorkGroupAlreadyExists
  } else if groupName == WorkGroupCreationError {
    return errors.New(WorkGroupCreationError)
  }

  WorkGroupId++
  m.WorkGroups[WorkGroupId] = groupName
  return nil
}

func (m AdminDataMock) HasWorkGroup(groupName string) bool {
  for _, name := range m.WorkGroups {
    if name == groupName {
      return true
    }
  }

  return false
}

func (m *AdminDataMock) DeleteWorkGroup(groupId uint) error {
  if groupId == WgIdNotExists {
    return db.WorkGroupNotExists
  }

  return nil
}

func (m *AdminDataMock) GetAllUsers() ([]models.User, error) {
  if m.gettingAllReturnsError {
    return nil, gettingAllError
  }

  return m.Users, nil
}

func (m *AdminDataMock) CreateUser(user models.User) error {
  if user.Name == UserNameAlreadyExists {
    return db.UserNameAlreadyExists
  } else if user.Name == UserNameCreationError {
    return errors.New(UserNameCreationError)
  }

  m.Users = append(m.Users, user)
  return nil
}

func (m AdminDataMock) HasUser(user models.User) bool {
  for _, u := range m.Users {
    if u == user {
      return true
    }
  }

  return false
}

func (m *AdminDataMock) DeleteUser(userId uint) error {
  if userId == UserIdNotExists {
    return db.UserIdNotExists
  } else if userId == UserIdDeletingError {
    return errors.New(UserIdDeletingErrorMessage)
  }

  var users []models.User
  for _, u := range m.Users {
    if u.ID != userId {
      users = append(users, u)
    }
  }
  m.Users = users

  return nil
}

func (m *AdminDataMock) GetAllTasks() ([]models.Task, error) {
  if m.gettingAllReturnsError {
    return nil, gettingAllError
  }

  var allTasks []models.Task

  for _, t := range m.Tasks {
    allTasks = append(allTasks, t...)
  }

  return allTasks, nil
}

func (m *AdminDataMock) AssignTasksToGroup(groupId uint, tasks []models.Task) error {
  if groupId == WgIdNotExists {
    return db.WorkGroupNotExists
  } else if groupId == WgIdAssigningError {
    return errors.New(WorkGroupAssigningError)
  }

  m.Tasks[groupId] = tasks
  return nil
}

func (m AdminDataMock) TasksAssigned(groupId uint, tasks []models.Task) bool {
  groupTasks, ok := m.Tasks[groupId]
  if len(m.Tasks) != len(tasks) || !ok {
    return false
  }

  for taskIndex, _ := range groupTasks {
    if groupTasks[taskIndex] != tasks[taskIndex] {
      return false
    }
  }

  return true
}

func (m *AdminDataMock) DeleteTask(taskId uint) error {
  if taskId == TaskIdNotExists {
    return db.TaskIdNotExists
  } else if taskId == TaskIdDeletingError {
    return errors.New(TaskIdDeletingErrorMessage)
  }

  var allTasks []models.Task
  for _, tasks := range m.Tasks {
    for _, task := range tasks {
      if task.ID != taskId {
        allTasks = append(allTasks, task)
      }
    }
  }

  return nil
}
