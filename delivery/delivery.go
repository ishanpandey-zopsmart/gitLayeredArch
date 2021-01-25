package delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"layered/architecture/entities"
	"layered/architecture/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func New(customer service.CustomerService) CustomerHandler {
	return CustomerHandler{service: customer}
}

func (c CustomerHandler) GetById(w http.ResponseWriter, r *http.Request) {
	pathparams := mux.Vars(r)
	id, err := strconv.Atoi(pathparams["id"])
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	resp, err := c.service.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]entities.Customer(nil)) // Empty struct for customer.
		return
	} else {
		if resp.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			//	json.NewEncoder(w).Encode([]entities.Customer(nil))
		} else {
			json.NewEncoder(w).Encode(resp)
		}
	}

}

func (c CustomerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp, err := c.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}

func (c CustomerHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name, ok := q["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(name[0])
	resp, err := c.service.GetByName(name[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (c CustomerHandler) PostCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var cust entities.Customer
	err = json.Unmarshal(body, &cust)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//w.Write([]byte("Invalid JSON format for insertion."))
		return
	}
	if cust.Name == "" || cust.DOB == "" || cust.Address.StreetName == "" || cust.Address.City == "" || cust.Address.State == "" {
		w.WriteHeader(http.StatusBadRequest)
		//w.Write([]byte("Missing Field Values."))
		return
	}

	// Calling the service.Create -> Calculates age and returns error if age is less.
	resp, age, _ := c.service.Create(cust)
	if age < 18 {
		w.WriteHeader(http.StatusMethodNotAllowed)
		//w.Write([]byte("Sorry, Age is less than 18."))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//w.Write([]byte("Oh snap, Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//
}

func (c CustomerHandler) PutCustomer(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)["id"]
	id, err := strconv.Atoi(pathParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//w.Write([]byte("Invalid ID as parameter query"))
		return
	}
	var customer entities.Customer
	bodyData, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(bodyData, &customer)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//	w.Write([]byte("Data format is not correct"))
		return
	}
	// If some field is empty, then return
	if customer.ID != 0 || customer.DOB != "" {
		w.WriteHeader(http.StatusBadRequest)
		//	w.Write([]byte("You can't update your ID or DOB."))
		return
	}
	if customer.Name == "" && customer.Address.State == "" && customer.Address.City == "" && customer.Address.StreetName == "" {
		w.WriteHeader(http.StatusBadRequest)
		//	w.Write([]byte("No data passed to Update, Failed."))
		return
	}

	res, err := c.service.Update(id, customer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//	w.Write([]byte("Something went wrong in query."))
		return
	}

	if res.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		//	w.Write([]byte("This record is not found in our database."))
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (c CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)["id"]

	id, err := strconv.Atoi(pathParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//	w.Write([]byte("This ID does not exist."))
		return
	}

	_, err = c.service.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//	w.Write([]byte("Internal Server Error."))
		return
	}

	//if customer == entities.Customer(nil) {
	//	w.WriteHeader(http.StatusNotFound)
	//	w.Write([]byte("This record is not found in our database."))
	//	return
	//}

	w.WriteHeader(http.StatusNoContent)
	//	w.Write([]byte("Data corresponding to passed ID is deleted successfully."))
}
