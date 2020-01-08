package models

type Group struct {
  ID uint `db:"id" json:"id"`
  Name string `db:"name" json:"name"`
}
