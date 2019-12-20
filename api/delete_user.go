package api

import (
	"fmt"
	"net/http"

	"app/db"

	"github.com/labstack/echo/v4"
)

// DeleteUser Handler
func DeleteUser(c echo.Context) error {

	idUser := c.Param("id")
	fmt.Println("Eliminando: " + idUser)

	dbCon := db.New()

	tx := dbCon.MustBegin()

	_, e := tx.Exec("DELETE FROM usuario WHERE id=$1", idUser)
	if e != nil {
		fmt.Println("E", e)
		return c.String(400, "Error en base de datos")
	}

	err := tx.Commit()
	if err != nil {
		fmt.Println(err)
		return c.String(400, "Error en base de datos")
	}

	// return c.JSON(http.StatusOK, "user")

	return c.NoContent(http.StatusOK)
}
