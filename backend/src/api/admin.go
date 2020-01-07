package api

import (
  "interfaces"
  "models"
  "net/http"
  . "users/admin"
)

var adminRequestHandler AdminRequestHandler

type AdminRequestHandler struct {
  admin Admin
}

func InitAdminRequestHandler(adminDataPlugin interfaces.AdminData) {
  adminRequestHandler.admin = NewAdmin(adminDataPlugin)
  bindAdminRoutesToHandlers()
}

func bindAdminRoutesToHandlers() {
  api := router.PathPrefix("/api/admin").Subrouter()

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

  err := handler.admin.DeleteTask(deleteTaskReq.TaskId)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, nil)
}
