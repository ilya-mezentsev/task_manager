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
    WorkerId uint `json:"user_id"`
    Task `json:"task"`
  }

  GroupWorkerTasksRequest struct {
    WorkerId uint `json:"user_id"`
  }

  CommentTaskRequest struct {
    TaskId uint `json:"task_id"`
    Comment string `json:"comment"`
  }

  DeleteTaskRequest struct {
    TaskId uint `json:"task_id"`
  }

  CompleteTaskRequest DeleteTaskRequest
)
