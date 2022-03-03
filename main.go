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
	_, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	r.POST("/user/register", handler.NewUserHandler)
	r.POST("/user/login", handler.UserLogin)
	r.POST("/user/:idclinic/payment", middleware.CheckJwtUser(), handler.Payment)
	r.GET("/clinic", middleware.CheckJwtUser(), handler.ReadClinic)
	r.POST("/clinic/register", handler.NewClinicalHandler)
	r.POST("/clinic/login", handler.ClinicLogin)
	r.Run()
}
