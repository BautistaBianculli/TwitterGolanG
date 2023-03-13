package routers


import (
	"encoding/json"
	"net/http"
	"github.com/BautistaBianculli/TwitterGolanG/bd"
	"github.com/BautistaBianculli/TwitterGolanG/models"
)
/*ModificarPerfil modifica el perfil del usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request){

	var usuario models.User

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil{
		http.Error(w, "Datos incorrectos "+err.Error(),400)
		return
	}
	var status bool
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil{
		http.Error(w, "Ocurrió un error al intetar modificar el registro de usuario." + err.Error(),400)
		return
	}
	if !status{
		http.Error(w, "No se modificaron registros. ",400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}