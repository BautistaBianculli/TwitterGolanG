package routers

import (
	"errors"
	"strings"

	"github.com/BautistaBianculli/TwitterGolanG/bd"
	"github.com/BautistaBianculli/TwitterGolanG/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email valor usado en todos los EndPoints */
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usa en todos los EndPoints */
var IDUsuario string

/*ProcesoToken proceso token apra extrar los valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Clave12345")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.UserCheck(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
