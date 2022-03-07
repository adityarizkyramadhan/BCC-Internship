package model

import "gorm.io/gorm"

type Komunitas struct {
	gorm.Model
	NamaKomunitas string
	JenisHewan    string
	Deskripsi     string
	AsalKota      string
	LinkKomunitas string
}
