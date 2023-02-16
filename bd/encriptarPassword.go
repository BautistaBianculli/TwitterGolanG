package bd

import  "golang.org/x/crypto/bcrypt"

/*ENCRIPTO LA PASSWORD 256 VECES 2^8*/
func EncriptarPassword(password string)(string, error){
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),costo)
	return string(bytes),err
}