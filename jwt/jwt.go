package jwt

import(
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/BautistaBianculli/TwitterGolanG/models"
)

/*GeneroJWT genera el ecnriptado del token con JWT*/
func GeneroJWT(usuario models.User)(string, error){
	miClave := []byte("Clave12345")
	/*Payload es DATA del JWT*/
	payload := jwt.MapClaims{

		"email" 			: 	usuario.Email,
		"nombre"			: 	usuario.Nombre,
		"apellidos" 		: 	usuario.Apellidos,
		"fecha_nacimiento" 	: 	usuario.FechaNacimiento,
		"biografia"			:	usuario.Biografia,
		"ubicacion"			:	usuario.Ubicacion,
		"sitioweb"			:	usuario.SitioWeb,
		"_id"				:	usuario.ID.Hex(),
		"exp"				: 	time.Now().Add(time.Hour * 24 ).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
		
	tokenStr, err := token.SignedString(miClave)
	if err !=nil {
		return tokenStr, err
	}

	return tokenStr,nil
}