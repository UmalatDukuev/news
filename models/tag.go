package models

type Tag struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
