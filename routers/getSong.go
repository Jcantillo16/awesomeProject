package routers

import (
	"awesomeProject/db"
	"encoding/json"
	"net/http"
)

func GetFilterName(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	name := r.URL.Query().Get("name")
	if len(name) < 1 {
		http.Error(w, "Debe enviar el parametro name", http.StatusBadRequest)
		return
	}
	cancion, err := db.BuscoCancion(name)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar la cancion "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cancion)
}

func GetFilterArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	artist := r.URL.Query().Get("artist")

	if len(artist) < 1 {
		http.Error(w, "Debe enviar el parametro artist", http.StatusBadRequest)
		return
	}
	cancion, err := db.BuscoCancion(artist)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar la cancion "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cancion)
}

func GetFilterAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	album := r.URL.Query().Get("album")
	if len(album) < 1 {
		http.Error(w, "Debe enviar el parametro album", http.StatusBadRequest)
		return
	}
	cancion, err := db.BuscoCancion(album)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar la cancion "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cancion)
}

func GetFilterBy(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	params := r.URL.Query()
	if len(params) < 1 {
		http.Error(w, "Debe enviar el parametro name y artist o album", http.StatusBadRequest)
		return
	}

	if len(params) > 1 {
		http.Error(w, "Debe enviar solo un parametro", http.StatusBadRequest)
		return
	}

	if len(params) == 1 {
		if params.Get("name") != "" {
			GetFilterName(w, r)
			return
		}
		if params.Get("artist") != "" {
			GetFilterArtist(w, r)
			return
		}
		if params.Get("album") != "" {
			GetFilterAlbum(w, r)
			return
		}
	}

}
