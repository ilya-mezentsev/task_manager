package interfaces

import "models"

type (
  AdminData interface {
    GetAllGroups() ([]models.Group, error)
    CreateWorkGroup(groupName string) error
    DeleteWorkGroup(groupId uint) error
    GetAllUsers() ([]models.User, error)
    CreateUser(user models.User) error
    DeleteUser(userId uint) error
    GetAllTasks() ([]models.Task, error)
    AssignTasksToGroup(groupId uint, tasks []models.Task) error
    DeleteTask(taskId uint) error
  }

  GroupLeadData interface {
    AssignTaskToWorker(workerId uint, task models.Task) error
    GetTasksByGroupId(groupId uint) ([]models.Task, error)
  }

  GroupWorkerData interface {
    GetTasksByUserId(userId uint) ([]models.Task, error)
    AddCommentToTask(taskId uint, comment string) error
    MarkTaskAsCompleted(taskId uint) error
  }
)
