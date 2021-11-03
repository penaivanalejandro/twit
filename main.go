package main

import (
	"log"

	"github.com/penaivanalejandro/twit/bd"
	"github.com/penaivanalejandro/twit/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
