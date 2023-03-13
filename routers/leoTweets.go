package routers	

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BautistaBianculli/TwitterGolanG/bd"
)
/*LeoTweets Leo los tweets*/
func LeoTweets(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len (ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) // STR A INT
	if err != nil{
		http.Error(w, "Error paginando los tweets", http.StatusBadRequest)
		return
	}

	page := int64 (pagina)
	respuesta, correcto := bd.LeoTweets(ID,page)
	if correcto == false{
		http.Error(w, "Error en la base de datos", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}