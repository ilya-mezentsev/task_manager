package api

import (
  "interfaces"
  "net/http"
  . "users/group_lead"
)

var groupLeadRequestHandler GroupLeadRequestHandler

type GroupLeadRequestHandler struct {
  groupLead GroupLead
}

func InitGroupLeadRequestHandler(groupLeadDataPlugin interfaces.GroupLeadData) {
  groupLeadRequestHandler.groupLead = NewGroupLead(groupLeadDataPlugin)
}

func BindGroupLeadRoutesToHandlers() {
  api := router.PathPrefix("/api/group/lead").Subrouter()

  api.HandleFunc("task", groupLeadRequestHandler.AssignTaskToWorker).Methods(http.MethodPost)
}

func (handler GroupLeadRequestHandler) AssignTaskToWorker(w http.ResponseWriter, r *http.Request) {}
