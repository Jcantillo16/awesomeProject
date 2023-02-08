package routers

import (
	"awesomeProject/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetFilterArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	artist := r.URL.Query().Get("artistName")
	fmt.Println("artist:", artist)

	if len(artist) < 1 {
		http.Error(w, "Debe enviar el parametro artist", http.StatusBadRequest)
		return
	}
	cancion, err := db.BuscoCancion(artist)
	if err != nil {
		http.Error(w, "Ocurrió un error "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cancion)
}
