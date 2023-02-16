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

	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT,handler))

}