package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes inicializa las rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	return router
}
