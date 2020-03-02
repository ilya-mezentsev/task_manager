package api

import (
  "api/helpers"
  "api/middleware"
  "interfaces"
  "models"
  "net/http"
  . "users/group_lead"
)

var groupLeadRequestHandler GroupLeadRequestHandler

type GroupLeadRequestHandler struct {
  groupLead GroupLead
  checker helpers.InputChecker
}

func InitGroupLeadRequestHandler(groupLeadDataPlugin interfaces.GroupLeadData) {
  groupLeadRequestHandler = GroupLeadRequestHandler{
    groupLead: NewGroupLead(groupLeadDataPlugin),
    checker: helpers.NewInputChecker(),
  }
  bindGroupLeadRoutesToHandlers()
}

func bindGroupLeadRoutesToHandlers() {
  api := router.PathPrefix("/group/lead").Subrouter()
  api.Use(middleware.RequiredGroupLeadRole)

  api.HandleFunc("/users", groupLeadRequestHandler.GetUsersByGroupId).Methods(http.MethodPost)
  api.HandleFunc("/tasks", groupLeadRequestHandler.GetTasksByGroupId).Methods(http.MethodPost)
  api.HandleFunc("/task", groupLeadRequestHandler.AssignTaskToWorker).Methods(http.MethodPost)
}

func (handler GroupLeadRequestHandler) GetUsersByGroupId(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var workGroupUsersReq models.WorkGroupUsersRequest
  decodeRequestBody(r, &workGroupUsersReq)

  if !handler.checker.IsSafeUint64(workGroupUsersReq.GroupId) {
    panic(getIncorrectGroupIdError(workGroupUsersReq.GroupId))
  }

  users, err := handler.groupLead.GetUsersByGroupId(workGroupUsersReq.GroupId)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, users)
}

func (handler GroupLeadRequestHandler) AssignTaskToWorker(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var assignTaskReq models.AssignTaskToGroupWorkerRequest
  decodeRequestBody(r, &assignTaskReq)

  if !handler.checker.IsSafeUint64(assignTaskReq.WorkerId) {
    panic(getIncorrectUserIdError(assignTaskReq.WorkerId))
  } else if !handler.checker.IsSafeUint64(assignTaskReq.Task.ID) {
    // we do not need to check another fields coz they are not used
    panic(getIncorrectTaskIdError(assignTaskReq.Task.ID))
  }

  err := handler.groupLead.AssignTaskToWorker(assignTaskReq.WorkerId, assignTaskReq.Task)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, nil)
}

func (handler GroupLeadRequestHandler) GetTasksByGroupId(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var groupTasksReq models.WorkGroupTasksRequest
  decodeRequestBody(r, &groupTasksReq)

  if !handler.checker.IsSafeUint64(groupTasksReq.GroupId) {
    panic(getIncorrectGroupIdError(groupTasksReq.GroupId))
  }

  tasks, err := handler.groupLead.GetTasksByGroupId(groupTasksReq.GroupId)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, tasks)
}
