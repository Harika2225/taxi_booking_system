package controllers

import (
	"encoding/json"
	"net/http"

	auth "com.example.drivermanagement/auth"
	"com.example.drivermanagement/handler"
	eureka "com.example.drivermanagement/eurekaregistry"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

type DriverController struct {
}

func (t DriverController) RegisterRoutes(r *mux.Router) {
	r.Handle("/api/createDriver", auth.Protect(http.HandlerFunc(handler.CreateDriver))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/api/getDriver", auth.Protect(http.HandlerFunc(handler.GetDrivers))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/getDriverById/{id}", auth.Protect(http.HandlerFunc(handler.GetDriverByID))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/updateDriverById/{id}", auth.Protect(http.HandlerFunc(handler.UpdateDriverByID))).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/api/deleteDriverById", auth.Protect(http.HandlerFunc(handler.DeleteDriverByID))).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/management/health/readiness", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP", "components": map[string]interface{}{"readinessState": map[string]interface{}{"status": "UP"}}})
	}).Methods(http.MethodGet)
	r.Handle("/rest/services/drivermanagementmodule", auth.Protect(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		logger.Infof("response sent")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"server": "UP"})
	}))).Methods(http.MethodGet)

	r.HandleFunc("/api/services/customermanagementmodule", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "customermanagementmodule") }).Methods(http.MethodGet)
	r.HandleFunc("/api/services/bookingmanagementmodule", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "bookingmanagementmodule") }).Methods(http.MethodGet)
	r.HandleFunc("/api/services/paymentmanagementmodule", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "paymentmanagementmodule") }).Methods(http.MethodGet)
	
	r.Handle("/api/bookingStatus", auth.Protect(http.HandlerFunc(handler.BookingStatus))).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode("helloworld")
	}).Methods(http.MethodGet)
	r.HandleFunc("/management/health/liveness", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP", "components": map[string]interface{}{"livenessState": map[string]interface{}{"status": "UP"}}})
	}).Methods(http.MethodGet)
}
