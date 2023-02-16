package bd

import (
	"github.com/BautistaBianculli/TwitterGolanG/models"
	"golang.org/x/crypto/bcrypt"

)

/*IntentoLogin realiza el chequeo de Login BD*/
func IntentoLogin(email string, password string)(models.User, bool){
	usuario, found,_ := UserCheck(email)
	if !found{
		return usuario, false

	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD,passwordBytes)
	if err != nil{
		return usuario, false
	}

	return usuario,true

}