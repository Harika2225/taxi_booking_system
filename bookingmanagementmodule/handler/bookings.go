package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	eureka "com.example.bookingmanagement/eurekaregistry"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

var bookingTableName = "booking"

// Booking struct represents the booking data model
type Booking struct {
	ID            int    `json:"id"`
	CustomerID    int    `json:"customer_id"`
	DriverID      int    `json:"driver_id"`
	Pickupaddress string `json:"pickupaddress"`
	Destination   string `json:"destination"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Status        string `json:"status"`
	Amount        int    `json:"amount"`
}

func SetJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// CreateBooking handles the creation of a new booking request
// POST /api/createBooking
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var newBooking Booking
	fmt.Println(r.Body, "ppppppp")
	// Decode the JSON request body into the newBooking struct
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	newBooking.Status = "Pending"
	// Check if the Booking table exists, and migrate only if needed
	if !dbClient.Migrator().HasTable(&Booking{}) {
		fmt.Println("Migrating Booking table...")
		if err := dbClient.Table(bookingTableName).AutoMigrate(&Booking{}); err != nil {
			fmt.Println("Error migrating Booking table:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println("Booking table migration successful")
	}
	// Assuming your Booking struct has a Time field of type string
	var cost int

	// Parse the time string to extract the hour
	timeValue, err := time.Parse("15:04", newBooking.Time)
	if err != nil {
		// Try parsing in 12-hour format if 24-hour format fails
		timeValue, err = time.Parse("3:04 PM", newBooking.Time)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			http.Error(w, "Invalid time format", http.StatusBadRequest)
			return
		}
	}

	hour := timeValue.Hour()
	if (hour >= 11 && hour <= 17) || (hour >= 11 && hour <= 5+12) {
		cost = 1000
	} else {
		cost = 500
	}
	newBooking.Amount = cost

	// Create a new booking record
	fmt.Println("Creating a new booking record...")
	if err := dbClient.Table(bookingTableName).Create(&newBooking).Error; err != nil {
		fmt.Println("Error creating booking:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("New booking record created:", newBooking)

	// Return the newly created booking in the response
	SetJSONContentType(w)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBooking)

	logger.Info("this is booking")

	eureka.ClientCommunication(r, w, "drivermanagementmodule", "api/bookingStatus", newBooking)
	logger.Info("called booking")

	fmt.Println("Successfully communicated with drivermanagementmodule for api/bookingStatus")
}

func BookingAccepted(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------------", r.Body)
	var receivedData map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&receivedData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close() //close reqest body after decoding

	fmt.Println("--------------------", receivedData)

	driverID, _ := receivedData["DriverID"].(float64)
	bookingID, _ := receivedData["ID"].(float64)
	status, _ := receivedData["Status"].(string)

	fmt.Println("DriverID:", driverID)
	fmt.Println("BookingID:", bookingID)
	fmt.Println("Status:", status)

	updateBooking := Booking{
		ID:       int(bookingID),
		DriverID: int(driverID),
		Status:   status,
	}

	err := dbClient.Table(bookingTableName).Model(&Booking{}).Where("id = ?", updateBooking.ID).UpdateColumns(updateBooking).Error
	if err != nil {
		fmt.Println("Error updating booking:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Retrieve the complete updated record from the database
	var updatedRecord Booking
	err = dbClient.Table(bookingTableName).Where("id = ?", updateBooking.ID).First(&updatedRecord).Error
	if err != nil {
		fmt.Println("Error retrieving updated booking record:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Booking record updated:", updatedRecord)
	
	SetJSONContentType(w)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedRecord)
	
	eureka.ClientCommunication(r, w, "customermanagementmodule", "api/booked", updatedRecord)
	fmt.Println("Successfully communicated with customermanagementmodule for api/booked")
}

// GetBookings retrieves a list of all bookings
// GET /api/getBookings
func GetBookings(w http.ResponseWriter, r *http.Request) {
	var bookings []Booking
	e := dbClient.Table(bookingTableName).Find(&bookings)
	if e.Error != nil {
		return
	}
	fmt.Println(bookings)
	SetJSONContentType(w)
	json.NewEncoder(w).Encode(bookings)
}

// GetBookingByID retrieves a specific booking's details by its ID
// GET /api/getBookingById/{id}
func GetBookingByID(w http.ResponseWriter, r *http.Request) {
	id := parseBookingID(w, r)

	// Retrieve the booking from the database by its ID
	var booking Booking
	if err := dbClient.Table(bookingTableName).First(&booking, id).Error; err != nil {
		// Booking with the given ID not found
		http.NotFound(w, r)
		return
	}

	SetJSONContentType(w)
	json.NewEncoder(w).Encode(booking)
}

// UpdateBookingByID updates an existing booking's details by its ID
// PUT /api/updateBookingById/{id}
func UpdateBookingByID(w http.ResponseWriter, r *http.Request) {
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
	SetJSONContentType(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingBooking)
}

// DeleteBookingByID deletes a specific booking by its ID
// DELETE /api/deleteBookingById/{id}
func DeleteBookingByID(w http.ResponseWriter, r *http.Request) {
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
