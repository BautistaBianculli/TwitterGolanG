package bd

import (
	"context"
	"log"
	"time"

	"github.com/BautistaBianculli/TwitterGolanG/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil busca un perfil en la BD*/
func BuscoPerfil(ID string) (models.User, error) {
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Twitter")
	col := db.Collection("users")

	var perfil models.User
	/*El ID es string y tengo guardado un ObjetID, tengo que convertir*/
	objID, _ := primitive.ObjectIDFromHex(ID)
	/*Condicion del search*/
	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(context, condition).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		log.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}

	return perfil, nil
}
