package middlewares

import "github.com/labstack/echo/v4"

func CompareUserId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("X-ID") != c.Param("id") {
			return echo.ErrForbidden
		}

		return next(c)
	}
}
