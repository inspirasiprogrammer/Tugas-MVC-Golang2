package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("root:@/digibank?charset=utf8&parseTime=True&loc=Local")), &gorm.Config{})
	if err != nil {
		panic("failed to connect to databases " + err.Error())
	}
	DB.AutoMigrate(new(Account), new(Transaction))
}
