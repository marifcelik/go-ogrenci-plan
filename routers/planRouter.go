package routers

import (
	"go-ders-programi/db"
	"go-ders-programi/middlewares"
	"go-ders-programi/models"
	"go-ders-programi/util"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetupPlanRouter(app *echo.Echo) {
	planRouter := app.Group("/plan")
	planRouter.Use(middlewares.AuthMiddleware)

	planRouter.GET("/", getAllPlans)
	planRouter.GET("/:id", getPlan, middlewares.CompareUserId)
	planRouter.POST("/", createPlan)
	planRouter.PATCH("/:id", updatePlan, middlewares.CompareUserId)
	planRouter.DELETE("/:id", deletePlan, middlewares.CompareUserId)
}

func getAllPlans(c echo.Context) error {
	var plans []models.Plan
	weekFilter := util.ParseUint(c.QueryParam("week"))
	monthFilter := util.ParseUint(c.QueryParam("month"))

	if weekFilter != 0 {
		db.DB.Where("student_id = ? AND start BETWEEN ? AND ?", util.GetUserId(c), time.Now(), time.Now().AddDate(0, 0, 7*int(monthFilter))).Find(&plans)
	} else if monthFilter != 0 {
		db.DB.Where("student_id = ? AND start BETWEEN ? AND ?", util.GetUserId(c), time.Now(), time.Now().AddDate(0, int(monthFilter), 0)).Find(&plans)
	} else {
		db.DB.Where("student_id = ?", util.GetUserId(c)).Find(&plans)
	}

	return c.JSON(http.StatusOK, plans)
}

func getPlan(c echo.Context) error {
	var plan models.Plan
	db.DB.First(&plan, "id = ?", c.Param("id"))
	return c.JSON(http.StatusOK, plan)
}

func createPlan(c echo.Context) error {
	plan := models.Plan{
		StudentId: util.GetUserId(c),
	}

	if err := c.Bind(&plan); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if plan.Start.IsZero() {
		plan.Start = time.Now().Truncate(time.Minute)
	}
	if plan.End.IsZero() {
		plan.End = plan.Start.Add(time.Hour * 2)
	}

	// bu kısım planın başlangıç ve bitiş tarihleri arasında başka bir plan var mı kontrol eder
	var count int64
	db.DB.Model(&models.Plan{}).
		Where("student_id = ?", plan.StudentId).
		Where("(? BETWEEN start AND end OR ? BETWEEN start AND end OR start BETWEEN ? AND ?)",
			plan.Start, plan.End, plan.Start, plan.End).
		Count(&count)

	// eğer varsa hata döndürür
	if count > 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "there is already a plan in this time"})
	}

	result := db.DB.Create(&plan)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, plan.ID)
}

func updatePlan(c echo.Context) error {
	var plan models.Plan
	result := db.DB.First(&plan, "id = ?", c.Param("id"))
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, result.Error)
	}

	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	result = db.DB.Save(&plan)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusOK, result.RowsAffected)
}

func deletePlan(c echo.Context) error {
	result := db.DB.Delete(&models.Plan{}, c.Param("id"))
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}
