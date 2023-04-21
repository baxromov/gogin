package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password    string `gorm:"not null"`
	LastLogin   string `gorm:"not null"`
	IsSuperuser bool   `gorm:"not null"`
	Username    string `gorm:"unique;not null"`
	LastName    string `gorm:"not null"`
	Email       string `gorm:"not null"`
	IsStaff     bool   `gorm:"not null"`
	IsActive    bool   `gorm:"not null"`
	DateJoined  string `gorm:"not null"`
	FirstName   string `gorm:"not null"`
}
