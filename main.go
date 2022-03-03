package main

import (
	"BCC-Internship/config"
	"BCC-Internship/handler"
	"BCC-Internship/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	//tempat endpoint user
	_, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	r.POST("/user/register", handler.NewUserHandler)
	r.POST("/user/login", handler.UserLogin)
	// "/user/:id
	//tempat endpoint klinik
	r.GET("/clinic", middleware.CheckJwtUser(), handler.ReadClinic)
	r.POST("/clinic/register", handler.NewClinicalHandler)
	r.POST("/clinic/login", handler.ClinicLogin)
	r.Run()
}
