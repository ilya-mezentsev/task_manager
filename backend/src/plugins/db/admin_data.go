package db

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "models"
)

const (
  AllUsersQuery = "SELECT * FROM users"
)

type UsersDataPlugin struct {
  database *sql.DB
}

func NewUsersDataPlugin(db *sql.DB) UsersDataPlugin {
  return UsersDataPlugin{database: db}
}

func (a UsersDataPlugin) GetAllUsers() ([]models.User, error) {
  usersRows, err := a.database.Query(AllUsersQuery)
  if err != nil {
    return nil, err
  }
  var users []models.User

  for usersRows.Next() {
    var user models.User
    if err = usersRows.Scan(&user.ID, &user.Name, &user.GroupId, &user.Password, &user.IsGroupLead); err != nil {
      return nil, err
    }

    users = append(users, user)
  }

  return users, nil
}
