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
type Doctor struct {
	gorm.Model
	Name      string
	Age       int
	Address   string
	Education string
	ClinicId  uint
}

type NewClinic struct {
	NameClinic     string `json:"nameClinic" binding:"required"`
	UsernameClinic string `json:"usernameClinic" binding:"required"`
	PasswordClinic string `json:"passwordClinic" binding:"required"`
	Contact        string `json:"contact" binding:"required"`
	Address        string `json:"address" binding:"required"`
}

type Clinic struct {
	gorm.Model
	NameClinic     string
	UsernameClinic string
	PasswordClinic string
	Contact        string
	Address        string
}

type Animal struct {
	gorm.Model
	Name        string
	Age         int
	Description string
	OwnerId     uint
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ClinicLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type DoctorLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
