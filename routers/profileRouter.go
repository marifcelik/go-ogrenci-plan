package routers

import (
	"go-ders-programi/db"
	"go-ders-programi/middlewares"
	"go-ders-programi/models"
	"go-ders-programi/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupProfileRouter(app *echo.Echo) {
	userRouter := app.Group("/profile")
	userRouter.Use(middlewares.AuthMiddleware)

	userRouter.GET("/", getProfile)
	userRouter.PATCH("/", updateProfile)
}

func getProfile(c echo.Context) error {
	var student models.Student
	student.ID = util.GetUserId(c)

	result := db.DB.First(&student)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusOK, student)
}

func updateProfile(c echo.Context) error {
	var student models.Student
	student.ID = util.GetUserId(c)

	result := db.DB.First(&student)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": result.Error.Error()})
	}

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	result = db.DB.Save(&student)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusOK, result.RowsAffected)
}
