package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"app/db"
	"app/model"
)

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

	users := []model.Usuario{}
	err := dbCon.Select(&users, "SELECT * FROM public.usuario")
	if err != nil {
		fmt.Println(err)
		return c.String(400, "Error en base de datos")
	}
	dbCon.Close()

	return c.JSON(http.StatusOK, users)
	// return c.String(http.StatusOK, "Hello, World!")
}
