package service

import (
	"layered/architecture/entities"
)

type Customer interface {
	// Service layer decides the arguments needed to implement DB transactions in Store layer.
	GetByID(id int) (entities.Customer, error)
	GetByName(name string) ([]entities.Customer, error)
	GetAll() ([]entities.Customer, error)
	Create(cust entities.Customer) (entities.Customer, error)
	Update(id int, customer entities.Customer) (entities.Customer, error)
	Delete(id int) (entities.Customer, error)
}
