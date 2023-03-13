package routers

import(
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/BautistaBianculli/TwitterGolanG/bd"
	"github.com/BautistaBianculli/TwitterGolanG/models"
)


/*SubirAvatar subo el avatar a la base de datos y lo copia en la carpeta*/
func SubirAvatar(w http.ResponseWriter, r* http.Request){
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename,".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE,0666)
	if err != nil{
		http.Error(w,"Error al subir imagen!" + err.Error(),http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f,file)
	if err != nil{
		http.Error(w,"Error al copiar la imagen!" + err.Error(),http.StatusBadRequest)
		return
	}

	var usuario models.User
	var status bool

	usuario.Avatar = IDUsuario + "." + extension

	status, err = bd.ModificoRegistro(usuario,IDUsuario)
	if err != nil || status == false{
		http.Error(w,"Error al grabar el avatar en la base de datos " + err.Error(),http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
		
}