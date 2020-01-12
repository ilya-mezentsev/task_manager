package group_lead

import (
  "interfaces"
  "models"
  "users"
)

type GroupLead struct {
  dataProvider interfaces.GroupLeadData
}

func NewGroupLead(provider interfaces.GroupLeadData) GroupLead {
  return GroupLead{dataProvider: provider}
}

func (gl GroupLead) AssignTaskToWorker(workerId uint, task models.Task) error {
  if err := gl.dataProvider.AssignTaskToWorker(workerId, task); err != nil {
    return users.ParseError("AssignTaskToWorker", err)
  }

  return nil
}

func (gl GroupLead) GetTasksByGroupId(groupId uint) ([]models.Task, error) {
  tasks, err := gl.dataProvider.GetTasksByGroupId(groupId)
  if err != nil {
    return nil, users.ParseError("GetTasksByGroupId", err)
  }

  return tasks, nil
}

func (gl GroupLead) GetUsersByGroupId(groupId uint) ([]models.User, error) {
  usersByGroupId, err := gl.dataProvider.GetUsersByGroupId(groupId)
  if err != nil {
    return nil, users.ParseError("GetUsersByGroupId", err)
  }

  return usersByGroupId, nil
}
