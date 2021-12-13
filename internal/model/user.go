package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Surname string
	Email   string
}

type UserDto struct {
	ID      uint
	Name    string
	Surname string
	Email   string
}
