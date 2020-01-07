package mock

import "models"

const (
  ExistsGroupName = "group1"
  DropGroupsTable = "drop table if exists groups;"
  CreateGroupsTable = `
  create table if not exists groups(
    id integer not null primary key autoincrement,
    name text not null unique
  )`
)

var (
  TestingGroupsQueries = []string{
    "insert into groups values(1, 'group1')",
    "insert into groups values(2, 'group2')",
    "insert into groups values(3, 'group3')",
  }
  TestingGroups = []models.Group{
    {ID: 1, Name: "group1"},
    {ID: 2, Name: "group2"},
    {ID: 3, Name: "group3"},
  }
  TestingGroup = models.Group{
    ID: 4,
    Name: "group4",
  }
)

func GroupListEqual(l1, l2 []models.Group) bool {
  if len(l1) != len(l2) {
    return false
  }

  for i, _ := range l1 {
    if l1[i] != l2[i] {
      return false
    }
  }

  return true
}
