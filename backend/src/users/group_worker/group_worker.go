package group_worker

import (
  "interfaces"
  "users"
)

type GroupWorker struct {
  dataProvider interfaces.GroupWorkerData
}

func NewGroupWorker(provider interfaces.GroupWorkerData) GroupWorker {
  return GroupWorker{dataProvider: provider}
}

func (gw GroupWorker) AddCommentToTask(taskId uint, comment string) error {
  if err := gw.dataProvider.AddCommentToTask(taskId, comment); err != nil {
    return users.ParseError("AddCommentToTask", err)
  }

  return nil
}

func (gw GroupWorker) MarkTaskAsCompleted(taskId uint) error {
  if err := gw.dataProvider.MarkTaskAsCompleted(taskId); err != nil {
    return users.ParseError("MarkTaskAsCompleted", err)
  }

  return nil
}
