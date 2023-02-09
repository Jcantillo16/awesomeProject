package routers

import (
	"awesomeProject/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetFilterAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	name := r.URL.Query().Get("name")
	fmt.Println("name", name)
	if len(name) < 1 {
		http.Error(w, "Debe enviar el parametro name", http.StatusBadRequest)
		return
	}
	//db.BuscoCancion(name)
	cancion, err := db.BuscoCancion(name)
	if err != nil {
		http.Error(w, "ERROR: La Cancion no existe "+"**********"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cancion)
}
