package main
import (
	"log"
	"github.com/BautistaBianculli/TwitterGolanG/handlers"
	"github.com/BautistaBianculli/TwitterGolanG/bd"
)


func main() {
	if bd.ConnectionCheck() == false {
		log.Fatal("No Conecction")
		return
	}
	handlers.Manejadores()
}
