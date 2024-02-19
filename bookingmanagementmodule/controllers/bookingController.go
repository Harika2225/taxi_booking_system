package controllers

import (
	"net/http"

	auth "com.example.bookingmanagement/auth"
	"com.example.bookingmanagement/handler"
	"github.com/gorilla/mux"
)

type BookingController struct {
}

func (t BookingController) RegisterRoutes(r *mux.Router) {
	r.Handle("/api/createBooking", auth.Protect(http.HandlerFunc(handler.CreateBooking))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/api/getBooking", auth.Protect(http.HandlerFunc(handler.GetBookings))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/getBookingById/{id}", auth.Protect(http.HandlerFunc(handler.GetBookingByID))).Methods(http.MethodGet, http.MethodOptions)
	r.Handle("/api/updateBookingById/{id}", auth.Protect(http.HandlerFunc(handler.UpdateBookingByID))).Methods(http.MethodPut, http.MethodOptions)
	r.Handle("/api/deleteBookingById", auth.Protect(http.HandlerFunc(handler.DeleteBookingByID))).Methods(http.MethodDelete, http.MethodOptions)
}
