package model

// DatoDummy model
type DatoDummy struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"nombre"`
	Time string `db:"time" json:"time"`
}
