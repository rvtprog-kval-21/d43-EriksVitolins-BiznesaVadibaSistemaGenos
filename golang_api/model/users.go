package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	Avatar    string
	Name      string
	LastName  string
	About     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
