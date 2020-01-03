package models

type User struct {
  ID uint `db:"id"`
  Name string `db:"name"`
  GroupId uint `db:"group_id"`
  Password string `db:"password"`
  IsGroupLead bool `db:"is_group_lead"`
}
