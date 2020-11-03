package controller

import (
	"net/http"
	"tugasmvc/app/model"
	"tugasmvc/app/utils"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type InDB struct {
	DB *gorm.DB
}
