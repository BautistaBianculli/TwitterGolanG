package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BautistaBianculli/TwitterGolanG/bd"
	"github.com/BautistaBianculli/TwitterGolanG/models"
)

/*GraboTweet perimte grabaar el tweet en la bd*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al Insertar el Tweet. Intente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
