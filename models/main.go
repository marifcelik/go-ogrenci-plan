package models

import (
	"time"

	"gorm.io/gorm"
)

type PlanStatus uint

const (
	Todo PlanStatus = iota
	Done
	Cancelled
)

type Plan struct {
	gorm.Model
	StudentId uint       `json:"student_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Start     time.Time  `json:"start"`
	End       time.Time  `json:"end"`
	State     PlanStatus `json:"state" gorm:"type:int unsigned;default:0"`
}

type Student struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname" gorm:"not null"`
	Username string `json:"username" gorm:"unique; not null"`
	Password string `json:"password" gorm:"not null"`
}
