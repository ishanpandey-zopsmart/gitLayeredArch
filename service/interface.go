package service

import (
	"layered/architecture/entities"
	"net/http"
)

type Customer interface {
	// Service layer decides the arguments needed to implement DB transactions in Store layer.
	GetByID(w http.ResponseWriter, id int)
	GetByName(w http.ResponseWriter, name string)
	CreateCustomer(w http.ResponseWriter, cust entities.Customer)
	UpdateCustomer(w http.ResponseWriter, id int, customer entities.Customer)
	DeleteCustomer(w http.ResponseWriter, id int)
}
