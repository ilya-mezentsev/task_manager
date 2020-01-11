package api

import (
  "interfaces"
  "models"
  "net/http"
  . "users/group_lead"
)

var groupLeadRequestHandler GroupLeadRequestHandler

type GroupLeadRequestHandler struct {
  groupLead GroupLead
}

func InitGroupLeadRequestHandler(groupLeadDataPlugin interfaces.GroupLeadData) {
  groupLeadRequestHandler.groupLead = NewGroupLead(groupLeadDataPlugin)
  bindGroupLeadRoutesToHandlers()
}

func bindGroupLeadRoutesToHandlers() {
  api := router.PathPrefix("/api/group/lead").Subrouter()

  api.HandleFunc("/tasks", groupLeadRequestHandler.GetTasksByGroupId).Methods(http.MethodGet)
  api.HandleFunc("/task", groupLeadRequestHandler.AssignTaskToWorker).Methods(http.MethodPost)
}

func (handler GroupLeadRequestHandler) AssignTaskToWorker(w http.ResponseWriter, r *http.Request) {}

func (handler GroupLeadRequestHandler) GetTasksByGroupId(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var groupTasksReq models.WorkGroupTasksRequest
  decodeRequestBody(r, &groupTasksReq)

  tasks, err := handler.groupLead.GetTasksByGroupId(groupTasksReq.GroupId)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, tasks)
}
