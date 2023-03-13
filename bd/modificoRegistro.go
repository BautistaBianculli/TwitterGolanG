package bd

import(
	"context"
	"time"

	"github.com/BautistaBianculli/TwitterGolanG/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)
/*ModificoREgistro permite modficiar el perfil del usuario*/
func ModificoRegistro (usuario models.User, ID string)(bool , error){
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	/*DB APUNTA A LA BASE DE DATOS, DB APUNTA A LA TABLA/COLECCION*/
	db:= MongoCN.Database("Twitter")
	col := db.Collection("users")

	/*Distintas alternativas para armar el registro de actualizacion a la BD*/
	registro := make(map[string]interface{})
	if len(usuario.Nombre) > 0 {
		registro["nombre"] = usuario.Nombre
	}
	if len(usuario.Apellidos) > 0 {
		registro["apellidos"] = usuario.Apellidos
	}
	if len(usuario.Avatar) > 0 {
		registro["avatar"] = usuario.Avatar
	}
	if len(usuario.Banner) > 0 {
		registro["banner"] = usuario.Banner
	}
	if len(usuario.Biografia) > 0 {
		registro["biografia"] = usuario.Biografia
	}
	if len(usuario.Ubicacion) > 0 {
		registro["ubicacion"] = usuario.Ubicacion
	}
	if len(usuario.SitioWeb) > 0 {
		registro["sitioWeb"] = usuario.SitioWeb
	}
	registro["fechaNacimiento"] = usuario.FechaNacimiento

	updateString := bson.M{
		"$set": registro,
	}

	objID, er := primitive.ObjectIDFromHex(ID)
	if er != nil{
		return false, er
	}

	filtro := bson.M{"_id" : bson.M{"$eq" : objID}}
	_, err := col.UpdateOne(contexto,filtro,updateString)
	if err != nil{
		return false , err
	}
	return true, nil
}