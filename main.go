package main

import (
	"awesomeProject/db"
	"awesomeProject/handlers"
	"log"
)

func main() {

	if db.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}

	handlers.Manejadores()
}
