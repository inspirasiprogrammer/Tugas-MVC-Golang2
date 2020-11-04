package main

import (
	"tugasmvc2/app/controllers"
	"tugasmvc2/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/api/v1/account/add", controllers.CreateAccount)
	router.POST("/api/v1/login", controllers.Login)
	router.GET("/api/v1/account", middleware.Auth, controllers.GetAccount)
	router.POST("/api/v1/transfer", middleware.Auth, controllers.Transfer)
	router.POST("/api/v1/withdraw", middleware.Auth, controllers.Withdraw)
	router.POST("/api/v1/deposit", middleware.Auth, controllers.Deposit)
	router.POST("/api/v1/interest", middleware.Auth, controllers.Interest)
	router.Run(":8080")
}
