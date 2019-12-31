package mock

import (
  "errors"
  "models"
)

const (
  UserNameAlreadyExists = "u_already_exists"
  WorkGroupAlreadyExists = "wg_already_exists"
  WorkGroupNotExists = "wg_not_exists"
  WgIdNotExists uint = 0
)

var (
  TestWgName = "test_wg"
  WorkGroupId uint = 0
)

type AdminDataMock struct {
  Users []models.User
  WorkGroups map[uint]string
  Tasks map[uint][]models.Task
}

func (m *AdminDataMock) CreateUser(user models.User) error {
  if user.Name == UserNameAlreadyExists {
    return errors.New(UserNameAlreadyExists)
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
    return errors.New(WorkGroupAlreadyExists)
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
    return errors.New(WorkGroupNotExists)
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
