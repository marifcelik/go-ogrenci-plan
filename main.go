package main

import (
	"net/http"
	"os"

	"go-ders-programi/db"
	"go-ders-programi/routers"
	"go-ders-programi/util"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if os.Getenv("APP_ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}
	db.InitDB()
}

func main() {
	app := echo.New()
	app.Pre(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusAccepted, "Hello World")
	})

	routers.SetupProfileRouter(app)
	routers.SetupPlanRouter(app)
	routers.SetupAuthRouter(app)

	app.Logger.Fatal(app.Start(":" + util.GetPort()))
}
