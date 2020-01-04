package db

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "models"
  "processing"
)

const (
  AllUsersQuery = "SELECT * FROM users"
  GetUserById = "SELECT * FROM users WHERE id = ?"
  CreateUserQuery = "INSERT INTO users VALUES(NULL, ?, ?, ?, ?)"
  DeleteUserQuery = "DELETE FROM users WHERE id = ?"
  UserNameAlreadyExistsMessage = "UNIQUE constraint failed: users.name"
)

type UsersDataPlugin struct {
  database *sql.DB
}

func NewUsersDataPlugin(db *sql.DB) UsersDataPlugin {
  return UsersDataPlugin{database: db}
}

func (u UsersDataPlugin) GetAllUsers() ([]models.User, error) {
  usersRows, err := u.database.Query(AllUsersQuery)
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

func (u UsersDataPlugin) GetUser(userId uint) (models.User, error) {
  var (
    user models.User
    emptyUser models.User
  )

  userRow := u.database.QueryRow(GetUserById, userId)
  switch err := userRow.Scan(&user.ID, &user.Name, &user.GroupId, &user.Password, &user.IsGroupLead); err {
  case nil: // no errors, it's ok
    return user, nil
  case sql.ErrNoRows: // we need to wrap this case here
    return emptyUser, processing.WorkerIdNotExists
  default:
    return emptyUser, err
  }
}

func (u UsersDataPlugin) CreateUser(user models.User) (uint, error) {
  statement, err := u.database.Prepare(CreateUserQuery)
  if err != nil {
    return 0, err
  }

  result, err := statement.Exec(u.getCreatingFieldsSequence(user)...)
  if err != nil {
    switch err.Error() {
    case UserNameAlreadyExistsMessage:
      return 0, processing.UserNameAlreadyExists
    default:
      return 0, err
    }
  }

  // ignore error here coz SQLite driver does not return it
  lastInsertedId, _ := result.LastInsertId()

  return uint(lastInsertedId), nil
}

func (u UsersDataPlugin) getCreatingFieldsSequence(user models.User) []interface{} {
  return []interface{}{
    user.Name,
    user.GroupId,
    user.Password,
    user.IsGroupLead,
  }
}

func (u UsersDataPlugin) DeleteUser(userId uint) error {
  statement, err := u.database.Prepare(DeleteUserQuery)
  if err != nil {
    return err
  }

  _, err = statement.Exec(userId)
  return err
}
