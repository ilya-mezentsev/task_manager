package admin

import (
  "interfaces"
  "models"
)

type Admin struct {
  dataProvider interfaces.AdminData
}

func NewAdmin(provider interfaces.AdminData) Admin {
  return Admin{dataProvider: provider}
}

func (a Admin) CreateUser(user models.User) error {
  if err := a.dataProvider.CreateUser(user); err != nil {
    return ParseAdminError("CreateUser", err)
  }

  return nil
}

func (a Admin) CreateWorkGroup(groupName string) error {
  if err := a.dataProvider.CreateWorkGroup(groupName); err != nil {
    return ParseAdminError("CreateWorkGroup", err)
  }

  return nil
}

func (a Admin) AssignTasksToWorkGroup(groupId uint, tasks []models.Task) error {
  if err := a.dataProvider.AssignTasksToGroup(groupId, tasks); err != nil {
    return ParseAdminError("AssignTasksToWorkGroup", err)
  }

  return nil
}