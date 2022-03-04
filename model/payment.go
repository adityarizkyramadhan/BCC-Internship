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
	SaveImage    SaveImage
}

type PaymentReturn struct {
	IDTransaction uint
	UserId        uint
	ClinicId      uint
	Status        bool
	JenisHewan    string
	Keluhan       string
	Ras           string
	JenisKelamin  string
	Umur          int
	Tanggal       string
}

type GetPayment struct {
	IDTransaction uint
	UserId        uint
	ClinicId      uint
	Status        bool
	JenisHewan    string
	Keluhan       string
	Ras           string
	JenisKelamin  string
	Umur          int
	LinkImage     string
}

type InputUriClinic struct {
	IdClinic int `uri:"idclinic" binding:"required"`
}

type PaymentInput struct {
	JenisHewan   string `json:"jenisHewan" binding:"required"`
	Keluhan      string `json:"keluhan" binding:"required"`
	Ras          string `json:"ras" binding:"required"`
	JenisKelamin string `json:"jenisKelamin" binding:"required"`
	Umur         int    `json:"umur" binding:"required"`
	Tanggal      string `json:"tanggal" binding:"required"`
}

type InputUriTransaction struct {
	IdTransaction int `uri:"idtransaction" binding:"required"`
}

type SaveImage struct {
	gorm.Model
	PaymentId uint
	Path      string
}

type ReturnImage struct {
	TransactionID uint
	Path          string
}
