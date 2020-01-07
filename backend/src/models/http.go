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

  DeleteTaskRequest struct {
    TaskId uint `json:"task_id"`
  }
)
