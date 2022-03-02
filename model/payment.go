package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserId       uint
	Amount       uint
	Status       bool
	LinkPayment  string
	JenisHewan   string
	Keluhan      string
	Ras          string
	JenisKelamin string
	Umur         int
	Tanggal      string
}
