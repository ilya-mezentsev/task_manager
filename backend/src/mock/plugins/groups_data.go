package mock

const (
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
)
