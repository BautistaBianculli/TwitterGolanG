package routers

import (
	"encoding/json"
	"net/http"

	"github.com/BautistaBianculli/TwitterGolanG/bd"
)

/*VerPerfil permite extraer los valores de Perfil*/
func VerPerfil(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID) == 0 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	perfil , err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "No se encontro el registro " + err.Error(), 400)
		return
	}

	w.Header().Set("context-type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}