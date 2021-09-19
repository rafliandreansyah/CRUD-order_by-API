package controllers

import (
	"assignment-2/structs"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func (Conn *DBConn) Order(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		getOrder(Conn, w, r)
	case "POST":
		postOrder(Conn, w, r)
	case "PATCH":
		updateOrder(Conn, w, r)
	default:
		notFound(w, r)
	}

}

/*
	Delete order
*/
func (Conn *DBConn) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {

		w.Header().Set("Content-Type", "application/json")

		var (
			order    structs.Order
			response structs.Response
		)

		orderId := r.URL.Path[len("/orders/"):]
		log.Println("order id->", orderId)

		result := Conn.DB.Delete(&order, orderId)
		log.Println("Delete result =>", result)

		response.Message = "Delete Success"
		response.Data = order

		jsonResponse, err := json.Marshal(&response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(jsonResponse))

	}
}

/*
	Create order
*/
func postOrder(Conn *DBConn, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var (
		order    structs.Order
		response structs.Response
	)

	// Decode r.body to struct
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		if err == io.EOF {
			msgError := "Request body must only contain a single JSON object"
			http.Error(w, msgError, http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := Conn.DB.Create(&order)
	log.Println("Created result =>", result)

	response.Message = "Ordered success"
	response.Data = order

	jsonData, err := json.Marshal(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(jsonData))

}

/*
	Get data order and each items
*/
func getOrder(Conn *DBConn, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var orders []structs.Order

	result := Conn.DB.Preload("Item").Find(&orders)
	log.Println("Get result =>", result)

	jsonData, err := json.Marshal(&orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonData))

}

/*
	Update data order
*/
func updateOrder(Conn *DBConn, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var (
		order    structs.Order
		response structs.Response
	)

	// Decode r.body to struct
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := Conn.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)
	log.Println("Update result =>", result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Data Updated"
	response.Data = &order

	jsonData, err := json.Marshal(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonData))

}

/*
	If request method not found handle on this function
*/

func notFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, `{"message": "Request method not allowed"}`)

}
