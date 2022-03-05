package main

import (
	"BCC-Internship/config"
	"BCC-Internship/handler"
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
	r.POST("/user/register", handler.NewUserHandler)
	r.POST("/user/login", handler.UserLogin)
	r.POST("/user/:idclinic/payment", middleware.CheckJwtUser(), handler.Payment)
	r.POST("/user/payment/:idtransaction/upload", middleware.CheckJwtUser(), handler.UploadStructPayment)
	r.GET("/user/seeclinic", middleware.CheckJwtUser(), handler.ReadClinic)
	r.POST("/clinic/register", handler.NewClinicalHandler)
	r.POST("/clinic/login", handler.ClinicLogin)
	r.GET("/clinic/payment", middleware.CheckJwtClinic(), handler.GetAllPaymentClinic)
	r.POST("/clinic/:idpayment/validate", middleware.CheckJwtClinic(), handler.UpdatePayment)
	r.GET("/clinic/paymentsuccess", middleware.CheckJwtClinic(), handler.SeeValidatePayment)
	r.GET("/clinic/showinvoice", middleware.CheckJwtClinic(), handler.ShowInvoices)
	r.GET("/clinic/seepatient", middleware.CheckJwtClinic(), pasien.GetPasien)
	//Klinik payment home atau klinik payment
	//riwayat transaksi
	//komunitas
	r.Run()
}
