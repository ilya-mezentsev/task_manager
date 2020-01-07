package api

import (
  "interfaces"
  "net/http"
  . "users/group_worker"
)

var groupWorkerRequestHandler GroupWorkerRequestHandler

type GroupWorkerRequestHandler struct {
  groupWorker GroupWorker
}

func InitGroupWorkerRequestHandler(groupWorkerDataPlugin interfaces.GroupWorkerData) {
  groupWorkerRequestHandler.groupWorker = NewGroupWorker(groupWorkerDataPlugin)
}

func BindGroupWorkerRoutesToHandlers() {
  api := router.PathPrefix("/api/group/worker").Subrouter()

  api.HandleFunc("tasks", groupWorkerRequestHandler.GetAllTasks).Methods(http.MethodGet)
  api.HandleFunc("task/comment", groupWorkerRequestHandler.CommentTask).Methods(http.MethodPatch)
  api.HandleFunc("task/complete", groupWorkerRequestHandler.CompleteTask).Methods(http.MethodPatch)
}

func (handler GroupWorkerRequestHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {}

func (handler GroupWorkerRequestHandler) CommentTask(w http.ResponseWriter, r *http.Request) {}

func (handler GroupWorkerRequestHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {}
