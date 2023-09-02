package routers

import (
	"go-ders-programi/db"
	"go-ders-programi/models"
	"go-ders-programi/util"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func SetupAuthRouter(app *echo.Echo) {
	authRouter := app.Group("/auth")

	authRouter.POST("/signin", handleSignin)
	authRouter.POST("/signup", handleSignup)
	authRouter.POST("/signout", handleSignout)
}

func handleSignin(c echo.Context) error {
	student := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&student); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if len(student.Password) < 8 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "password must be at least 8 characters"})
	}

	var studentDB models.Student
	db.DB.First(&studentDB, "username = ?", student.Username)
	if studentDB.ID == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "username does not exist"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(studentDB.Password), []byte(student.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": studentDB.ID,
		"exp": time.Now().Add(util.GetJWTExprMinutes()).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = tokenStr
	cookie.Expires = time.Now().Add(util.GetJWTExprMinutes())
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Path = "/"
	cookie.Secure = false
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

func handleSignup(c echo.Context) error {
	student := models.Student{}

	if c.Bind(&student) != nil {
		return echo.ErrBadRequest
	}

	if len(student.Password) < 8 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "password must be at least 8 characters"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "internal server error on password hashing")
	}

	student.Password = string(hash)
	result := db.DB.Table("students").Create(&student)
	if result.Error != nil {
		if strings.HasPrefix(result.Error.Error(), "Error 1062") {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "username already exists"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal server error on database"})
	}

	return c.JSON(http.StatusOK, student)
}

func handleSignout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(time.Minute * -1)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}
