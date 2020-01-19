package api

import (
  "api/helpers"
  "api/middleware"
  "interfaces"
  "models"
  "net/http"
  . "users/admin"
)

var adminRequestHandler AdminRequestHandler

type AdminRequestHandler struct {
  admin Admin
  checker helpers.InputChecker
}

func InitAdminRequestHandler(adminDataPlugin interfaces.AdminData) {
  adminRequestHandler = AdminRequestHandler{
    admin: NewAdmin(adminDataPlugin),
    checker: helpers.NewInputChecker(),
  }
  bindAdminRoutesToHandlers()
}

func bindAdminRoutesToHandlers() {
  api := router.PathPrefix("/admin").Subrouter()
  api.Use(middleware.RequiredAdminRole)

  api.HandleFunc("/groups", adminRequestHandler.GetAllGroups).Methods(http.MethodGet)
  api.HandleFunc("/group", adminRequestHandler.CreateGroup).Methods(http.MethodPost)
  api.HandleFunc("/group", adminRequestHandler.DeleteGroup).Methods(http.MethodDelete)

  api.HandleFunc("/users", adminRequestHandler.GetAllUsers).Methods(http.MethodGet)
  api.HandleFunc("/user", adminRequestHandler.CreateUser).Methods(http.MethodPost)
  api.HandleFunc("/user", adminRequestHandler.DeleteUser).Methods(http.MethodDelete)

  api.HandleFunc("/tasks", adminRequestHandler.GetAllTasks).Methods(http.MethodGet)
  api.HandleFunc("/tasks", adminRequestHandler.AssignTasksToWorkGroup).Methods(http.MethodPost)
  api.HandleFunc("/task", adminRequestHandler.DeleteTask).Methods(http.MethodDelete)
}

func (handler AdminRequestHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  groups, err := handler.admin.GetAllGroups()
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, groups)
}

func (handler AdminRequestHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var createGroupReq models.CreateWorkGroupRequest
  decodeRequestBody(r, &createGroupReq)

  if !handler.checker.IsStringCorrect(createGroupReq.GroupName) {
    panic(getIncorrectGroupNameError(createGroupReq.GroupName))
  }

  err := handler.admin.CreateWorkGroup(createGroupReq.GroupName)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, nil)
}

func (handler AdminRequestHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var deleteGroupReq models.DeleteWorkGroupRequest
  decodeRequestBody(r, &deleteGroupReq)

  if !handler.checker.IsSafeUint64(deleteGroupReq.GroupId) {
    panic(getIncorrectGroupIdError(deleteGroupReq.GroupId))
  }

  err := handler.admin.DeleteWorkGroup(deleteGroupReq.GroupId)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, nil)
}

func (handler AdminRequestHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  users, err := handler.admin.GetAllUsers()
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, users)
}

func (handler AdminRequestHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var createUserReq models.CreateUserRequest
  decodeRequestBody(r, &createUserReq)

  groupId, userName := createUserReq.User.GroupId, createUserReq.User.Name
  switch {
  case !handler.checker.IsSafeUint64(groupId):
    panic(getIncorrectUserGroupIdError(groupId))
  case !handler.checker.IsStringCorrect(userName):
    panic(getIncorrectUserNameError(userName))
  }

  err := handler.admin.CreateUser(createUserReq.User)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, nil)
}

func (handler AdminRequestHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var deleteUserReq models.DeleteUserRequest
  decodeRequestBody(r, &deleteUserReq)

  if !handler.checker.IsSafeUint64(deleteUserReq.UserId) {
    panic(getIncorrectUserIdError(deleteUserReq.UserId))
  }

  err := handler.admin.DeleteUser(deleteUserReq.UserId)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, nil)
}

func (handler AdminRequestHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  tasks, err := handler.admin.GetAllTasks()
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, tasks)
}

func (handler AdminRequestHandler) AssignTasksToWorkGroup(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var assignTasksReq models.AssignTasksToWorkGroupRequest
  decodeRequestBody(r, &assignTasksReq)

  if !handler.checker.IsSafeUint64(assignTasksReq.GroupId) {
    panic(getIncorrectGroupIdError(assignTasksReq.GroupId))
  }
  for _, task := range assignTasksReq.Tasks {
    if !handler.checker.IsStringCorrect(task.Title) {
      panic(getIncorrectTaskTitleError(task.Title))
    } else if !handler.checker.IsLongTextCorrect(task.Description) {
      panic(getIncorrectTaskDescriptionError(task.Description))
    }
  }

  err := handler.admin.AssignTasksToWorkGroup(assignTasksReq.GroupId, assignTasksReq.Tasks)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, nil)
}

func (handler AdminRequestHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var deleteTaskReq models.DeleteTaskRequest
  decodeRequestBody(r, &deleteTaskReq)

  if !handler.checker.IsSafeUint64(deleteTaskReq.TaskId) {
    panic(getIncorrectTaskIdError(deleteTaskReq.TaskId))
  }

  err := handler.admin.DeleteTask(deleteTaskReq.TaskId)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, nil)
}
