package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/micro/micro/v3/service/logger"
)

var customertableName = "customer"

// Customer struct represents the customer data model
type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName" gorm:"column:firstname"`
	LastName  string `json:"lastName" gorm:"column:lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

// Sample in-memory store for customer data
var customerStore = make(map[int]Customer)
var lastCustomerID = 0

// CreateCustomerHandler handles the creation of a new customer record
// POST /customer
func CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var newCustomer Customer
	json.NewDecoder(r.Body).Decode(&newCustomer)
	e := dbClient.Table(customertableName).Create(&newCustomer)
	if e != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	lastCustomerID++
	// logger.infof(lastCustomerID, "idddddd")
	newCustomer.ID = lastCustomerID
	customerStore[newCustomer.ID] = newCustomer

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

// GetCustomersHandler retrieves a list of all customers
// GET /customer
func GetCustomersHandler(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	e := dbClient.Table(customertableName).Find(&customers)
	if e.Error != nil {
		return
	}
	fmt.Print(customers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// GetCustomerByIDHandler retrieves a specific customer's details by its ID
// GET /customer/:id
func GetCustomerByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := parseCustomerID(w, r)
	customer, exists := customerStore[id]
	var customers []Customer
	e := dbClient.Table(customertableName).Find(&customers)
	if e.Error != nil {
		return
	}
	if !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// UpdateCustomerByIDHandler updates an existing customer's details by its ID
// PUT /customer/:id
func UpdateCustomerByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := parseCustomerID(w, r)
	var customers []Customer
	e := dbClient.Table(customertableName).Find(&customers)
	if e.Error != nil {
		return
	}
	_, exists := customerStore[id]
	if !exists {
		http.NotFound(w, r)
		return
	}

	var updatedCustomer Customer
	err := json.NewDecoder(r.Body).Decode(&updatedCustomer)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedCustomer.ID = id
	customerStore[id] = updatedCustomer

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCustomer)
}


// DeleteCustomerByIDHandler deletes a specific customer by its ID
// DELETE /customer/:id
func DeleteCustomerByIDHandler(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	e := dbClient.Table(customertableName).Find(&customers)
	if e.Error != nil {
		return
	}
	queryParams, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, "Error parsing query parameters", http.StatusBadRequest)
		return
	}
	idValues, ok := queryParams["id"]
	if !ok || len(idValues) == 0 {
		http.Error(w, "Missing or empty 'id' parameter", http.StatusBadRequest)
		return
	}
	id := idValues[0]
	result := dbClient.Table(customertableName).Where("id = ?", id).Delete(&customers)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Unable to delete Please Verify`))
		return
	}
	logger.Infof("Deleted Note with Id:" + id)
	json.NewEncoder(w).Encode(result.RowsAffected)
}

func parseCustomerID(w http.ResponseWriter, r *http.Request) int {
	idStr := r.URL.Path[len("/customer/"):]
	id, err := fmt.Sscanf(idStr, "%d")
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
	}
	return id
}
