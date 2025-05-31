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
	ctx.String(http.StatusOK, "My Response")

	return nil
}