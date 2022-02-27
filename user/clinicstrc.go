package user

import "gorm.io/gorm"

type Clinic struct {
	gorm.Model
	NameClinic     string
	UsernameClinic string
	PasswordClinic string
	Contact        string
	Address        string
}

type GetClinic struct {
	ID         uint
	NameClinic string
	Contact    string
	Address    string
}

type NewClinic struct {
	NameClinic     string `json:"nameClinic" binding:"required"`
	UsernameClinic string `json:"usernameClinic" binding:"required"`
	PasswordClinic string `json:"passwordClinic" binding:"required"`
	Contact        string `json:"contact" binding:"required"`
	Address        string `json:"address" binding:"required"`
}

type ClinicLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
