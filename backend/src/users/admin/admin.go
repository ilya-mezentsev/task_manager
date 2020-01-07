package admin

import (
  "interfaces"
  "models"
  "users"
)

type Admin struct {
  dataProvider interfaces.AdminData
}

func NewAdmin(provider interfaces.AdminData) Admin {
  return Admin{dataProvider: provider}
}

func (a Admin) GetAllGroups() ([]models.Group, error) {
  return a.dataProvider.GetAllGroups()
}

func (a Admin) CreateWorkGroup(groupName string) error {
  if err := a.dataProvider.CreateWorkGroup(groupName); err != nil {
    return users.ParseError("CreateWorkGroup", err)
  }

  return nil
}

func (a Admin) DeleteWorkGroup(groupId uint) error {
  if err := a.dataProvider.DeleteWorkGroup(groupId); err != nil {
    return users.ParseError("DeleteWorkGroup", err)
  }

  return nil
}

func (a Admin) GetAllUsers() ([]models.User, error) {
  return a.dataProvider.GetAllUsers()
}

func (a Admin) CreateUser(user models.User) error {
  if err := a.dataProvider.CreateUser(user); err != nil {
    return users.ParseError("CreateUser", err)
  }

  return nil
}

func (a Admin) DeleteUser(userId uint) error {
  if err := a.dataProvider.DeleteUser(userId); err != nil {
    return users.ParseError("DeleteUser", err)
  }

  return nil
}

func (a Admin) GetAllTasks() ([]models.Task, error) {
  return a.dataProvider.GetAllTasks()
}

func (a Admin) AssignTasksToWorkGroup(groupId uint, tasks []models.Task) error {
  for _, task := range tasks {
    task.GroupId = groupId
  }

  if err := a.dataProvider.AssignTasksToGroup(groupId, tasks); err != nil {
    return users.ParseError("AssignTasksToWorkGroup", err)
  }

  return nil
}

func (a Admin) DeleteTask(taskId uint) error {
  if err := a.dataProvider.DeleteTask(taskId); err != nil {
    return users.ParseError("DeleteTask", err)
  }

  return nil
}
