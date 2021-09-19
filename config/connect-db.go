package config

import (
	"assignment-2/structs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	username  = "root"
	password  = "amaterasu"
	host      = "127.0.0.1:3306"
	dbname    = "orders_by"
	parsetime = "true"
)

func InitDB() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%s&loc=Local", username, password, host, dbname, parsetime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect db =>", err.Error())
		return nil
	}

	db.AutoMigrate(&structs.Order{}, &structs.Item{})

	return db

}
