package middleware

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userRole := ctx.Request().Header.Get("User-Role")
		userRole = strings.ToLower(userRole)

		if userRole == roleAdmin {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil

	}
}