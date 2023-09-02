package routers

import (
	"github.com/labstack/echo/v4"
)

func SetupProfileRouter(app *echo.Echo) {
	userRouter := app.Group("/profile")

	userRouter.GET("/", getProfile)

	userRouter.GET("/settings", func(c echo.Context) error {
		return c.String(200, "User Settings")
	})
}

func getProfile(c echo.Context) error {
	return c.String(200, "User Profile")
}
