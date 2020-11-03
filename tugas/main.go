package main

import (
	"tugasmvc/app/handler"
	"tugasmvc/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/api/v1/account/add", handler.CreateAccount)
	router.POST("/api/v1/login", handler.Login)
	router.GET("/api/v1/account", middleware.Auth, handler.GetAccount)
	router.POST("/api/v1/transfer", middleware.Auth, handler.Transfer)
	router.POST("/api/v1/withdraw", middleware.Auth, handler.Withdraw)
	router.POST("/api/v1/deposit", middleware.Auth, handler.Deposit)
	r0uter.Run(":8080")
}
