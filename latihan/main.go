package main

import (
	"tugasmvc/app/config"
	"tugasmvc/app/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", inDB.CreateAccount)
	router.Run(":8080")
}
