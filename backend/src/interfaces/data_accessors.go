package interfaces

import "models"

type (
  AdminData interface {
    CreateUser(user models.User) error
    CreateWorkGroup(groupName string) error
    AssignTasksToGroup(groupId uint, tasks []models.Task) error
  }

  GroupLeadData interface {
    AssignTaskToWorker(task models.Task, workerId uint) error
  }

  GroupWorkerData interface {
    AddCommentToTask(taskId uint, comment string) error
    MarkTaskAsCompleted(taskId uint) error
  }
)
