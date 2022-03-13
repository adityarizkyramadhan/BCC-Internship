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
	r.Static("/imagesPayment", "./imagesPayment")
	r.Static("/imageClinics", "./imageClinics")
	// set grup endpoint
	user := r.Group("/user")
	{
		user.POST("/register", handler.NewUserHandler)
		user.POST("/login", handler.UserLogin)
		user.POST("/:idclinic/payment", middleware.CheckJwtUser(), handler.Payment)
		user.POST("/payment/:idtransaction/upload", middleware.CheckJwtUser(), handler.UploadStructPayment)
		user.GET("/seeclinic", handler.ReadClinic)
		user.GET("/clinic/search", handler.SearchClinic)
		user.GET("/history", middleware.CheckJwtUser(), handler.GetHistory)
		user.GET("/community", middleware.CheckJwtUser(), komunitas.GetKomunitas)
		user.GET("/community/search", middleware.CheckJwtUser(), komunitas.SearchKomunitas)
	}
	//Klinik
	clinic := r.Group("/clinic")
	{
		clinic.POST("/register", handler.NewClinicalHandler)
		clinic.POST("/upload/:idclinic", handler.UploadClinicImage)
		clinic.POST("/login", handler.ClinicLogin)
		clinic.GET("/payment", middleware.CheckJwtClinic(), handler.GetAllPaymentClinic)
		clinic.POST("/:idpayment/validate", middleware.CheckJwtClinic(), handler.UpdatePayment)
		clinic.GET("/paymentsuccess", middleware.CheckJwtClinic(), handler.SeeValidatePayment)
		clinic.GET("/showinvoice", middleware.CheckJwtClinic(), handler.ShowInvoices)
		clinic.GET("/seepatient", middleware.CheckJwtClinic(), pasien.GetPasien)
		clinic.POST("/:idpayment/updatepatient", middleware.CheckJwtClinic(), pasien.UpdatePatient)
	}
	//komunitas
	r.POST("/petmate/komunitas", komunitas.AddKomunitas)
	r.Run(":5000")
}
