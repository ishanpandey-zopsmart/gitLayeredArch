package service

import "C"
import (
	"layered/architecture/entities"
	"layered/architecture/store"
)

type CustomerService struct {
	store store.Customer
}

func New(customer store.Customer) CustomerService {
	return CustomerService{store: customer}
}

func (c CustomerService) GetByID(id int) (entities.Customer, error){
	resp, err := c.store.GetByID(id)
	return resp, err
}

func (c CustomerService) GetAll() ([]entities.Customer, error){
	resp, err := c.store.GetAll()
	return resp, err
}

func (c CustomerService) GetByName(name string) ([]entities.Customer, error){
	resp, err := c.store.GetByName(name)
	return resp, err
}

func (c CustomerService) Create(cust entities.Customer) (entities.Customer, int, error){

	 if timestamp := DateSubstract(cust.DOB); timestamp/(3600*24*12*30) < 18 {
	 	return entities.Customer{}, timestamp/(3600*24*12*30), nil
	} else {
		res, err := c.store.Create(cust)
		return res, timestamp/(3600*24*12*30), err
	}
}

func (c CustomerService) Update(id int, customer entities.Customer) (entities.Customer, error){
		cust, err := c.store.Update(id, customer)
		return cust, err
}

func (c CustomerService) Delete(id int) (entities.Customer, error){
	customer, err := c.store.Delete(id)
	return customer, err
}
