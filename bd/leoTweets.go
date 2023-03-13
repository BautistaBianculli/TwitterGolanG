package bd

import (
	"context"
	"log"
	"time"

	"github.com/BautistaBianculli/TwitterGolanG/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweets(ID string, pagina int64)([]*models.DevulevoTweets, bool){
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	/*DB APUNTA A LA BASE DE DATOS, DB APUNTA A LA TABLA/COLECCION*/
	db:= MongoCN.Database("Twitter")
	col := db.Collection("tweet")

	var result []*models.DevulevoTweets

	condition := bson.M{
		"userid": ID,
	}
	
	options := options.Find()
	options.SetLimit(20)//DOCUMENTOS LIMITES PARA PAGINAR
	options.SetSort(bson.D{{Key:"fecha",Value: -1}}) // ORDENADAMOS LO ANTERIOR por orden desc con el -1 por el campo fecha
	options.SetSkip((pagina -1)*20)

	cursor, err := col.Find(contexto,condition,options)
	if err != nil{
		log.Fatal(err.Error())
		return result, false
	}
	for cursor.Next(context.TODO()){
		var registro models.DevulevoTweets
		err := cursor.Decode(&registro)
		if err != nil{
			return result, false
		}
		result = append(result,&registro)
	}
	return result,true
}