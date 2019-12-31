package interfaces

import "models"

type Admin interface {
  CreateUser(user models.User) error
  CreateWorkGroup(groupName string) error
  AssignTasksToGroup(tasks []models.Task, groupId uint) error
}

type GroupLead interface {
  AssignTaskToWorker(task models.Task, workerId uint) error
}

type GroupWorker interface {
  AddCommentToTask(taskId uint, comment string) error
  MarkTaskAsCompleted(taskId uint) error
}
