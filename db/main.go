package db

import (
	"go-ders-programi/models"
	"go-ders-programi/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(util.GetDBUrl()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Plan{}, &models.Student{})
}
