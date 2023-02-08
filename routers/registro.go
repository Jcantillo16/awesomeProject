package routers

//func Registro(w http.ResponseWriter, r *http.Request) {
//
//	var t models.Song
//	err := json.NewDecoder(r.Body).Decode(&t)
//	if err != nil {
//		http.Error(w, "hay datos requeridos"+err.Error(), 400)
//		return
//	}
//	if len(t.Name) == 0 {
//		http.Error(w, "El Nombre de la cancion es requerido", 400)
//		return
//	}
//	_, status, err := db.InsertoRegistro(
//	if err != nil {
//		http.Error(w, "Ocurrió un error al intentar realizar el registro de la cancion "+err.Error(), 400)
//		return
//	}
//	if status == false {
//		http.Error(w, "No se ha logrado insertar el registro de a cancion ", 400)
//		return
//	}
//	w.WriteHeader(http.StatusCreated)
//	fmt.Println("Registro insertado correctamente")
//	json.NewEncoder(w).Encode(t)
//
//}
