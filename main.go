package main

import (
	"BCC-Internship/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//tempat endpoint user
	r.POST("/user/new", handler.NewUserHandler)
	r.POST("/user/login", handler.UserLogin)
	//tempat endpoint klinik
	r.POST("/clinic/new", handler.NewClinicalHandler)
	r.POST("/clinic/login", handler.ClinicLogin)
	r.Run()
}
