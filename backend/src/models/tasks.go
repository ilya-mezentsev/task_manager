package models

type Task struct {
  ID uint `db:"id"`
  Title string `db:"title"`
  Description string `db:"description"`
  GroupId uint `db:"group_id"`
  UserId uint `db:"user_id"`
  IsComplete bool `db:"is_complete"`
  Comment string `db:"comment"`
}
