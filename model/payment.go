package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserId       uint
	Status       bool
	JenisHewan   string
	Keluhan      string
	Ras          string
	JenisKelamin string
	Umur         int
	Tanggal      string
}

type PaymentInput struct {
	UserId       uint   `json:"user_id" binding:"required"`
	Status       bool   `json:"status" binding:"required"`
	JenisHewan   string `json:"jenis_hewan" binding:"required"`
	Keluhan      string `json:"keluhan" binding:"required"`
	Ras          string `json:"ras" binding:"required"`
	JenisKelamin string `json:"jenis_kelamin" binding:"required"`
	Umur         int    `json:"umur" binding:"required"`
	Tanggal      string `json:"tanggal" binding:"required"`
}
