package handler

import (
	"github.com/gorilla/mux"
)

//TODO finish here
func RegisterAuthRoutes(r *mux.Router, authService auth.Servie) {
	r.HandleFunc("/login", authService.)
}
