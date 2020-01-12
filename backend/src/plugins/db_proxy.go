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

func (proxy DBProxy) GetAllTasks() ([]models.Task, error) {
  log.Println("requesting all tasks")

  allTasks, err := proxy.tasksData.GetAllTasks()
  if err != nil {
    log.Println("error while requesting all tasks:", err)
    return nil, err
  }

  return allTasks, nil
}

func (proxy DBProxy) GetTasksByGroupId(groupId uint) ([]models.Task, error) {
  log.Printf("requesting tasks by group id: <%d>\n", groupId)

  tasksByGroupId, err := proxy.tasksData.GetTasksByGroupId(groupId)
  if err != nil {
    log.Printf("error while requesting tasks by group id (id <%d>): %v\n", groupId, err)
    return nil, err
  }

  return tasksByGroupId, nil
}

func (proxy DBProxy) GetTasksByUserId(userId uint) ([]models.Task, error) {
  log.Printf("requesting tasks by user id: <%d>\n", userId)

  tasksByUserId, err := proxy.tasksData.GetTasksByUserId(userId)
  if err != nil {
    log.Printf("error while requesting tasks by user id (id <%d>): %v\n", userId, err)
    return nil, err
  }

  return tasksByUserId, nil
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

func (proxy DBProxy) DeleteTask(taskId uint) error {
  log.Printf("deleting task (id <%d>)\n", taskId)

  err := proxy.tasksData.DeleteTask(taskId)
  if err != nil {
    log.Printf("error while deleting task (id <%d>): %v\n", taskId, err)
    return err
  }

  return nil
}

func (proxy DBProxy) GetAllUsers() ([]models.User, error) {
  log.Println("requesting all users")

  allUsers, err := proxy.usersData.GetAllUsers()
  if err != nil {
    log.Println("error while requesting all users:", err)
    return nil, err
  }

  return allUsers, nil
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

func (proxy DBProxy) DeleteUser(userId uint) error {
  log.Printf("deleting user (id <%d>)\n", userId)

  err := proxy.usersData.DeleteUser(userId)
  if err != nil {
    log.Printf("error while deleing user (id <%d>): %v\n", userId, err)
    return err
  }

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

func (proxy DBProxy) GetAllGroups() ([]models.Group, error) {
  log.Println("requesting all groups")

  allGroups, err := proxy.groupsData.GetAllGroups()
  if err != nil {
    log.Println("error while requesting all groups:", err)
    return nil, err
  }

  return allGroups, nil
}

func (proxy DBProxy) DeleteWorkGroup(groupId uint) error {
  log.Printf("deleting work group (id <%d>)\n", groupId)

  err := proxy.groupsData.DeleteWorkGroup(groupId)
  if err != nil {
    log.Printf("error while deleting work group (id <%d>)\n", groupId)
    return err
  }

  return nil
}
