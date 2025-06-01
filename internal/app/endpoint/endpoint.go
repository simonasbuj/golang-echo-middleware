package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"

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

type Response struct {
	Message 		string `json:"message"`
	AnotherField 	string `json:"anotherField"`
}

func (e *Endpoint) DaysUntil(ctx echo.Context) error {

	daysUntil2027 := e.s.DaysUntil2027()

	response := Response{
		Message: fmt.Sprintf("Days until 2027: %d", daysUntil2027),
		AnotherField: "This is another field v2",
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to serialize JSON response",
			"oneMore": "hi",
		})
	}

	ctx.JSONBlob(http.StatusOK, jsonBytes)

	return nil
}