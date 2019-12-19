package api

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func GetFullName(c echo.Context) error {

	// Bind the input data to ExampleRequest
	// exampleRequest := new(model.ExampleRequest)
	// if err := c.Bind(exampleRequest); err != nil {
	//   return err
	// }

	// Manipulate the input data
	greeting := "Hola mundo"

	return c.JSONBlob(
		http.StatusOK,
		[]byte(
			fmt.Sprintf(`{
		  "first_name": %q,
		  "last_name": %q,
		  "msg": "Hello %s"
		}`, "Luis", "Gonzalez", greeting),
		),
	)
}
