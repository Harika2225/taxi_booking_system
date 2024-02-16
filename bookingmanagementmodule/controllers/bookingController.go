package controllers

import (
	"encoding/json"
	"net/http"

	auth "com.example.bookingmanagement/auth"
	eureka "com.example.bookingmanagement/eurekaregistry"
	"com.example.bookingmanagement/handler"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

type BookingController struct {
}

func (t BookingController) RegisterRoutes(r *mux.Router) {
	r.Handle("/api/createBooking", auth.Protect(http.HandlerFunc(handler.CreateBooking))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/api/getBooking", auth.Protect(http.HandlerFunc(handler.GetBookings))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/getBookingById/{id}", auth.Protect(http.HandlerFunc(handler.GetBookingByID))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/updateBookingById/{id}", auth.Protect(http.HandlerFunc(handler.UpdateBookingByID))).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/api/deleteBookingById", auth.Protect(http.HandlerFunc(handler.DeleteBookingByID))).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/management/health/readiness", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP", "components": map[string]interface{}{"readinessState": map[string]interface{}{"status": "UP"}}})
	}).Methods(http.MethodGet)
	r.Handle("/rest/services/bookingmanagementmodule", auth.Protect(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		logger.Infof("response sent")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"server": "UP"})
	}))).Methods(http.MethodGet)

	r.HandleFunc("/api/services/customermanagementmodule", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "customermanagementmodule") }).Methods(http.MethodGet)
	r.HandleFunc("/api/services/drivermanagementmodule", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "paymentmanagementmodule") }).Methods(http.MethodGet)
	r.HandleFunc("/api/services/paymentmanagementmodule", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "paymentmanagementmodule") }).Methods(http.MethodGet)

	r.Handle("/api/boookingStatus",auth.Protect(http.HandlerFunc(handler.BookingAccepted))).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode("helloworld")
	}).Methods(http.MethodGet)
	r.HandleFunc("/management/health/liveness", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP", "components": map[string]interface{}{"livenessState": map[string]interface{}{"status": "UP"}}})
	}).Methods(http.MethodGet)
}
