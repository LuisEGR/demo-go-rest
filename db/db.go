package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 5434
	user     = "admin"
	password = "admin123"
	dbname   = "admin"
)

// New - Returns new DB instance
func New() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbCon, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	return dbCon
}
