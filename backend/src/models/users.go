package models

type (
  User struct {
    ID uint `db:"id" json:"id"`
    Name string `db:"name" json:"name"`
    GroupId uint `db:"group_id" json:"group_id"`
    Password string `db:"password" json:"password"`
    IsGroupLead bool `db:"is_group_lead" json:"is_group_lead"`
  }

  UserSession struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Role string `json:"role"`
  }
)
