package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	//Docker
	//DB, err = gorm.Open(mysql.Open(fmt.Sprintf("irwan:masterkey@tcp(172.17.0.2)/digibank")), &gorm.Config{})
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("root:@tcp(127.0.0.1)/digibank")), &gorm.Config{})
	if err != nil {
		panic("failed to connect to databases " + err.Error())
	}
	DB.AutoMigrate(new(Account), new(Transaction))
}
