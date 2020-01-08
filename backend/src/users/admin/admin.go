package admin

import (
  "fmt"
  "interfaces"
  "models"
  "users"
  "utils"
)

type Admin struct {
  dataProvider interfaces.AdminData
}

func NewAdmin(provider interfaces.AdminData) Admin {
  return Admin{dataProvider: provider}
}

func (a Admin) GetAllGroups() ([]models.Group, error) {
  allGroups, err := a.dataProvider.GetAllGroups()
  if err != nil {
    return nil, users.ParseError("GetAllGroups", err)
  }

  return allGroups, nil
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
  allUsers, err := a.dataProvider.GetAllUsers()
  if err != nil {
    return nil, users.ParseError("GetAllUsers", err)
  }

  return allUsers, nil
}

func (a Admin) CreateUser(user models.User) error {
  user.Password = a.getUserPassword(user)

  if err := a.dataProvider.CreateUser(user); err != nil {
    return users.ParseError("CreateUser", err)
  }

  return nil
}

func (a Admin) getUserPassword(user models.User) string {
  return utils.GetHash(fmt.Sprintf("%s_%d", user.Name, user.GroupId))
}

func (a Admin) DeleteUser(userId uint) error {
  if err := a.dataProvider.DeleteUser(userId); err != nil {
    return users.ParseError("DeleteUser", err)
  }

  return nil
}

func (a Admin) GetAllTasks() ([]models.Task, error) {
  allTasks, err := a.dataProvider.GetAllTasks()
  if err != nil {
    return nil, users.ParseError("GetAllTasks", err)
  }

  return allTasks, nil
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
