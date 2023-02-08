package routers

import (
	"net/http"
)

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
