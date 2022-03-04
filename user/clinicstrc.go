package user

import "gorm.io/gorm"

type Clinic struct {
	gorm.Model
	NameClinic     string
	UsernameClinic string
	PasswordClinic string
	SpreadSheet    string
	Contact        string
	Address        string
}

type GetClinic struct {
	ID          uint
	NameClinic  string
	Contact     string
	Address     string
	SpreadSheet string
	Token       string
}

type PrintClinic struct {
	ID          uint
	NameClinic  string
	Contact     string
	Address     string
	SpreadSheet string
}

type NewClinic struct {
	NameClinic     string `json:"nameClinic" binding:"required"`
	UsernameClinic string `json:"usernameClinic" binding:"required,min=8,max=20"`
	PasswordClinic string `json:"passwordClinic" binding:"required,min=8,max=20"`
	SpreadSheet    string `json:"spreadSheet" binding:"required"`
	Contact        string `json:"contact" binding:"required"`
	Address        string `json:"address" binding:"required"`
}

type ClinicLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ClinicMasuk struct {
	ID          uint
	NameClinic  string
	Contact     string
	Address     string
	SpreadSheet string
}
