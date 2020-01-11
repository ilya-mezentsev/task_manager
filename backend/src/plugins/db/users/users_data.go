package users

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "models"
  "plugins/db"
)

const (
  AllUsersQuery = "SELECT * FROM users"
  GetUsersByGroupId = "SELECT * FROM users WHERE group_id = ?"
  GetUserById = "SELECT * FROM users WHERE id = ?"
  CreateUserQuery = "INSERT INTO users VALUES(NULL, ?, ?, ?, ?)"
  DeleteUserQuery = "DELETE FROM users WHERE id = ?"
  UserNameAlreadyExistsMessage = "UNIQUE constraint failed: users.name"
  // users table has only work_id foreign key
  WGIdNotExists = "FOREIGN KEY constraint failed"
)

type DataPlugin struct {
  database *sql.DB
}

func NewDataPlugin(db *sql.DB) DataPlugin {
  return DataPlugin{database: db}
}

func (u DataPlugin) GetAllUsers() ([]models.User, error) {
  return u.getUsersSequence(AllUsersQuery)
}

func (u DataPlugin) GetUsersByGroupId(groupId uint) ([]models.User, error) {
  return u.getUsersSequence(GetUsersByGroupId, groupId)
}

func (u DataPlugin) getUsersSequence(query string, args ...interface{}) ([]models.User, error) {
  usersRows, err := u.database.Query(query, args...)
  if err != nil {
    return nil, err
  }
  var users []models.User

  for usersRows.Next() {
    var user models.User
    err = usersRows.Scan(&user.ID, &user.Name, &user.GroupId, &user.Password, &user.IsGroupLead)
    if err != nil {
      return nil, err
    }

    users = append(users, user)
  }

  return users, nil
}

func (u DataPlugin) GetUser(userId uint) (models.User, error) {
  var (
    user models.User
    emptyUser models.User
  )

  userRow := u.database.QueryRow(GetUserById, userId)
  switch err := userRow.Scan(&user.ID, &user.Name, &user.GroupId, &user.Password, &user.IsGroupLead); err {
  case nil: // no errors, it's ok
    return user, nil
  case sql.ErrNoRows: // we need to wrap this case here
    return emptyUser, db.WorkerIdNotExists
  default:
    return emptyUser, err
  }
}

func (u DataPlugin) CreateUser(user models.User) (uint, error) {
  statement, err := u.database.Prepare(CreateUserQuery)
  if err != nil {
    return 0, err
  }

  result, err := statement.Exec(u.getCreatingFieldsSequence(user)...)
  if err != nil {
    switch err.Error() {
    case UserNameAlreadyExistsMessage:
      return 0, db.UserNameAlreadyExists
    case WGIdNotExists:
      return 0, db.WorkGroupNotExists
    default:
      return 0, err
    }
  }

  // ignore error here coz SQLite driver does not return it
  lastInsertedId, _ := result.LastInsertId()

  return uint(lastInsertedId), nil
}

func (u DataPlugin) getCreatingFieldsSequence(user models.User) []interface{} {
  return []interface{}{
    user.Name,
    user.GroupId,
    user.Password,
    user.IsGroupLead,
  }
}

func (u DataPlugin) DeleteUser(userId uint) error {
  statement, err := u.database.Prepare(DeleteUserQuery)
  if err != nil {
    return err
  }

  res, err := statement.Exec(userId)
  if err != nil {
    return err
  }

  if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
    return db.UserIdNotExists
  }

  return nil
}
