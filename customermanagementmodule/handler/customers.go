package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
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
type Booking struct {
	ID            int    `json:"id"`
	CustomerID    int    `json:"customer_id"`
	DriverID      int    `json:"driver_id"`
	Pickupaddress string `json:"pickupaddress"`
	Destination   string `json:"destination"`
	Date          string `json:"date"`
	Status        string `json:"status"`
	Amount        int    `json:"amount"`
}

func SetJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// Handler for api/booked in customermanagementmodule
func BookedHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming data (booking details) from the request
	var receivedBooking Booking
	if err := json.NewDecoder(r.Body).Decode(&receivedBooking); err != nil {
	    http.Error(w, "Invalid request body", http.StatusBadRequest)
	    return
	}

	// Assuming you have a function to retrieve the booking status from the database
	// currentStatus, err := GetBookingStatus(receivedBooking.ID)
	// if err != nil {
	//     // Handle the error, log, or respond appropriately
	//     http.Error(w, "Failed to retrieve booking status", http.StatusInternalServerError)
	//     return
	// }

	// // Check the current status and respond accordingly
	// if currentStatus == "InProgress" {
	//     // Respond with a success message or the current booking details
	//     fmt.Fprintf(w, "Booking is in progress")
	// } else {
	//     // Respond with a message indicating that the booking is not in progress
	//     fmt.Fprintf(w, "Booking is not in progress")
	// }
}

// CreateCustomer handles the creation of a new customer record
// POST /customer
// Assuming you have a GORM dbClient initialized earlier
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var newCustomer Customer

	// Decode the JSON request body into the newCustomer struct
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the customer with the same email already exists
	var existingCustomer Customer
	if err := dbClient.Where("email = ?", newCustomer.Email).First(&existingCustomer).Error; err == nil {
		// Customer with the same email already exists, return existing customer details
		SetJSONContentType(w)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(existingCustomer)
		return
	}

	// Check if the customer table exists, and migrate only if needed
	if !dbClient.Migrator().HasTable(&Customer{}) {
		if err := dbClient.AutoMigrate(&Customer{}); err != nil {
			fmt.Println("Error creating the customer table:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Create a new customer record
	if err := dbClient.Create(&newCustomer).Error; err != nil {
		fmt.Println("Error creating customer:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the newly created customer in the response
	SetJSONContentType(w)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

// GetCustomers retrieves a list of all customers
// GET /customer
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	e := dbClient.Table(customertableName).Find(&customers)
	if e.Error != nil {
		return
	}
	fmt.Print(customers)
	SetJSONContentType(w)
	json.NewEncoder(w).Encode(customers)
}

// GetCustomerByID retrieves a specific customer's details by its ID
// GET /customer/:id
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id := parseCustomerID(w, r)
	fmt.Print(id)

	var customer Customer
	if err := dbClient.Table(customertableName).First(&customer, id).Error; err != nil {
		// Booking with the given ID not found
		http.NotFound(w, r)
		return
	}

	SetJSONContentType(w)
	json.NewEncoder(w).Encode(customer)
}

// UpdateCustomerByID updates an existing customer's details by its ID
// PUT /customer/:id
func UpdateCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	var updatedCustomer Customer
	if err := json.NewDecoder(r.Body).Decode(&updatedCustomer); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the customer with the given ID exists
	var existingCustomer Customer
	if err := dbClient.Table(customertableName).First(&existingCustomer, customerID).Error; err != nil {
		// Customer with the given ID not found
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}
	fmt.Println(updatedCustomer)
	// Update the existing customer with the new data
	if err := dbClient.Table(customertableName).Model(&existingCustomer).Updates(updatedCustomer).Error; err != nil {
		// Error updating customer
		fmt.Println("Error updating customer:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the updated customer in the response
	SetJSONContentType(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingCustomer)
}

// DeleteCustomerByID deletes a specific customer by its ID
// DELETE /customer/:id
func DeleteCustomerByID(w http.ResponseWriter, r *http.Request) {
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

func parseCustomerID(w http.ResponseWriter, r *http.Request) uint {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return 0
	}
	return uint(id)
}
