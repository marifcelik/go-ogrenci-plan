package middlewares

import (
	"fmt"
	"go-ders-programi/util"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwtCookie, err := c.Cookie("jwt")
		if err != nil {
			return echo.ErrUnauthorized
		}

		token, err := jwt.Parse(jwtCookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(util.GetJWTSecret()), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Request().Header.Set("X-ID", strconv.Itoa(int(claims["sub"].(float64))))
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
	}
}
