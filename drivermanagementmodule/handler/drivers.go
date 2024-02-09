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

var driverTableName = "driver"

// Driver struct represents the driver data model
type Driver struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	License   string `json:"license"`
}

// CreateDriverHandler handles the creation of a new driver record
// POST /api/createDriver
func CreateDriverHandler(w http.ResponseWriter, r *http.Request) {
	var newDriver Driver
	// Decode the JSON request body into the newDriver struct
	err := json.NewDecoder(r.Body).Decode(&newDriver)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the Driver table exists, and migrate only if needed
	if !dbClient.Migrator().HasTable(&Driver{}) {
		fmt.Println("Migrating Driver table...")
		if err := dbClient.AutoMigrate(&Driver{}); err != nil {
			fmt.Println("Error migrating Driver table:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println("Driver table migration successful")
	}

	// Check if the driver with the same license already exists
	var existingDriver Driver
	if err := dbClient.Where("license = ?", newDriver.License).First(&existingDriver).Error; err == nil {
		// Driver with the same license already exists, return existing driver details
		fmt.Println("Driver with the same license already exists. Returning existing driver details.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(existingDriver)
		return
	}

	// Create a new driver record
	fmt.Println("Creating a new driver record...")
	if err := dbClient.Table(driverTableName).Create(&newDriver).Error; err != nil {
		fmt.Println("Error creating driver:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("New driver record created:", newDriver)

	// Check if the driver was stored in the database
	var retrievedDriver Driver
	if err := dbClient.First(&retrievedDriver, newDriver.ID).Error; err != nil {
		fmt.Println("Error retrieving driver from the database:", err)
	} else {
		fmt.Println("Driver retrieved from the database:", retrievedDriver)
	}

	// Return the newly created driver in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newDriver)
}

// GetDriversHandler retrieves a list of all drivers
// GET /api/getDriver
func GetDriversHandler(w http.ResponseWriter, r *http.Request) {
	var drivers []Driver
	e := dbClient.Table(driverTableName).Find(&drivers)
	if e.Error != nil {
		return
	}
	fmt.Println(drivers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drivers)
}

// GetDriverByIDHandler retrieves a specific driver's details by its ID
// GET /api/getDriverById/{id}
func GetDriverByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := parseDriverID(w, r)

	// Retrieve the driver from the database by its ID
	var driver Driver
	if err := dbClient.Table(driverTableName).First(&driver, id).Error; err != nil {
		// Driver with the given ID not found
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(driver)
}

// UpdateDriverByIDHandler updates an existing driver's details by its ID
// PUT /api/updateDriverById/{id}
func UpdateDriverByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing driver ID", http.StatusBadRequest)
		return
	}

	driverID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	// Check if the driver with the given ID exists
	var existingDriver Driver
	if err := dbClient.Table(driverTableName).First(&existingDriver, driverID).Error; err != nil {
		// Driver with the given ID not found
		fmt.Println("Driver not found")
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}

	var updatedDriver Driver
	if err := json.NewDecoder(r.Body).Decode(&updatedDriver); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update the existing driver with the new data
	if err := dbClient.Table(driverTableName).Model(&existingDriver).Updates(updatedDriver).Error; err != nil {
		// Error updating driver
		fmt.Println("Error updating driver:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the updated driver in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingDriver)
}

// DeleteDriverByIDHandler deletes a specific driver by its ID
// DELETE /api/deleteDriverById/{id}
func DeleteDriverByIDHandler(w http.ResponseWriter, r *http.Request) {
	var drivers []Driver
	e := dbClient.Table(driverTableName).Find(&drivers)
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
	result := dbClient.Table(driverTableName).Where("id = ?", id).Delete(&drivers)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Unable to delete Please Verify`))
		return
	}
	logger.Infof("Deleted Note with Id:" + id)
	json.NewEncoder(w).Encode(result.RowsAffected)
}

func parseDriverID(w http.ResponseWriter, r *http.Request) uint {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return 0
	}
	return uint(id)
}
