package api

import (
  "bytes"
  "database/sql"
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "log"
  "mock"
  mock2 "mock/plugins"
  "net/http"
  "net/http/httptest"
  "os"
  "plugins"
  "plugins/db"
  "testing"
  . "utils"
)

var database *sql.DB

func init() {
  dbFile := os.Getenv("TEST_DB_FILE")
  if dbFile == "" {
    fmt.Println("TEST_DB_FILE env var is not set")
    os.Exit(1)
  }

  var err error
  database, err = sql.Open("sqlite3", dbFile)
  if err != nil {
    fmt.Println("An error while opening db file:", err)
    os.Exit(1)
  }

  InitAdminRequestHandler(plugins.NewDBProxy(database))
  initTestTables()
}

func dropTestTables() {
  for _, q := range mock.DropTestTablesQueries {
    db.ExecQuery(database, q)
  }
}

func initTestTables() {
  dropTestTables()
  for _, q := range mock.CreateTestTablesQueries {
    db.ExecQuery(database, q)
  }
  for _, q := range mock.AddDataToTestTablesQueries {
    db.ExecQuery(database, q)
  }
}

func makeRequest(t *testing.T, method, endpoint, data string) io.ReadCloser {
  srv := httptest.NewServer(router)
  defer srv.Close()

  client := &http.Client{}
  req, err := http.NewRequest(
    method,
    fmt.Sprintf("%s/api/%s", srv.URL, endpoint),
    bytes.NewBuffer([]byte(data)),
  )
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })

  req.Header.Set("Content-Type", "application/json; charset=utf-8")
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  return resp.Body
}

func TestMain(m *testing.M) {
  log.SetOutput(ioutil.Discard)
  os.Exit(m.Run())
}

func TestGetAllGroupsSuccess(t *testing.T) {
  var response mock.AllGroupsResponse
  responseBody := makeRequest(t, http.MethodGet, "admin/groups", "")
  err := json.NewDecoder(responseBody).Decode(&response)

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(response.Status == "ok", func() {
    t.Log(GetExpectationString("ok", response.Status))
    t.Fail()
  })
  Assert(mock2.GroupListEqual(response.Data, mock2.TestingGroups), func() {
    t.Log(GetExpectationString(mock2.TestingGroups, response.Data))
    t.Fail()
  })
}
