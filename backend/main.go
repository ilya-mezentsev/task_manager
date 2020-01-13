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

var staticFilesDirPath string

func init() {
  staticFilesDirPath = os.Getenv("STATIC_DIR")
  if staticFilesDirPath == "" {
    log.Println("STATIC_DIR env var is not set")
    os.Exit(1)
  }

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
  r := api.GetRouter()
  r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFilesDirPath)))
  api.InitAdminRequestHandler(dbProxy)
  api.InitGroupLeadRequestHandler(dbProxy)
  api.InitGroupWorkerRequestHandler(dbProxy)
}

func main() {
  srv := &http.Server{
    Handler: api.GetRouter(),
    Addr: "127.0.0.1:8181",
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }

  log.Fatal(srv.ListenAndServe())
}
