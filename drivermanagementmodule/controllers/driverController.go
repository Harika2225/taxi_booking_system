package controllers

import (
	"net/http"
	auth "com.example.drivermanagement/auth"
	"com.example.drivermanagement/handler"
	"github.com/gorilla/mux"
)

type DriverController struct {
}

func (t DriverController) RegisterRoutes(r *mux.Router) {
	r.Handle("/api/createDriver", auth.Protect(http.HandlerFunc(handler.CreateDriver))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/api/getDriver", auth.Protect(http.HandlerFunc(handler.GetDrivers))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/getDriverById/{id}", auth.Protect(http.HandlerFunc(handler.GetDriverByID))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/updateDriverById/{id}", auth.Protect(http.HandlerFunc(handler.UpdateDriverByID))).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/api/deleteDriverById", auth.Protect(http.HandlerFunc(handler.DeleteDriverByID))).Methods(http.MethodDelete, http.MethodOptions)
}
