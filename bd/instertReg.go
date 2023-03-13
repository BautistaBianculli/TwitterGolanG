package bd

import (
	"context"
	"time"

	"github.com/BautistaBianculli/TwitterGolanG/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsterReg es la parada final con la BD para insertar los datos del usuario*/
func InsertReg (usuario models.User) (string, bool, error){

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	/*DB APUNTA A LA BASE DE DATOS, DB APUNTA A LA TABLA/COLECCION*/
	db:= MongoCN.Database("Twitter")
	col := db.Collection("users")

	usuario.Password, _  = EncriptarPassword(usuario.Password)

	result, err := col.InsertOne(contexto,usuario)
	if err != nil{
		return "",false,err
	}

	/*InsertOne devulve un ID y hay que agarralo*/
	ObjId, _:= result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true,nil
}