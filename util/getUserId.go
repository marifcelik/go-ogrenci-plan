package util

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserId(c echo.Context) uint {
	return ParseUint(c.Request().Header.Get("X-ID"))
}

func ParseUint(s string) uint {
	u64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint(u64)
}
