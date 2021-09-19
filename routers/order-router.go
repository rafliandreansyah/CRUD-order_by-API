package routers

import (
	"assignment-2/controllers"
	"net/http"

	"gorm.io/gorm"
)

func Router(db *gorm.DB) {

	DBConn := &controllers.DBConn{DB: db}

	http.HandleFunc("/orders", DBConn.Order)
	http.HandleFunc("/orders/", DBConn.DeleteOrder)

}
