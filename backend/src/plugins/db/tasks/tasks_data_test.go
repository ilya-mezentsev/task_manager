package tasks

import (
  "database/sql"
  "fmt"
  mock "mock/plugins"
  "models"
  "os"
  "plugins/db"
  "testing"
  . "utils"
)

var (
  dbFile string
  database *sql.DB
  tasksData DataPlugin
)

func init() {
  dbFile = os.Getenv("TEST_DB_FILE")
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

  tasksData = NewDataPlugin(database)
  db.ExecQuery(database, mock.TurnOnForeignKeys)
  db.CreateGroups(database)
}

func dropTasksTable() {
  db.ExecQuery(database, mock.DropTasksTable)
}

func initTasksTable() {
  dropTasksTable()
  db.ExecQuery(database, mock.CreateTasksTable)
  for _, q := range mock.TestingTasksQueries {
    db.ExecQuery(database, q)
  }
}

func getTaskById(taskId uint) models.Task {
  var task models.Task

  taskRow := database.QueryRow("SELECT * FROM tasks WHERE id = ?", taskId)
  err := taskRow.Scan(
    &task.ID, &task.Title, &task.Description, &task.GroupId,
    &task.UserId, &task.IsComplete, &task.Comment,
  )
  if err != nil {
    fmt.Println("error while getting task by id:", err)
    os.Exit(1)
  }

  return task
}

func TestTasksDataPlugin_GetAllTasksSuccess(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  tasks, err := tasksData.GetAllTasks()

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mock.TasksListEqual(tasks, mock.TestingTasks), func() {
    t.Log(GetExpectationString(mock.TestingTasks, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_GetAllTasksErrorTableNotExists(t *testing.T) {
  dropTasksTable()

  tasks, err := tasksData.GetAllTasks()
  Assert(err != nil, func() {
    t.Log("should be err")
    t.Fail()
  })
  Assert(tasks == nil, func() {
    t.Log(GetExpectationString(nil, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_GetTasksByGroupIdSuccess(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  tasks, err := tasksData.GetTasksByGroupId(1)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mock.TasksListEqual(tasks, mock.TestingTasksByGroupId), func() {
    t.Log(GetExpectationString(mock.TestingTasksByGroupId, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_GetTasksByGroupIdThatNotExists(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  tasks, err := tasksData.GetTasksByGroupId(4)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(tasks == nil, func() {
    t.Log(GetExpectationString(nil, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_GetTasksByGroupIdErrorTableNotExists(t *testing.T) {
  dropTasksTable()

  tasks, err := tasksData.GetTasksByGroupId(1)
  Assert(err != nil, func() {
    t.Log("should be err")
    t.Fail()
  })
  Assert(tasks == nil, func() {
    t.Log(GetExpectationString(nil, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_GetTasksByUserIdSuccess(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  tasks, err := tasksData.GetTasksByUserId(1)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(mock.TasksListEqual(tasks, mock.TestingTasksByUserId), func() {
    t.Log(GetExpectationString(mock.TestingTasksByUserId, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_GetTasksByUserIdThatNotExists(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  tasks, err := tasksData.GetTasksByUserId(4)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(tasks == nil, func() {
    t.Log(GetExpectationString(nil, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_GetTasksByUserIdErrorTableNotExists(t *testing.T) {
  dropTasksTable()

  tasks, err := tasksData.GetTasksByUserId(1)
  Assert(err != nil, func() {
    t.Log("should be err")
    t.Fail()
  })
  Assert(tasks == nil, func() {
    t.Log(GetExpectationString(nil, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_CreateTasksSuccessForOne(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.CreateTasks([]models.Task{mock.TestingTask})
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })

  tasks, _ := tasksData.GetAllTasks()
  expectedTasks := append(mock.TestingTasks, mock.TestingTask)
  Assert(mock.TasksListEqual(expectedTasks, tasks), func() {
    t.Log(GetExpectationString(expectedTasks, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_CreateTasksSuccessForList(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.CreateTasks(mock.TestingTasksAdditional)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })

  tasks, _ := tasksData.GetAllTasks()
  expectedTasks := append(mock.TestingTasks, mock.TestingTasksAdditional...)
  Assert(mock.TasksListEqual(expectedTasks, tasks), func() {
    t.Log(GetExpectationString(expectedTasks, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_CreateTasksErrorTableNotExists(t *testing.T) {
  dropTasksTable()

  err := tasksData.CreateTasks(mock.TestingTasksAdditional)
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestTasksDataPlugin_CreateTasksErrorGroupIdNotExists(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.CreateTasks([]models.Task{mock.TaskWithNotExistsGroupId})
  AssertErrorsEqual(err, db.WorkGroupNotExists, func() {
    t.Log(GetExpectationString(db.WorkGroupNotExists, err))
    t.Fail()
  })

  tasks, _ := tasksData.GetAllTasks()
  Assert(mock.TasksListEqual(mock.TestingTasks, tasks), func() {
    t.Log(GetExpectationString(mock.TestingTasks, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_MarkTaskAsCompleteSuccess(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.MarkTaskAsComplete(1)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })

  task := getTaskById(1)
  Assert(task.IsComplete, func() {
    t.Log("task should be completed")
    t.Fail()
  })
}

func TestTasksDataPlugin_MarkTaskAsCompleteErrorTableNotExists(t *testing.T) {
  dropTasksTable()

  err := tasksData.MarkTaskAsComplete(1)
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestTasksDataPlugin_MarkTaskAsCompleteErrorTaskIdNotExists(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.MarkTaskAsComplete(11)
  AssertErrorsEqual(err, db.TaskIdNotExists, func() {
    t.Log(GetExpectationString(db.TaskIdNotExists, err))
    t.Fail()
  })
}

func TestTasksDataPlugin_CommentTaskSuccess(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.CommentTask(1, mock.TestingComment)
  Assert(err == nil, func() {
    t.Log("should not ber error:", err)
    t.Fail()
  })

  task := getTaskById(1)
  Assert(task.Comment == mock.TestingComment, func() {
    t.Log(GetExpectationString(mock.TestingComment, task.Comment))
    t.Fail()
  })
}

func TestTasksDataPlugin_CommentTaskErrorTableNotExists(t *testing.T) {
  dropTasksTable()

  err := tasksData.CommentTask(1, "")
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestTasksDataPlugin_CommentTaskTaskIdNotExists(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.CommentTask(11, "")
  AssertErrorsEqual(err, db.TaskIdNotExists, func() {
    t.Log(GetExpectationString(db.TaskIdNotExists, err))
    t.Fail()
  })
}

func TestTasksDataPlugin_AssignTaskToWorkerSuccess(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.AssignTaskToWorker(1, 1)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  task := getTaskById(1)
  Assert(task.UserId == 1, func() {
    t.Log(GetExpectationString(1, task.ID))
    t.Fail()
  })
}

func TestTasksDataPlugin_AssignTaskToWorkerErrorTableNotExists(t *testing.T) {
  dropTasksTable()

  err := tasksData.AssignTaskToWorker(1, 1)
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestTasksDataPlugin_AssignTaskToWorkerErrorTaskIdNotExists(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.AssignTaskToWorker(11, 1)
  AssertErrorsEqual(err, db.TaskIdNotExists, func() {
    t.Log(GetExpectationString(db.TaskIdNotExists, err))
    t.Fail()
  })
}

func TestTasksDataPlugin_DeleteTaskSuccess(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.DeleteTask(1)
  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  tasks, _ := tasksData.GetAllTasks()
  expectedTasks := mock.TestingTasks[1:]
  Assert(mock.TasksListEqual(tasks, expectedTasks), func() {
    t.Log(GetExpectationString(expectedTasks, tasks))
    t.Fail()
  })
}

func TestTasksDataPlugin_DeleteTaskErrorTableNotExists(t *testing.T) {
  dropTasksTable()

  err := tasksData.DeleteTask(1)
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestTasksDataPlugin_DeleteTaskErrorIdNotExists(t *testing.T) {
  initTasksTable()
  defer dropTasksTable()

  err := tasksData.DeleteTask(11)
  AssertErrorsEqual(err, db.TaskIdNotExists, func() {
    t.Log(GetExpectationString(db.TaskIdNotExists, err))
    t.Fail()
  })
}
