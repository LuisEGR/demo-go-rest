package model

// Usuario model
type Usuario struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"nombre"`
	Time string `db:"time" json:"time"`
}
