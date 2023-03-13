package bd

import (
	"context"
	"time"

	"github.com/BautistaBianculli/TwitterGolanG/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodos(ID string, page int64, search string, tipo string)([]*models.User, bool){
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	/*DB APUNTA A LA BASE DE DATOS, DB APUNTA A LA TABLA/COLECCION*/
	db:= MongoCN.Database("Twitter")
	col := db.Collection("users")

	var results []* models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1)*20)
	findOptions.SetLimit(20)
	
	query := bson.M{
		"nombre":bson.M{"$regex": `(?i)`+ search},
	}

	cursor, err := col.Find(contexto,query,findOptions)
	if err != nil{
		return results, false
	}

	var encontrado, incluir bool

	for cursor.Next(contexto){
		var s models.User
		err := cursor.Decode(&s)
		if err != nil{
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false 

		encontrado, _ = ConsultoRelacion(r)
		if tipo == "new" && !encontrado{
			incluir = true
		}

		if tipo == "follow" && encontrado{
			incluir = true
		}
		if r.UsuarioRelacionID == ID{
			incluir = false
		}

		if incluir{
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cursor.Err()
	if err != nil{
		return results,false
	}
	cursor.Close(contexto)
	return results, true
}