package endpoint

import (
	"net/http"
	"strconv"

	// "golang-echo-middleware/internal/app/service"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysUntil2027() int
}

type Endpoint struct {
	s Service
}


func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}


func (e *Endpoint) DaysUntil(ctx echo.Context) error {
	// ctx.String(http.StatusOK, "{\"message\": \"Days until 2027\"}")

	daysUntil2027 := e.s.DaysUntil2027()

	responseJSON := `
		{
			"message": "Days until 2027: ` + strconv.Itoa(daysUntil2027) + `",
			"another_field": "This is another field"
		}
	`
	ctx.Blob(http.StatusOK, "application/json", []byte(responseJSON))

	return nil
}