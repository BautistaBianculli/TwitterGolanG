package routers

import (
	"net/http"
	"encoding/json"
	

	"github.com/BautistaBianculli/TwitterGolanG/bd"
	"github.com/BautistaBianculli/TwitterGolanG/models"
)

func ConsultaRelacion(w http.ResponseWriter, r * http.Request){
	ID := r.URL.Query().Get("id")

	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID
	var respuesta models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(relacion)
	if err != nil || !status {
		respuesta.Status = false 
	}else{
		respuesta.Status = true
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}