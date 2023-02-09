package routers

import (
	"net/http"
)

func GetFilterBy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	name := params.Get("name")
	artist := params.Get("artist")
	album := params.Get("album")

	if name != "" && artist != "" && album != "" {
		w.Write([]byte("Se filtra por nombre de la cancion: " + name + " nombre del artista: " +
			artist + " y nombre del album: " + album + "..."))
		GetFilterAll(w, r)

	} else if name != "" {
		w.Write([]byte("Se filtra por nombre de la cancion: " + name + "... "))
		GetFilterName(w, r)

	} else if artist != "" {
		w.Write([]byte("Se filtra por nombre del artista: " + artist + "... "))
		GetFilterArtist(w, r)

	} else if album != "" {
		w.Write([]byte("Se filtra por nombre del album: " + album + "... "))
		GetFilterAlbum(w, r)

	} else {
		http.Error(w, "Debe enviar al menos un parametro de busqueda", http.StatusBadRequest)
		return

	}
}
