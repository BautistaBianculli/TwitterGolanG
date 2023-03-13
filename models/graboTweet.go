package models

import "time"

/*GaboTweet es el formato de estructura que tendra el tweet*/
type GraboTweet struct {
	UserID 		string 		`bson:"userid" json:"userid,omitempty"` 
	Mensaje 	string 		`bson:"mensaje" json:"mensaje,omitempty"` 
	Fecha 		time.Time 	`bson:"fecha" json:"fecha,omitempty"` 
}