package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

var bookingTableName = "booking"

// Booking struct represents the booking data model
type Booking struct {
	ID            int    `json:"id"`
	CustomerID    int    `json:"customer_id"`
	Pickupaddress string `json:"pickupaddress"`
	Destination   string `json:"destination"`
	Date          string `json:"date"`
	Status        string `json:"status"`
}

// CreateBookingHandler handles the creation of a new booking request
// POST /api/createBooking
func CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	var newBooking Booking
	fmt.Println(r.Body, "ppppppp")
	// Decode the JSON request body into the newBooking struct
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Check if the Booking table exists, and migrate only if needed
	if !dbClient.Migrator().HasTable(&Booking{}) {
		fmt.Println("Migrating Booking table...")
		if err := dbClient.AutoMigrate(&Booking{}); err != nil {
			fmt.Println("Error migrating Booking table:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println("Booking table migration successful")
	}

	// Create a new booking record
	fmt.Println("Creating a new booking record...")
	if err := dbClient.Table(bookingTableName).Create(&newBooking).Error; err != nil {
		fmt.Println("Error creating booking:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("New booking record created:", newBooking)

	// Return the newly created booking in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBooking)
}

// GetBookingsHandler retrieves a list of all bookings
// GET /api/getBookings
func GetBookingsHandler(w http.ResponseWriter, r *http.Request) {
	var bookings []Booking
	e := dbClient.Table(bookingTableName).Find(&bookings)
	if e.Error != nil {
		return
	}
	fmt.Println(bookings)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// GetBookingByIDHandler retrieves a specific booking's details by its ID
// GET /api/getBookingById/{id}
func GetBookingByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := parseBookingID(w, r)

	// Retrieve the booking from the database by its ID
	var booking Booking
	if err := dbClient.Table(bookingTableName).First(&booking, id).Error; err != nil {
		// Booking with the given ID not found
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

// UpdateBookingByIDHandler updates an existing booking's details by its ID
// PUT /api/updateBookingById/{id}
func UpdateBookingByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing booking ID", http.StatusBadRequest)
		return
	}

	bookingID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	// Check if the booking with the given ID exists
	var existingBooking Booking
	if err := dbClient.Table(bookingTableName).First(&existingBooking, bookingID).Error; err != nil {
		// Booking with the given ID not found
		fmt.Println("Booking not found")
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	var updatedBooking Booking
	if err := json.NewDecoder(r.Body).Decode(&updatedBooking); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update the existing booking with the new data
	if err := dbClient.Table(bookingTableName).Model(&existingBooking).Updates(updatedBooking).Error; err != nil {
		// Error updating booking
		fmt.Println("Error updating booking:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the updated booking in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingBooking)
}

// DeleteBookingByIDHandler deletes a specific booking by its ID
// DELETE /api/deleteBookingById/{id}
func DeleteBookingByIDHandler(w http.ResponseWriter, r *http.Request) {
	var bookings []Booking
	e := dbClient.Table(bookingTableName).Find(&bookings)
	if e.Error != nil {
		return
	}
	queryParams := r.URL.Query()
	idValues, ok := queryParams["id"]
	if !ok || len(idValues) == 0 {
		http.Error(w, "Missing or empty 'id' parameter", http.StatusBadRequest)
		return
	}
	id := idValues[0]
	result := dbClient.Table(bookingTableName).Where("id = ?", id).Delete(&bookings)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Unable to delete. Please verify`))
		return
	}
	logger.Infof("Deleted Booking with ID:" + id)
	json.NewEncoder(w).Encode(result.RowsAffected)
}

func parseBookingID(w http.ResponseWriter, r *http.Request) uint {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return 0
	}
	return uint(id)
}
