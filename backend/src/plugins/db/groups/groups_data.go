package groups

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "models"
  "processing"
)

const (
  GetAllGroupsQuery = "SELECT * FROM groups"
  CreateGroupQuery = "INSERT INTO groups VALUES(NULL, ?)"
  GroupNameAlreadyExistsMessage = "UNIQUE constraint failed: groups.name"
)

type DataPlugin struct {
  database *sql.DB
}

func NewDataPlugin(provider *sql.DB) DataPlugin {
  return DataPlugin{database: provider}
}

func (g DataPlugin) GetAllGroups() ([]models.Group, error) {
  groupsRows, err := g.database.Query(GetAllGroupsQuery)
  if err != nil {
    return nil, err
  }

  var groups []models.Group
  for groupsRows.Next() {
    var group models.Group
    err = groupsRows.Scan(&group.ID, &group.Name)
    if err != nil {
      return nil, err
    }

    groups = append(groups, group)
  }

  return groups, nil
}

func (g DataPlugin) CreateWorkGroup(groupName string) error {
  statement, err := g.database.Prepare(CreateGroupQuery)
  if err != nil {
    return err
  }

  _, err = statement.Exec(groupName)
  if err != nil {
    switch err.Error() {
    case GroupNameAlreadyExistsMessage:
      return processing.WorkGroupAlreadyExists
    default:
      return err
    }
  }

  return nil
}
