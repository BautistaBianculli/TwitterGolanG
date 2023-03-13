package bd

import(
	"context"
	"time"
	"github.com/BautistaBianculli/TwitterGolanG/models"
)

/*InsertoRelacion graba la relacion en la BD*/
func InsertoRelacion(relacion models.Relacion)(bool, error){
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Twitter")
	col := db.Collection("relacion")

	_, err := col.InsertOne(context,relacion)
	if err != nil{
		return false, err
	}

	return true, nil
}