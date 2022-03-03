package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserId       uint
	ClinicId     uint
	Status       bool
	JenisHewan   string
	Keluhan      string
	Ras          string
	JenisKelamin string
	Umur         int
	Tanggal      string
}

type PaymentInput struct {
	ClinicId     uint   `json:"clinicId" binding:"required"`
	JenisHewan   string `json:"jenisHewan" binding:"required"`
	Keluhan      string `json:"keluhan" binding:"required"`
	Ras          string `json:"ras" binding:"required"`
	JenisKelamin string `json:"jenisKelamin" binding:"required"`
	Umur         int    `json:"umur" binding:"required"`
	Tanggal      string `json:"tanggal" binding:"required"`
}
