package plugins

import (
  "database/sql"
  "log"
  "models"
  "plugins/db/groups"
  "plugins/db/tasks"
  "plugins/db/users"
)

type DBProxy struct {
  groupsData groups.DataPlugin
  tasksData tasks.DataPlugin
  usersData users.DataPlugin
}

func NewDBProxy(db *sql.DB) DBProxy {
  return DBProxy{
    groupsData: groups.NewDataPlugin(db),
    tasksData: tasks.NewDataPlugin(db),
    usersData: users.NewDataPlugin(db),
  }
}

func (proxy DBProxy) AddCommentToTask(taskId uint, comment string) error {
  log.Printf("adding comment '%s', to task with id <%d>\n", comment, taskId)

  err := proxy.tasksData.CommentTask(taskId, comment)
  if err != nil {
    log.Printf("error while commenting task (id <%d>): %v\n", taskId, err)
    return err
  }

  return nil
}

func (proxy DBProxy) MarkTaskAsCompleted(taskId uint) error {
  log.Printf("mark task (id <%d>) as complete\n", taskId)

  err := proxy.tasksData.MarkTaskAsComplete(taskId)
  if err != nil {
    log.Printf("error while completing task (id <%d>): %v\n", taskId, err)
    return err
  }

  return nil
}

func (proxy DBProxy) AssignTaskToWorker(workerId uint, task models.Task) error {
  log.Printf("assigning task (id <%d>) to worker (id <%d>)\n", task.ID, workerId)

  err := proxy.tasksData.AssignTaskToWorker(task.ID, workerId)
  if err != nil {
    log.Printf("error while assigning task (id <%d>) to worker (id <%d>): %v\n", task.ID, workerId, err)
    return err
  }

  return nil
}

func (proxy DBProxy) CreateUser(user models.User) error {
  log.Printf("creating user (%v)\n", user)

  userId, err := proxy.usersData.CreateUser(user)
  if err != nil {
    log.Printf("error while creating user: %v\n", err)
    return err
  }

  log.Printf("created user id: %d\n", userId)
  return nil
}

func (proxy DBProxy) CreateWorkGroup(groupName string) error {
  log.Printf("starting to create work group: %s\n", groupName)

  err := proxy.groupsData.CreateWorkGroup(groupName)
  if err != nil {
    log.Printf("error while creating work group: %s\n", err)
    return err
  }

  return nil
}

func (proxy DBProxy) AssignTasksToGroup(groupId uint, tasks []models.Task) error {
  log.Printf("assigning tasks to work group (id <%d>)\n", groupId)
  for i, task := range tasks {
    log.Printf("\t%d: %v\n", i, task)
  }

  err := proxy.tasksData.CreateTasks(tasks)
  if err != nil {
    log.Printf("error while assigning tasks to work group (id <%d>): %v\n", groupId, err)
    return err
  }

  return nil
}
