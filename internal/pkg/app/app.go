package app

import (
	"fmt"
	"log"

	"golang-echo-middleware/internal/app/endpoint"
	"golang-echo-middleware/internal/app/service"
	"golang-echo-middleware/internal/app/middleware"

	"github.com/labstack/echo/v4"
)

type App struct {
	e 		*endpoint.Endpoint
	echo 	*echo.Echo
}

func New() (*App, error) {
	a := &App{
		e: endpoint.New(service.New()),
		echo: echo.New(),
	}

	// Register middleware
	a.echo.Use(middleware.CheckRole)

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
