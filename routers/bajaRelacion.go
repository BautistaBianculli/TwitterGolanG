package routers

import(
	"net/http"
	"github.com/BautistaBianculli/TwitterGolanG/bd"
	"github.com/BautistaBianculli/TwitterGolanG/models"
)



func BajaRelacion(w http.ResponseWriter, r* http.Request){
	ID:=r.URL.Query().Get("id")
	if len(ID) < 1{
		http.Error(w, "Enviar el parametro id",http.StatusBadRequest)
		return
	}

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil || status == false{
		http.Error(w, "Error al borrar la relacion " + err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}