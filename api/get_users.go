package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"app/db"
	"app/model"
)

const (
	host     = "localhost"
	port     = 5434
	user     = "admin"
	password = "admin123"
	dbname   = "admin"
)

// CREATE TABLE public.usuario (
// 	id numeric NULL,
// 	name varchar NULL,
// 	"time" timestamp NULL DEFAULT CURRENT_TIMESTAMP
// );

// INSERT INTO public.usuario (id, "name", "time") VALUES(0, 'user0', CURRENT_TIMESTAMP);
// INSERT INTO public.usuario (id, "name", "time") VALUES(1, 'user1', CURRENT_TIMESTAMP);
// INSERT INTO public.usuario (id, "name", "time") VALUES(2, 'user2', CURRENT_TIMESTAMP);

// GetUsers Handler
func GetUsers(c echo.Context) error {

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	// db, err := sqlx.Connect("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	dbCon := db.New()

	users := []model.DatoDummy{}
	dbCon.Select(&users, "SELECT * FROM public.usuario")
	if err != nil {
		fmt.Println(err)
		return c.String(400, "Error en base de datos")
	}

	return c.JSON(http.StatusOK, users)
	// return c.String(http.StatusOK, "Hello, World!")
}
