package controllers

import (
	"encoding/json"
	"net/http"

	auth "com.example.bookingmanagement/auth"
	"com.example.bookingmanagement/handler"
	eureka "com.example.paymentmanagement/eurekaregistry"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

type CommunicationController struct {
}

func (t CommunicationController) RegisterRoutes(r *mux.Router) {
	r.Handle("/rest/services/bookingmanagementmodule", auth.Protect(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		logger.Infof("response sent")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"server": "UP"})
	}))).Methods(http.MethodGet)

	r.HandleFunc("/api/services/customermanagementmodule", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "customermanagementmodule") }).Methods(http.MethodGet)
	r.HandleFunc("/api/services/drivermanagementmodule", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "paymentmanagementmodule") }).Methods(http.MethodGet)
	r.HandleFunc("/api/services/paymentmanagementmodule", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "paymentmanagementmodule") }).Methods(http.MethodGet)

	r.Handle("/api/bookingAccepted", auth.Protect(http.HandlerFunc(handler.BookingAccepted))).Methods(http.MethodPost, http.MethodOptions)
}
