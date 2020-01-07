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
  AssignTaskToWorker = "UPDATE tasks SET user_id = ? WHERE id = ?"
  MarkTaskAsCompleteQuery = "UPDATE tasks SET is_complete = 1 WHERE id = ?"
  CommentTaskQuery = "UPDATE tasks SET comment = ? WHERE id = ?"
  // tasks table has only work_id foreign key
  WGIdNotExistsMessage = "FOREIGN KEY constraint failed"
)

type DataPlugin struct {
  database *sql.DB
}

func NewDataPlugin(driver *sql.DB) DataPlugin {
  return DataPlugin{database: driver}
}

func (t DataPlugin) GetAllTasks() ([]models.Task, error) {
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

func (t DataPlugin) CreateTasks(tasks []models.Task) error {
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

func (t DataPlugin) getCreatingFieldsSequence(task models.Task) []interface{} {
  return []interface{}{
    task.Title,
    task.Description,
    task.GroupId,
    task.UserId,
    task.IsComplete,
    task.Comment,
  }
}

func (t DataPlugin) parseDBError(err error) error {
  switch err.Error() {
  case WGIdNotExistsMessage:
    return processing.WorkGroupNotExists
  default:
    return err
  }
}

func (t DataPlugin) assignRollbackErrorIfExists(err *error, rollbackError error) {
  if rollbackError != nil {
    *err = rollbackError
  }
}

func (t DataPlugin) AssignTaskToWorker(taskId, workerId uint) error {
  return t.execUpdating(AssignTaskToWorker, workerId, taskId)
}

func (t DataPlugin) MarkTaskAsComplete(taskId uint) error {
  return t.execUpdating(MarkTaskAsCompleteQuery, taskId)
}

func (t DataPlugin) CommentTask(taskId uint, comment string) error {
  return t.execUpdating(CommentTaskQuery, comment, taskId)
}

func (t DataPlugin) execUpdating(query string, args ...interface{}) error {
  statement, err := t.database.Prepare(query)
  if err != nil {
    return err
  }

  res, err := statement.Exec(args...)
  if err != nil {
    return err
  }

  // ignore error here coz sqlite does not return it
  if affectedRows, _ := res.RowsAffected(); affectedRows == 0 {
    return processing.TaskIdNotExists
  }
  return nil
}
