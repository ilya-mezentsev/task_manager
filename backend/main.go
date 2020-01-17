package main

import (
  "api"
  "database/sql"
  "log"
  "net/http"
  "os"
  "plugins"
  "time"
)

func init() {
  dbFile := os.Getenv("DB_FILE")
  if dbFile == "" {
    log.Println("DB_FILE env var is not set")
    os.Exit(1)
  }

  database, err := sql.Open("sqlite3", dbFile)
  if err != nil {
    log.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  dbProxy := plugins.NewDBProxy(database)
  dbProxy.InitDBStructure(database)
  api.InitLoginRequestHandler(dbProxy)
  api.InitAdminRequestHandler(dbProxy)
  api.InitGroupLeadRequestHandler(dbProxy)
  api.InitGroupWorkerRequestHandler(dbProxy)
}

func main() {
  srv := &http.Server{
    Handler: api.GetRouter(),
    Addr: "0.0.0.0:8080",
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }

  log.Fatal(srv.ListenAndServe())
}
