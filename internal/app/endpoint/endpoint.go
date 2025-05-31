package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


type Endpoint struct {}


func New() *Endpoint {
	return &Endpoint{}
}


func (e *Endpoint) DaysUntil(ctx echo.Context) error {
	// ctx.String(http.StatusOK, "{\"message\": \"Days until 2027\"}")

	responseJSON := `
		{
			"message": "Days until 2027",
			"another_field": "This is another field"
		}
	`
	ctx.Blob(http.StatusOK, "application/json", []byte(responseJSON))

	return nil
}