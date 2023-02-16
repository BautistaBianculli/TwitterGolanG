package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*mongoCN Objeto de conection a BD */
var MongoCN = ConnectionBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:4E7A2100AF@twitterudemy.cu2qmno.mongodb.net/?retryWrites=true&w=majority")

/*ConnectionBD funcion conectar BD */
func ConnectionBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa")
	return client
}

/*ConnectionCheck Chequea si la conexion esta abierta mediante un PING a BD*/
func ConnectionCheck() bool {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
