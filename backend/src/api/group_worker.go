package api

import "net/http"

func BindGroupWorkerRoutesToHandlers() {
  api := router.PathPrefix("/api/group/worker").Subrouter()

  api.HandleFunc("tasks", nil).Methods(http.MethodGet)
  api.HandleFunc("task/comment", nil).Methods(http.MethodPatch)
  api.HandleFunc("task/complete", nil).Methods(http.MethodPatch)
}
