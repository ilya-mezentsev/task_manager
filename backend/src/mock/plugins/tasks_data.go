package mock

import "models"

const (
  DropTasksTable = "drop table if exists tasks;"
  CreateTasksTable = `
  create table if not exists tasks(
    id integer not null primary key autoincrement,
    title text not null,
    description text default '',
    group_id integer not null,
    user_id integer default 0,
    is_complete integer default 0,
    comment text default '',
    foreign key(group_id) references groups(id)
  )`
)

var (
  TestingTasksQueries = []string{
    "insert into tasks(id, title, group_id) values(1, 'title1', 1)",
    "insert into tasks(id, title, group_id) values(2, 'title2', 2)",
    "insert into tasks(id, title, group_id) values(3, 'title1', 3)",
  }
  TestingTasks = []models.Task{
    {ID: 1, Title: "title1", GroupId: 1},
    {ID: 2, Title: "title2", GroupId: 2},
    {ID: 3, Title: "title1", GroupId: 3},
  }
)

func TasksListEqual(l1, l2 []models.Task) bool {
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
