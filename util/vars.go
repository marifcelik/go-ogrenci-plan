package util

import (
	"os"
	"strconv"
	"time"
)

func GetDBUrl() string {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "root:root@tcp(localhost:3306)/project?charset=utf8mb4&parseTime=True&loc=Local"
	}
	return connStr
}

func GetJWTSecret() string {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "jwt_super_secret_and_super_long_secret_key"
	}
	return jwtSecret
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}

func GetJWTExprMinutes() time.Duration {
	expirationMinutesStr := os.Getenv("JWT_EXPIRATION_MINUTES")
	expirationMinutes, err := strconv.Atoi(expirationMinutesStr)
	if err != nil {
		expirationMinutes = 5
	}
	return time.Duration(expirationMinutes) * time.Minute
}
