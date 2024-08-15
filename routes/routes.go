package routes

import (
	"online-store/controllers/authcontroller"
	ordercontroller "online-store/controllers/orderController"
	"online-store/controllers/productcontroller"
	"online-store/middlewares"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/{id}/password", authcontroller.ChangePassword).Methods("PUT")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)

	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.HandleFunc("/products", productcontroller.Create).Methods("POST")
	api.HandleFunc("/product/{id}", productcontroller.Show).Methods("GET")
	api.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	api.HandleFunc("/product/{id}", productcontroller.Delete).Methods("DELETE")

	api.HandleFunc("/orders", ordercontroller.Index).Methods("GET")
	api.HandleFunc("/orders", ordercontroller.Create).Methods("POST")
	api.HandleFunc("/order/{id}", ordercontroller.Show).Methods("GET")
	api.HandleFunc("/order/{id}/status", ordercontroller.UpdateStatus).Methods("PUT")
	api.HandleFunc("/order/{id}", ordercontroller.Delete).Methods("DELETE")

	return r
}
