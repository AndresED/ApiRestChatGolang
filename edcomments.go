package main

import (
	"flag"
	"log"

	"github.com/AndresED/api/migration"
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
}
