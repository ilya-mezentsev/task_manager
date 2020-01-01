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
