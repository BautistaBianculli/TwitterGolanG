package models

/*RespuestaLogin tiene el token que se devuelve on el logeo*/
type RespuestaLogin struct{
	Token string `json:"token,omitempty"`
}