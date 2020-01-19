package api

import (
  "api/helpers"
  "api/middleware"
  "interfaces"
  "models"
  "net/http"
  . "users/group_worker"
)

var groupWorkerRequestHandler GroupWorkerRequestHandler

type GroupWorkerRequestHandler struct {
  groupWorker GroupWorker
  checker helpers.InputChecker
}

func InitGroupWorkerRequestHandler(groupWorkerDataPlugin interfaces.GroupWorkerData) {
  groupWorkerRequestHandler = GroupWorkerRequestHandler{
    groupWorker: NewGroupWorker(groupWorkerDataPlugin),
    checker: helpers.NewInputChecker(),
  }
  bindGroupWorkerRoutesToHandlers()
}

func bindGroupWorkerRoutesToHandlers() {
  api := router.PathPrefix("/group/worker").Subrouter()
  api.Use(middleware.RequiredGroupWorkerRole)

  api.HandleFunc("/tasks", groupWorkerRequestHandler.GetTasksByWorkerId).Methods(http.MethodGet)
  api.HandleFunc("/task/comment", groupWorkerRequestHandler.CommentTask).Methods(http.MethodPatch)
  api.HandleFunc("/task/complete", groupWorkerRequestHandler.CompleteTask).Methods(http.MethodPatch)
}

func (handler GroupWorkerRequestHandler) GetTasksByWorkerId(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var groupWorkerTasksReq models.GroupWorkerTasksRequest
  decodeRequestBody(r, &groupWorkerTasksReq)

  if !handler.checker.IsSafeUint64(groupWorkerTasksReq.WorkerId) {
    panic(getIncorrectUserIdError(groupWorkerTasksReq.WorkerId))
  }

  tasks, err := handler.groupWorker.GetTasksByUserId(groupWorkerTasksReq.WorkerId)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, tasks)
}

func (handler GroupWorkerRequestHandler) CommentTask(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var commentTaskReq models.CommentTaskRequest
  decodeRequestBody(r, &commentTaskReq)

  if !handler.checker.IsSafeUint64(commentTaskReq.TaskId) {
    panic(getIncorrectTaskIdError(commentTaskReq.TaskId))
  } else if !handler.checker.IsLongTextCorrect(commentTaskReq.Comment) {
    panic(getIncorrectTaskCommentError(commentTaskReq.Comment))
  }

  err := handler.groupWorker.AddCommentToTask(commentTaskReq.TaskId, commentTaskReq.Comment)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, nil)
}

func (handler GroupWorkerRequestHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {
  defer sendErrorIfPanicked(w)

  var completeTaskReq models.CompleteTaskRequest
  decodeRequestBody(r, &completeTaskReq)

  if !handler.checker.IsSafeUint64(completeTaskReq.TaskId) {
    panic(getIncorrectTaskIdError(completeTaskReq.TaskId))
  }

  err := handler.groupWorker.MarkTaskAsCompleted(completeTaskReq.TaskId)
  if err != nil {
    panic(err)
  }

  encodeAndSendResponse(w, nil)
}
