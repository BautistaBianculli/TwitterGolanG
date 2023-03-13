package bd

import (
	"context"
	"time"

	"github.com/BautistaBianculli/TwitterGolanG/models"
	"go.mongodb.org/mongo-driver/bson"
)


func LeoTweetsSeguidores(ID string, page int)([]models.DevulevoTweetsSeguidores, bool){
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	/*DB APUNTA A LA BASE DE DATOS, DB APUNTA A LA TABLA/COLECCION*/
	db:= MongoCN.Database("Twitter")
	col := db.Collection("relacion")

	skip := (page - 1) *20

	conditions:= make([]bson.M,0)
	conditions = append(conditions, bson.M{"$match":bson.M{"usuarioid":ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":	"tweet",
			"localField": "usuariorelacionid",
			"foreignField" : "userid",
			"as" : "tweet",
		}})

	conditions = append(conditions, bson.M{"$unwind":"$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"fecha": -1}})
	conditions = append(conditions, bson.M{"$skip":skip})
	conditions = append(conditions, bson.M{"$limit":20})

	cursor , _ := col.Aggregate(contexto,conditions)
	var results []models.DevulevoTweetsSeguidores
	err := cursor.All(contexto, &results)
	if err != nil{
		return results, false
	}

	return results, true

}