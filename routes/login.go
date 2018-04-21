package routes

import (
	"github.com/AndresED/api/controllers"
	"github.com/gorilla/mux"
)

// SetLoginRouter -> rutas para el registro de usuario
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
