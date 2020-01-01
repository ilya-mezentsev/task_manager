package interfaces

import "models"

type (
  AdminData interface {
    CreateUser(user models.User) error
    CreateWorkGroup(groupName string) error
    AssignTasksToGroup(groupId uint, tasks []models.Task) error
  }

  GroupLeadData interface {
    AssignTaskToWorker(workerId uint, task models.Task) error
  }

  GroupWorkerData interface {
    AddCommentToTask(taskId uint, comment string) error
    MarkTaskAsCompleted(taskId uint) error
  }
)
