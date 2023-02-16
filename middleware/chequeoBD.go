package middleware

import (
	"net/http"
	"github.com/BautistaBianculli/TwitterGolanG/bd"
)
/* ChequeoBD middleware permite concoer el estado de la abse de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if bd.ConnectionCheck() == false {
			http.Error(w, "Conexion Perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w,r)
	}
}