package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/AndresED/api/migration"
	"github.com/AndresED/api/routes"
	"github.com/urfave/negroni"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la migración")
		migration.Migrate()
		log.Println("Migración finalizada")
	}
	//inicializa las rutas
	router := routes.InitRoutes()
	//Inicializa los midleware
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	log.Println("Inicializando servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizo la ejecución del programa")
}
