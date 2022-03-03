package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Contact  string
	Username string
	Password string
	Address  string
}

type NewUser struct {
	Name     string `json:"name" binding:"required"`
	Contact  string `json:"contact" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserMasuk struct {
	ID      uint
	Name    string
	Contact string
	Address string
}

type ReturnUser struct {
	ID      uint
	Name    string
	Contact string
	Address string
	Token   string
}
