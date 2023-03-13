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
	router.HandleFunc("/leoTweet",middleware.ChequeoBD(middleware.ValidoJWT(routers.LeoTweets))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT,handler))

}