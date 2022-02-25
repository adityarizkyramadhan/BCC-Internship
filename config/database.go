package config

import (
	"BCC-Internship/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabases() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/klinikhewan?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&user.User{}, &user.Doctor{}, &user.Clinic{}, &user.Animal{})
	if err != nil {
		panic(err)
	}
	return db, err
}
