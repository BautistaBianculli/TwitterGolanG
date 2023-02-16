package routers

import(
	"encoding/json"
	"net/http"
	"time"
	
	"github.com/BautistaBianculli/TwitterGolanG/bd"
	"github.com/BautistaBianculli/TwitterGolanG/models"
	"github.com/BautistaBianculli/TwitterGolanG/jwt"
	
)
/*Login recibe Login*/
func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","aplication/json")


	var usuario models.User

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil{
		http.Error(w, "Usuario y/o contraseña invalido\n" + err.Error(),400)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "El mail del usuario es requerido\n",400)
		return
	}

	document, exist := bd.IntentoLogin(usuario.Email,usuario.Password)
	if !exist {
		http.Error(w, "Usuario y/o contraseña invalido\n",400)
		return
	} 

	/*GENERO TOKEN*/

	jwtKey, err := jwt.GeneroJWT(document)
	if err != nil {
		http.Error(w, "Error al generar el token\n" + err.Error(),400)
		return
	}

	respuesta := models.RespuestaLogin{
		Token : jwtKey,
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

	/*GRABO LA COOKIE*/

	expirationTokenTime := time.Now().Add(24*time.Hour)
	http.SetCookie(w,&http.Cookie{
		Name: 	"token",
		Value:	jwtKey,
		Expires: expirationTokenTime,
	})
}