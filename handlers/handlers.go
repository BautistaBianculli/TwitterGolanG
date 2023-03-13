package handlers

import (
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/BautistaBianculli/TwitterGolanG/routers"
	"github.com/BautistaBianculli/TwitterGolanG/middleware"

)
/*Manejadores setteo mi puerto y pongo a escuchar al SRV*/
func Manejadores(){
	router := mux.NewRouter()

	router.HandleFunc("/registro",middleware.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login",middleware.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil",middleware.ChequeoBD(middleware.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil",middleware.ChequeoBD(middleware.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet",middleware.ChequeoBD(middleware.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweet",middleware.ChequeoBD(middleware.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet",middleware.ChequeoBD(middleware.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar",middleware.ChequeoBD(middleware.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subirBanner",middleware.ChequeoBD(middleware.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerAvatar",middleware.ChequeoBD(middleware.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/obtenerBanner",middleware.ChequeoBD(middleware.ValidoJWT(routers.ObtenerBanner))).Methods("GET")


	router.HandleFunc("/altaRelacion",middleware.ChequeoBD(middleware.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion",middleware.ChequeoBD(middleware.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion",middleware.ChequeoBD(middleware.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")


	router.HandleFunc("/listaUsuarios",middleware.ChequeoBD(middleware.ValidoJWT(routers.ListaUsuarios))).Methods("GET")



	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT,handler))

}