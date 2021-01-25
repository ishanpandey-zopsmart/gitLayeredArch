package main

import (
	"database/sql"
	"log"
	"net/http"

	"layered/architecture/delivery"
	"layered/architecture/service"
	"layered/architecture/store"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	var db, dbErr = sql.Open("mysql", "root:password@/customer_service")
	if dbErr != nil {
		panic(dbErr)
	}
	defer db.Close()

	dbErr = db.Ping()
	if dbErr != nil {
		panic(dbErr.Error()) // proper error handling instead of panic in your app
	}
	r := mux.NewRouter()
	datastore := store.New(db)
	s := service.New(datastore)
	handler := delivery.New(s)

	r.HandleFunc("/customer", handler.GetByName).Queries("name", "{name}").Methods(http.MethodGet)
	r.HandleFunc("/customer", handler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id:[0-9]+}", handler.GetById).Methods(http.MethodGet)

	r.HandleFunc("/customer", handler.PostCustomer).Methods(http.MethodPost)
	r.HandleFunc("/customer/{id}", handler.PutCustomer).Methods(http.MethodPut)
	r.HandleFunc("/customer/{id}", handler.DeleteCustomer).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8000", r))
}
