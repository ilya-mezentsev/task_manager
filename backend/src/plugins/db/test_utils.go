package db

import (
  "database/sql"
  "fmt"
  mock "mock/plugins"
  "os"
)

func ExecQuery(database *sql.DB, q string, args ...interface{}) {
  statement, err := database.Prepare(q)
  if err != nil {
    fmt.Println("An error while preparing db statement:", err)
    os.Exit(1)
  }

  _, err = statement.Exec(args...)
  if err != nil {
    fmt.Println("An error while creating db structure:", err)
    os.Exit(1)
  }
}

func CreateGroups(database *sql.DB) {
  ExecQuery(database, mock.DropGroupsTable)
  ExecQuery(database, mock.CreateGroupsTable)
  for _, q := range mock.TestingGroupsQueries {
    ExecQuery(database, q)
  }
}
