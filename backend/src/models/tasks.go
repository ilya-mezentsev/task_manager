package models

type Task struct {
  ID uint `db:"id" json:"id"`
  Title string `db:"title" json:"title"`
  Description string `db:"description" json:"description"`
  GroupId uint `db:"group_id" json:"group_id"`
  UserId uint `db:"user_id" json:"user_id"`
  IsComplete bool `db:"is_complete" json:"is_complete"`
  Comment string `db:"comment" json:"comment"`
}
