package routers

import (
	"encoding/json"
	"net/http"
	"github.com/BautistaBianculli/TwitterGolanG/bd"
	"github.com/BautistaBianculli/TwitterGolanG/models"
)

/**/
func Registro(w http.ResponseWriter, r *http.Request){
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil{
		http.Error(w, "Error en los datos recibidos "+err.Error(),400)
		return
	}

	if len(t.Email) == 0{
		http.Error(w, "Email vacio\n",400)
		return
	}
	if len(t.Password) <= 6{
		http.Error(w, "La password es muy pobre\n",400)
		return
	}

	_,encontrado,_ :=bd.UserCheck(t.Email)
	if encontrado == true{
		http.Error(w, "Usuario ya registrado\n", 400)
		return
	}

	_,status, err := bd.InsertReg(t)
	if err != nil{
		http.Error(w, "Ocurrio un error al registrar el usuario \n"+err.Error(),400)
		return
	}

	if status == false{
		http.Error(w, "No se logro isnertar el registro del usuario\n",400)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
}