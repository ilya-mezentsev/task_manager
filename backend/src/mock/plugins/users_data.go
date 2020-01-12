package mock

import "models"

const (
  TurnOnForeignKeys = "PRAGMA foreign_keys = ON;"
  DropUsersTable = "drop table if exists users;"
  CreateUsersTable = `
  create table if not exists users(
    id integer not null primary key autoincrement,
    name text not null unique,
    group_id integer not null,
    password text not null,
    is_group_lead integer default 0,
    foreign key(group_id) references groups(id) on delete cascade
  )`
)

var (
  TestingUsersQueries = []string{
    "insert into users values(1, 'name1', 1, 'e261f61e493f309cca25ee89e98518c0', 0);",
    "insert into users values(2, 'name2', 2, 'e261f61e493f309cca25ee89e98518c0', 0);",
    "insert into users values(3, 'name3', 1, 'e261f61e493f309cca25ee89e98518c0', 1);",
  }
  TestingUsers = []models.User{
    {
      ID: 1,
      Name: "name1",
      GroupId: 1,
      Password: "e261f61e493f309cca25ee89e98518c0",
      IsGroupLead: false,
    },
    {
      ID: 2,
      Name: "name2",
      GroupId: 2,
      Password: "e261f61e493f309cca25ee89e98518c0",
      IsGroupLead: false,
    },
    {
      ID: 3,
      Name: "name3",
      GroupId: 1,
      Password: "e261f61e493f309cca25ee89e98518c0",
      IsGroupLead: true,
    },
  }
  TestingUsersByGroupId = []models.User{
    {
      ID: 2,
      Name: "name2",
      GroupId: 2,
      Password: "e261f61e493f309cca25ee89e98518c0",
      IsGroupLead: false,
    },
  }
  EmptyUser models.User
  TestingUser = models.User{
    Name: "name4",
    GroupId: 3,
    Password: "e261f61e493f309cca25ee89e98518c0",
    IsGroupLead: false,
  }
  TestingUserWithExistsName = models.User{
    Name: "name1",
  }
  TestingUserWithNotExistsGroupId = models.User{
    GroupId: 11,
  }
  TestingCredentials = [2]string{"name1", "e261f61e493f309cca25ee89e98518c0"}
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
