package bd

import (
	"context"
	"time"

	"github.com/BautistaBianculli/TwitterGolanG/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func InsertoTweet(tweet models.GraboTweet)(string, bool, error){
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	/*DB APUNTA A LA BASE DE DATOS, DB APUNTA A LA TABLA/COLECCION*/
	db:= MongoCN.Database("Twitter")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid" : tweet.UserID,
		"mensaje" : tweet.Mensaje,
		"fecha" : tweet.Fecha,
	}

	result, err := col.InsertOne(contexto, registro)

	if err != nil{
		return "",false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(),true,nil
}