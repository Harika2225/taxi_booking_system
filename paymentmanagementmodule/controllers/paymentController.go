package controllers

import (
	"net/http"

	auth "com.example.paymentmanagement/auth"
	"com.example.paymentmanagement/handler"
	"github.com/gorilla/mux"
)

type PaymentController struct {
}

func (t PaymentController) RegisterRoutes(r *mux.Router) {
	r.Handle("/api/createPayment", auth.Protect(http.HandlerFunc(handler.CreatePayment))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/api/getPayment", auth.Protect(http.HandlerFunc(handler.GetPayments))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/getPaymentById/{id}", auth.Protect(http.HandlerFunc(handler.GetPaymentByID))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/updatePaymentById/{id}", auth.Protect(http.HandlerFunc(handler.UpdatePaymentByID))).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/api/deletePaymentById", auth.Protect(http.HandlerFunc(handler.DeletePaymentByID))).Methods(http.MethodDelete, http.MethodOptions)
}
