package main

import (
	"BCC-Internship/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/user/newuser", handler.NewUserHandler)
	r.Run()
}
