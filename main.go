package main

import (
	"app/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Ejemplo con colores
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "\x1b[90m[${time_rfc3339}]\x1b[0m [${method} ${uri}][status=${status}]\n",
	// }))

	e.Use(middleware.Recover())

	// Listar Usuarios
	e.GET("/users", api.GetUsers)

	// Borrar Usuario
	e.DELETE("/user/:id", api.DeleteUser)

	// Agregar Usuario
	e.POST("/user", api.PostUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
