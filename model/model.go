package model

import (
	"gorm.io/gorm"
)

var SecretKey = []byte("sasjdakdlkasjk")

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Position string
}
