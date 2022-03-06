package user

import "gorm.io/gorm"

type StatusPasien struct {
	gorm.Model
	Status    bool
	IDPayment uint
}
