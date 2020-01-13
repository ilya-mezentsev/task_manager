package main

import (
  "api"
  "database/sql"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "os"
  "plugins"
  "regexp"
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
  api.InitAdminRequestHandler(dbProxy)
  api.InitGroupLeadRequestHandler(dbProxy)
  api.InitGroupWorkerRequestHandler(dbProxy)
}

func main() {
  r := api.GetRouter()
  r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFilesDirPath)))
  r.Methods("GET").MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
   match, _ := regexp.MatchString(`/(^api)`, r.URL.Path)

   return match
  }).Handler(http.RedirectHandler("/", http.StatusPermanentRedirect))

  srv := &http.Server{
    Handler: r,
    Addr: "127.0.0.1:8181",
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }

  log.Fatal(srv.ListenAndServe())
}
