package bd

import (
	"context"
	"time"

	"github.com/BautistaBianculli/TwitterGolanG/models"
	"go.mongodb.org/mongo-driver/bson"
)
/*UserCheck recibe un email de parametro y chequea si existe*/
func UserCheck(email string )(models.User,bool,string){
	contexto, cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	db:= MongoCN.Database("Twitter")
	col := db.Collection("users")

	condicion := bson.M{"email" : email}

	var result models.User

	err := col.FindOne(contexto,condicion).Decode(&result)
	ID := result.ID.Hex()
	if err != nil{
		return result, false, ID
	}
	return result,true,ID
}