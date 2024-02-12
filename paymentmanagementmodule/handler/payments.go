package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

var paymentTableName = "payment"

// Payment struct represents the payment data model
type Payment struct {
	ID              int     `json:"id"`
	Amount          float64 `json:"amount"`
	PaymentDate     string  `json:"payment_date"`
	CustomerID      int     `json:"customer_id"`
	PaymentMethodID int     `json:"payment_method_id"`
}

// CreatePaymentHandler handles the creation of a new payment record
// POST /payment
func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var newPayment Payment
	err := json.NewDecoder(r.Body).Decode(&newPayment)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the Payment table exists and migrate only if needed
	if !dbClient.Migrator().HasTable(&Payment{}) {
		fmt.Println("Migrating Payment table...")
		if err := dbClient.AutoMigrate(&Payment{}); err != nil {
			fmt.Println("Error migrating Payment table:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println("Payment table migration successful")
	}

	// Create a new payment record
	fmt.Println("Creating a new payment record...")
	if err := dbClient.Table(paymentTableName).Create(&newPayment).Error; err != nil {
		fmt.Println("Error creating payment:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("New payment record created:", newPayment)

	// Return the newly created payment in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPayment)
}

// GetPaymentsHandler retrieves a list of all payments
// GET /payment
func GetPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	var payments []Payment
	if err := dbClient.Table(paymentTableName).Find(&payments).Error; err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

// GetPaymentByIDHandler retrieves a specific payment's details by its ID
// GET /payment/{id}
func GetPaymentByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := parsePaymentID(w, r)

	// Retrieve the payment from the database by its ID
	var payment Payment
	if err := dbClient.Table(paymentTableName).First(&payment, id).Error; err != nil {
		// Payment with the given ID not found
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

// UpdatePaymentByIDHandler updates an existing payment's details by its ID
// PUT /payment/{id}
func UpdatePaymentByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing payment ID", http.StatusBadRequest)
		return
	}

	paymentID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	// Check if the payment with the given ID exists
	var existingPayment Payment
	if err := dbClient.Table(paymentTableName).First(&existingPayment, paymentID).Error; err != nil {
		// Payment with the given ID not found
		fmt.Println("Payment not found")
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}

	var updatedPayment Payment
	if err := json.NewDecoder(r.Body).Decode(&updatedPayment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update the existing payment with the new data
	if err := dbClient.Table(paymentTableName).Model(&existingPayment).Updates(updatedPayment).Error; err != nil {
		// Error updating payment
		fmt.Println("Error updating payment:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the updated payment in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingPayment)
}

// DeletePaymentByIDHandler deletes a specific payment by its ID
// DELETE /payment/{id}
func DeletePaymentByIDHandler(w http.ResponseWriter, r *http.Request) {
	var payments []Payment
	if err := dbClient.Table(paymentTableName).Find(&payments).Error; err != nil {
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
	result := dbClient.Table(paymentTableName).Where("id = ?", id).Delete(&payments)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Unable to delete Please Verify`))
		return
	}
	fmt.Println("Deleted Payment with ID:", id)
	json.NewEncoder(w).Encode(result.RowsAffected)
}

func parsePaymentID(w http.ResponseWriter, r *http.Request) uint {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return 0
	}
	return uint(id)
}
