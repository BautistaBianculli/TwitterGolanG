package bd

import(
	"context"
	"time"
	"github.com/BautistaBianculli/TwitterGolanG/models"
)



func BorroRelacion(t models.Relacion)(bool,error){
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Twitter")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(context,t)
	if err != nil{
		return false ,err
	}
	return true, nil
}	