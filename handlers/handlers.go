package handlers

import (
	"awesomeProject/middlewares"
	"awesomeProject/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// Manejadores seteo mi puerto, el Handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	//router.HandleFunc("/registro", middlewares.ChequeoBD(routers.Registro)).Methods("POST")

	//ruta para buscar una cancion
	router.HandleFunc("/getSong/", middlewares.ChequeoBD(routers.GetFilterBy)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
