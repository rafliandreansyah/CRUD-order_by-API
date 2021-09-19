package main

import (
	"assignment-2/config"
	"assignment-2/routers"
	"net/http"
)

func main() {

	db := config.InitDB()

	routers.Router(db)

	http.ListenAndServe(":8080", nil)

}
