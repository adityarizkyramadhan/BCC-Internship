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
	Umur         string
	Tanggal      string
	Layanan      string
	Harga        string
	Jam          string
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
	Umur          string
	Tanggal       string
	Harga         string
	Jam           string
	Layanan       string
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
	Umur          string
	Tanggal       string
	Layanan       string
	Harga         string
	Jam           string
	SaveImage     SaveImage
}

type InputUriClinic struct {
	IdClinic int `uri:"idclinic" binding:"required"`
}

type PaymentInput struct {
	JenisHewan   string `json:"jenisHewan" binding:"required"`
	Keluhan      string `json:"keluhan" binding:"required"`
	Ras          string `json:"ras" binding:"required"`
	JenisKelamin string `json:"jenisKelamin" binding:"required"`
	Umur         string `json:"umur" binding:"required"`
	Tanggal      string `json:"tanggal" binding:"required"`
	Jam          string `json:"jam" binding:"required"`
	Layanan      string `json:"layanan" binding:"required"`
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

type InputUriPayment struct {
	IdPayment int `uri:"idpayment" binding:"required"`
}

type InputStatusPayment struct {
	Status bool `json:"status" binding:"required"`
}

type DataInvoice struct {
	IDTransaction uint
	UserId        uint
	ClinicId      uint
	Status        bool
}

type ReturnInvoice struct {
	BanyakPayment int
	TotalBayar    int
}
