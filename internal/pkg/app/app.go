package app

import (
	"fmt"
	"log"

	"golang-echo-middleware/internal/app/endpoint"

	"github.com/labstack/echo/v4"
)


type App struct {
	e 		*endpoint.Endpoint
	echo 	*echo.Echo
}


func New() (*App, error) {
	a := &App{
		e: endpoint.New(),
		echo: echo.New(),
	}

	// Register routes
	a.echo.GET("/days-until-2027", a.e.DaysUntil)

	return a, nil
}



func (a *App) Run() error {
	log.Println("Server is running...")

	err := a.echo.Start(":8080")
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}