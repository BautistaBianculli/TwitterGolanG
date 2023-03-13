package bd

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



/*BorroTweet Elimino el Tweet del usuario*/
func BorroTweet(ID string, UserID string) error {
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	/*DB APUNTA A LA BASE DE DATOS, DB APUNTA A LA TABLA/COLECCION*/
	db:= MongoCN.Database("Twitter")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id" : objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(contexto,condition)

	return err 
}