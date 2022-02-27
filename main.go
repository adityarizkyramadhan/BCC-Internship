package main

import (
	"BCC-Internship/config"
	"BCC-Internship/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//tempat endpoint user
	_, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	r.POST("/user/register", handler.NewUserHandler)
	r.POST("/user/login", handler.UserLogin)
	//tempat endpoint klinik
	r.GET("/clinic", handler.ReadClinic)
	r.POST("/clinic/register", handler.NewClinicalHandler)
	r.POST("/clinic/login", handler.ClinicLogin)
	r.Run()
}
