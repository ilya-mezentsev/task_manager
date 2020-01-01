package mock

import (
  "errors"
  "models"
  "processing"
)

const (
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

  UnableToCreateUserInternal = errors.New("unable to create user: internal error")
  UnableToCreateUserNameAlreadyExists = errors.New("unable to create user: user name already exists")

  UnableToCreateWgInternal = errors.New("unable to create work group: internal error")
  UnableToCreateWgAlreadyExists = errors.New("unable to create work group: work group already exists")

  UnableToAssignTasksInternal = errors.New("unable to assign tasks: internal error")
  UnableToAssignTasksNotExists = errors.New("unable to assign tasks: work group not exists")
)

type AdminDataMock struct {
  Users []models.User
  WorkGroups map[uint]string
  Tasks map[uint][]models.Task
}

func (m *AdminDataMock) CreateUser(user models.User) error {
  if user.Name == UserNameAlreadyExists {
    return processing.UserNameAlreadyExists
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

func (m *AdminDataMock) CreateWorkGroup(groupName string) error {
  if groupName == WorkGroupAlreadyExists {
    return processing.WorkGroupAlreadyExists
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

func (m *AdminDataMock) AssignTasksToGroup(groupId uint, tasks []models.Task) error {
  if groupId == WgIdNotExists {
    return processing.WorkGroupNotExists
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
