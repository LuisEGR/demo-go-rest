package api

import (
	"fmt"
	"net/http"

	"app/db"
	"app/model"

	"github.com/labstack/echo/v4"
)

// PostUser Handler
func PostUser(c echo.Context) error {

	dbCon := db.New()

	user := model.Usuario{}

	c.Bind(&user)

	println("Registrando: " + user.Name)
	tx := dbCon.MustBegin()

	_, e := tx.NamedExec("INSERT INTO usuario(id, name) VALUES (:id, :name)", user)
	if e != nil {
		fmt.Println("E", e)
		return c.String(400, "Error en base de datos")
	}

	err := tx.Commit()
	if err != nil {
		fmt.Println(err)
		return c.String(400, "Error en base de datos")
	}

	dbCon.Close()
	return c.JSON(http.StatusCreated, user)
}
