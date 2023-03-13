package models
/*Relacion de usuarios*/
type Relacion struct{
	UsuarioID 			string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID 	string `bson:"usuariorealcionid" json:"usuarioRelacionId"`
}
