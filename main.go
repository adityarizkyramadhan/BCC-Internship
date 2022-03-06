package main

import (
	"BCC-Internship/config"
	"BCC-Internship/handler"
	"BCC-Internship/komunitas"
	"BCC-Internship/middleware"
	"BCC-Internship/pasien"

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
	// User
	r.POST("/user/register", handler.NewUserHandler)
	r.POST("/user/login", handler.UserLogin)
	r.POST("/user/:idclinic/payment", middleware.CheckJwtUser(), handler.Payment)
	r.POST("/user/payment/:idtransaction/upload", middleware.CheckJwtUser(), handler.UploadStructPayment)
	r.GET("/user/seeclinic", middleware.CheckJwtUser(), handler.ReadClinic)
	r.GET("/user/history", middleware.CheckJwtUser(), handler.GetHistory)
	r.GET("/user/community", middleware.CheckJwtUser(), komunitas.GetKomunitas)
	r.GET("/user/community/search", middleware.CheckJwtUser(), komunitas.SearchKomunitas)

	//Klinik
	r.POST("/clinic/register", handler.NewClinicalHandler)
	r.POST("/clinic/login", handler.ClinicLogin)
	r.GET("/clinic/payment", middleware.CheckJwtClinic(), handler.GetAllPaymentClinic)
	r.POST("/clinic/:idpayment/validate", middleware.CheckJwtClinic(), handler.UpdatePayment)
	r.GET("/clinic/paymentsuccess", middleware.CheckJwtClinic(), handler.SeeValidatePayment)
	r.GET("/clinic/showinvoice", middleware.CheckJwtClinic(), handler.ShowInvoices)
	r.GET("/clinic/seepatient", middleware.CheckJwtClinic(), pasien.GetPasien)
	r.POST("/clinic/:idpayment/updatepatient", middleware.CheckJwtClinic(), pasien.UpdatePatient)
	//komunitas
	r.POST("/petmate/komunitas", komunitas.AddKomunitas)

	r.Run()
}
