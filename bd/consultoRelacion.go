package bd 

import (
	"context"
	"time"

	"github.com/BautistaBianculli/TwitterGolanG/models"
	"go.mongodb.org/mongo-driver/bson"
)
/*ConsultoRelacion consulta la relacion entre dos usuarios*/
func ConsultoRelacion (t models.Relacion)(bool,error){
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	/*DB APUNTA A LA BASE DE DATOS, DB APUNTA A LA TABLA/COLECCION*/
	db:= MongoCN.Database("Twitter")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid" : t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := col.FindOne(contexto,condicion).Decode(&resultado)
	if err != nil{
		return false, err
	}
	return true , nil
}