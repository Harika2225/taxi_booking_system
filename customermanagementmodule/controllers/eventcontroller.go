package controllers

import (
	"encoding/json"
	"net/http"

	auth "com.example.customermanagement/auth"
	"com.example.customermanagement/handler"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

type EventController struct {
}

func (t EventController) RegisterRoutes(r *mux.Router) {
	r.Handle("/event", auth.Protect(http.HandlerFunc(handler.AddEvent))).Methods(http.MethodPost)
	r.Handle("/events", auth.Protect(http.HandlerFunc(handler.GetEvents))).Methods(http.MethodGet)
	r.Handle("/events/{id}", auth.Protect(http.HandlerFunc(handler.ReadEventById))).Methods(http.MethodGet)
	r.Handle("/update", auth.Protect(http.HandlerFunc(handler.UpdateEvent))).Methods(http.MethodPatch)
	r.Handle("/delete/{id}", auth.Protect(http.HandlerFunc(handler.DeleteEvent))).Methods(http.MethodDelete)
	r.Handle("/api/notes", auth.Protect(http.HandlerFunc(handler.AddNote))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/api/notes", auth.Protect(http.HandlerFunc(handler.GetNotes))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/notes/{id}", auth.Protect(http.HandlerFunc(handler.ReadNoteById))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/update/notes", auth.Protect(http.HandlerFunc(handler.UpdateNote))).Methods(http.MethodPatch, http.MethodOptions)
	r.Handle("/api/notes", auth.Protect(http.HandlerFunc(handler.DeleteNote))).Methods(http.MethodDelete, http.MethodOptions)

	r.Handle("/api/createCustomer", auth.Protect(http.HandlerFunc(handler.CreateCustomerHandler))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/api/getCustomer", auth.Protect(http.HandlerFunc(handler.GetCustomersHandler))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/getCustomerById/{id}", auth.Protect(http.HandlerFunc(handler.GetCustomerByIDHandler))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/updateCustomerById/{id}", auth.Protect(http.HandlerFunc(handler.UpdateCustomerByIDHandler))).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/api/deleteCustomerById", auth.Protect(http.HandlerFunc(handler.DeleteCustomerByIDHandler))).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/management/health/readiness", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP", "components": map[string]interface{}{"readinessState": map[string]interface{}{"status": "UP"}}})
	}).Methods(http.MethodGet)
	r.Handle("/rest/services/customermanagementmodule", auth.Protect(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		logger.Infof("response sent")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"server": "UP"})
	}))).Methods(http.MethodGet)

	r.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode("helloworld")
	}).Methods(http.MethodGet)
	r.HandleFunc("/management/health/liveness", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP", "components": map[string]interface{}{"livenessState": map[string]interface{}{"status": "UP"}}})
	}).Methods(http.MethodGet)
}
