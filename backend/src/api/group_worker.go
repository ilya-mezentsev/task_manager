package api

import (
  "api/helpers"
  "interfaces"
  "net/http"
  . "users/group_worker"
)

var groupWorkerRequestHandler GroupWorkerRequestHandler

type GroupWorkerRequestHandler struct {
  groupWorker GroupWorker
  checker helpers.InputChecker
}

func InitGroupWorkerRequestHandler(groupWorkerDataPlugin interfaces.GroupWorkerData) {
  groupWorkerRequestHandler.groupWorker = NewGroupWorker(groupWorkerDataPlugin)
  groupLeadRequestHandler.checker = helpers.NewInputChecker()
  bindGroupWorkerRoutesToHandlers()
}

func bindGroupWorkerRoutesToHandlers() {
  api := router.PathPrefix("/api/group/worker").Subrouter()

  api.HandleFunc("tasks", groupWorkerRequestHandler.GetAllTasks).Methods(http.MethodGet)
  api.HandleFunc("task/comment", groupWorkerRequestHandler.CommentTask).Methods(http.MethodPatch)
  api.HandleFunc("task/complete", groupWorkerRequestHandler.CompleteTask).Methods(http.MethodPatch)
}

func (handler GroupWorkerRequestHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {}

func (handler GroupWorkerRequestHandler) CommentTask(w http.ResponseWriter, r *http.Request) {}

func (handler GroupWorkerRequestHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {}
