package controllers

import "gorm.io/gorm"

type DBConn struct {
	DB *gorm.DB
}
