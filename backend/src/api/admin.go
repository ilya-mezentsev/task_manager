package api

import "net/http"

func BindAdminRoutesToHandlers() {
  api := router.PathPrefix("/api/admin").Subrouter()

  api.HandleFunc("groups", nil).Methods(http.MethodGet)
  api.HandleFunc("group", nil).Methods(http.MethodPost)
  api.HandleFunc("group", nil).Methods(http.MethodDelete)

  api.HandleFunc("users", nil).Methods(http.MethodGet)
  api.HandleFunc("user", nil).Methods(http.MethodPost)
  api.HandleFunc("user", nil).Methods(http.MethodDelete)

  api.HandleFunc("tasks", nil).Methods(http.MethodGet)
  api.HandleFunc("tasks", nil).Methods(http.MethodPost)
  api.HandleFunc("task", nil).Methods(http.MethodDelete)
}
