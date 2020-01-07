package api

import (
  "interfaces"
  "net/http"
  . "users/admin"
)

var adminRequestHandler AdminRequestHandler

type AdminRequestHandler struct {
  admin Admin
}

func InitAdminRequestHandler(adminDataPlugin interfaces.AdminData) {
  adminRequestHandler.admin = NewAdmin(adminDataPlugin)
}

func BindAdminRoutesToHandlers() {
  api := router.PathPrefix("/api/admin").Subrouter()

  api.HandleFunc("groups", adminRequestHandler.GetAllTasks).Methods(http.MethodGet)
  api.HandleFunc("group", adminRequestHandler.CreateGroup).Methods(http.MethodPost)
  api.HandleFunc("group", adminRequestHandler.DeleteGroup).Methods(http.MethodDelete)

  api.HandleFunc("users", adminRequestHandler.GetAllUsers).Methods(http.MethodGet)
  api.HandleFunc("user", adminRequestHandler.CreateUser).Methods(http.MethodPost)
  api.HandleFunc("user", adminRequestHandler.DeleteUser).Methods(http.MethodDelete)

  api.HandleFunc("tasks", adminRequestHandler.GetAllTasks).Methods(http.MethodGet)
  api.HandleFunc("tasks", adminRequestHandler.AssignTasksToWorkGroup).Methods(http.MethodPost)
  api.HandleFunc("task", adminRequestHandler.DeleteTask).Methods(http.MethodDelete)
}

func (handler AdminRequestHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {}

func (handler AdminRequestHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {}

func (handler AdminRequestHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {}

func (handler AdminRequestHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {}

func (handler AdminRequestHandler) CreateUser(w http.ResponseWriter, r *http.Request) {}

func (handler AdminRequestHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {}

func (handler AdminRequestHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {}

func (handler AdminRequestHandler) AssignTasksToWorkGroup(w http.ResponseWriter, r *http.Request) {}

func (handler AdminRequestHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {}
