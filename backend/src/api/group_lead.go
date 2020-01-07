package api

import "net/http"

func BindGroupLeadRoutesToHandlers() {
  api := router.PathPrefix("/api/group/lead").Subrouter()

  api.HandleFunc("task", nil).Methods(http.MethodPost)
}
