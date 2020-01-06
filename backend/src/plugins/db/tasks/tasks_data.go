package tasks

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "models"
  "processing"
)

const (
  AllTasksQuery = "SELECT * FROM tasks"
  GetTaskById = "SELECT * FROM tasks WHERE id = ?"
  AddTaskQuery = "INSERT INTO tasks VALUES(NULL, ?, ?, ?, ?, ?, ?)"
  // tasks table has only work_id foreign key
  WGIdNotExistsMessage = "FOREIGN KEY constraint failed"

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

func (t TasksDataPlugin) GetTaskById(taskId uint) (models.Task, error) {
  var (
    task models.Task
    emptyTask models.Task
  )

  taskRow := t.database.QueryRow(GetTaskById, taskId)
  err := taskRow.Scan(
    &task.ID, &task.Title, &task.Description, &task.GroupId,
    &task.UserId, &task.IsComplete, &task.Comment,
  )
  switch err {
  case nil: // no errors, it's ok
    return task, nil
  case sql.ErrNoRows:
    return emptyTask, processing.TaskIdNotExists
  default:
    return emptyTask, err
  }
}

func (t TasksDataPlugin) CreateTasks(tasks []models.Task) error {
  tx, err := t.database.Begin()
  if err != nil {
    return err
  }

  for _, task := range tasks {
    _, err := tx.Exec(AddTaskQuery, t.getCreatingFieldsSequence(task)...)
    if err != nil {
      err = t.parseDBError(err)
      t.assignRollbackErrorIfExists(&err, tx.Rollback())

      return err
    }
  }

  return tx.Commit()
}

func (t TasksDataPlugin) getCreatingFieldsSequence(task models.Task) []interface{} {
  return []interface{}{
    task.Title,
    task.Description,
    task.GroupId,
    task.UserId,
    task.IsComplete,
    task.Comment,
  }
}

func (t TasksDataPlugin) parseDBError(err error) error {
  switch err.Error() {
  case WGIdNotExistsMessage:
    return processing.WorkGroupNotExists
  default:
    return err
  }
}

func (t TasksDataPlugin) assignRollbackErrorIfExists(err *error, rollbackError error) {
  if rollbackError != nil {
    *err = rollbackError
  }
}

func (t TasksDataPlugin) MarkTaskAsComplete(taskId uint) error {
  return nil
}

func (t TasksDataPlugin) CommentTask(taskId uint, comment string) error {
  return nil
}
