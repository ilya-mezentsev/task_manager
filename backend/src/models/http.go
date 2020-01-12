package models

type (
  CreateWorkGroupRequest struct {
    GroupName string `json:"group_name"`
  }

  DeleteWorkGroupRequest struct {
    GroupId uint `json:"group_id"`
  }

  CreateUserRequest struct {
    User User `json:"user"`
  }

  DeleteUserRequest struct {
    UserId uint `json:"user_id"`
  }

  AssignTasksToWorkGroupRequest struct {
    GroupId uint `json:"group_id"`
    Tasks []Task `json:"tasks"`
  }

  WorkGroupTasksRequest struct {
    GroupId uint `json:"group_id"`
  }

  AssignTaskToGroupWorkerRequest struct {
    WorkerId uint `json:"worker_id"`
    Task `json:"task"`
  }

  DeleteTaskRequest struct {
    TaskId uint `json:"task_id"`
  }
)
