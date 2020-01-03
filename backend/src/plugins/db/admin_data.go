package db

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "models"
)

const (
  AllUsersQuery = "SELECT * FROM users"
)

type AdminData struct {
  database *sql.DB
}

func NewAdminDataProvider(db *sql.DB) AdminData {
  return AdminData{database: db}
}

func (a AdminData) GetAllUsers() ([]models.User, error) {
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
