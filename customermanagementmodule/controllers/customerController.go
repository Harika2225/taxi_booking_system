package controllers

import (
	"net/http"

	auth "com.example.customermanagement/auth"
	"com.example.customermanagement/handler"
	"github.com/gorilla/mux"
)

type CustomerController struct {
}

func (t CustomerController) RegisterRoutes(r *mux.Router) {
	r.Handle("/api/createCustomer", auth.Protect(http.HandlerFunc(handler.CreateCustomer))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/api/getCustomer", auth.Protect(http.HandlerFunc(handler.GetCustomers))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/getCustomerById/{id}", auth.Protect(http.HandlerFunc(handler.GetCustomerByID))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/updateCustomerById/{id}", auth.Protect(http.HandlerFunc(handler.UpdateCustomerByID))).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/api/deleteCustomerById", auth.Protect(http.HandlerFunc(handler.DeleteCustomerByID))).Methods(http.MethodDelete, http.MethodOptions)

}
