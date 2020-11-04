package main

import (
	"tugasmvc2/app/controllers"
	"tugasmvc2/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// add CORS
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"GET","POST"}
	cfg.AllowHeaders = []string{"Authorization","Origin","Accept","X-Requested-With"," Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	router.Use(cors.New(cfg))

	//router.Use(cors.Default())
	router.POST("/api/v1/account/add", controllers.CreateAccount)
	router.POST("/api/v1/login", controllers.Login)
	router.GET("/api/v1/account", middleware.Auth, controllers.GetAccount)
	router.POST("/api/v1/transfer", middleware.Auth, controllers.Transfer)
	router.POST("/api/v1/withdraw", middleware.Auth, controllers.Withdraw)
	router.POST("/api/v1/deposit", middleware.Auth, controllers.Deposit)
	router.POST("/api/v1/interest", middleware.Auth, controllers.Interest)
	router.Run(":8080")
}
