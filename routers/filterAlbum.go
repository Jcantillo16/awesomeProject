package routers

import (
	"awesomeProject/db"
	"encoding/json"
	"net/http"
)

func GetFilterAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	album := r.URL.Query().Get("album")
	if len(album) < 1 {
		http.Error(w, "Debe enviar el parametro album", http.StatusBadRequest)
		return
	}
	db.BuscoCancion(album)
	cancion, err := db.BuscoCanciones(album)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar la cancion "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cancion)
}
