package mock

import "models"

const (
  ClearUsersTable = "delete from users;"
  DropUsersTable = "drop table if exists users;"
  CreateUsersTable = `
  create table if not exists users(
    id integer not null primary key autoincrement,
    name text not null unique,
    group_id integer not null,
    password text not null,
    is_group_lead integer default 0
  )`
)

var (
  TestingUsersQueries = []string{
    "insert into users values(1, 'name1', 1, 'some_pass', 0);",
    "insert into users values(2, 'name2', 2, 'some_pass', 0);",
    "insert into users values(3, 'name3', 1, 'some_pass', 1);",
  }
  EmptyUser models.User
  TestingUser = models.User{
    Name: "name4",
    GroupId: 3,
    Password: "some_pass",
    IsGroupLead: false,
  }
  TestingUserWithExistsName = models.User{
    Name: "name1",
  }
  TestingUsers = []models.User{
    {
      ID: 1,
      Name: "name1",
      GroupId: 1,
      Password: "some_pass",
      IsGroupLead: false,
    },
    {
      ID: 2,
      Name: "name2",
      GroupId: 2,
      Password: "some_pass",
      IsGroupLead: false,
    },
    {
      ID: 3,
      Name: "name3",
      GroupId: 1,
      Password: "some_pass",
      IsGroupLead: true,
    },
  }
)

func UserListEqual(l1 , l2 []models.User) bool {
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
