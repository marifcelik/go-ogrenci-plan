package routers

import (
	"github.com/labstack/echo/v4"
)

func SetupPlanRouter(app *echo.Echo) {
	planRouter := app.Group("/plan")

	planRouter.GET("/", getAllPlans)
	planRouter.GET("/:id", getPlan)
	planRouter.POST("/", createPlan)
	planRouter.PATCH("/:id", updatePlan)
	planRouter.DELETE("/:id", deletePlan)
}

func getAllPlans(c echo.Context) error {
	return c.String(200, "getAllPlans")
}

func getPlan(c echo.Context) error {
	return nil
}

func createPlan(c echo.Context) error {
	return nil
}

func updatePlan(c echo.Context) error {
	return nil
}

func deletePlan(c echo.Context) error {
	return nil
}
