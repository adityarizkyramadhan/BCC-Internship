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
	r.Static("/image", "./image")
	r.POST("/user/register", handler.NewUserHandler)
	r.POST("/user/login", handler.UserLogin)
	r.POST("/user/:idclinic/payment", middleware.CheckJwtUser(), handler.Payment)
	r.POST("/user/payment/:idtransaction/upload", middleware.CheckJwtUser(), handler.UploadStructPayment)
	r.GET("/user/seeclinic", middleware.CheckJwtUser(), handler.ReadClinic)
	r.POST("/clinic/register", handler.NewClinicalHandler)
	r.GET("/clinic/seepayment", middleware.CheckJwtClinic(), handler.GetAllPaymentClinic)
	r.POST("/clinic/:idpayment/update", middleware.CheckJwtClinic(), handler.UpdatePayment)
	r.POST("/clinic/login", handler.ClinicLogin)
	r.Run()
}
