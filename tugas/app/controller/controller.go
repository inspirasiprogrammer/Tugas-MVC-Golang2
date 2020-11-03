package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (inDB *InDB) CreateAccount(c *gin.Context) {

}
