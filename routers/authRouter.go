package routers

import (
	"go-ders-programi/db"
	"go-ders-programi/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type jsonRes map[string]interface{}

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
		return c.JSON(http.StatusBadRequest, jsonRes{"error": "password must be at least 8 characters"})
	}

	var studentDB models.Student
	db.DB.First(&studentDB, "username = ?", student.Username)
	if studentDB.ID == 0 {
		return c.JSON(http.StatusBadRequest, jsonRes{"error": "username does not exist"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(studentDB.Password), []byte(student.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, jsonRes{"error": "invalid password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		Subject:   strconv.Itoa(int(studentDB.ID)),
		Issuer:    "go-ders-programi",
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsonRes{"error": err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = tokenStr
	cookie.Expires = time.Now().Add(time.Minute * 1)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, jsonRes{"status": "ok"})
}

func handleSignup(c echo.Context) error {
	student := struct {
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if c.Bind(&student) != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if len(student.Password) < 8 {
		return c.JSON(http.StatusBadRequest, jsonRes{"error": "password must be at least 8 characters"})
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
			return c.JSON(http.StatusBadRequest, jsonRes{"error": "username already exists"})
		}
		return c.JSON(http.StatusInternalServerError, jsonRes{"error": "internal server error on database"})
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

	return c.JSON(http.StatusOK, jsonRes{"status": "ok"})
}
