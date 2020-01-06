package tasks

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "models"
)

const (
  AllTasksQuery = "SELECT * FROM tasks"
)

type TasksDataPlugin struct {
  database *sql.DB
}

func NewTasksDataPlugin(driver *sql.DB) TasksDataPlugin {
  return TasksDataPlugin{database: driver}
}

func (t TasksDataPlugin) GetAllTasks() ([]models.Task, error) {
  tasksRows, err := t.database.Query(AllTasksQuery)
  if err != nil {
    return nil, err
  }
  var tasks []models.Task

  for tasksRows.Next() {
    var task models.Task
    err = tasksRows.Scan(
      &task.ID, &task.Title, &task.Description, &task.GroupId, &task.UserId, &task.IsComplete, &task.Comment,
    )
    if err != nil {
      return nil, err
    }

    tasks = append(tasks, task)
  }

  return tasks, nil
}
