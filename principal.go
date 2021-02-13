package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "./handlers/api/v1"
	"./models"
	"github.com/gorilla/mux"
)

func main() {
	crearTablas()
	fmt.Println("hola")
	mux := mux.NewRouter()
	mux.HandleFunc("/api/v1/persona/{idpersona:[0-9]+}", v1.GetPersonaByID).Methods("GET")
	mux.HandleFunc("/api/v1/persona", v1.AddPersona).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func crearTablas() {
	models.CrearTablaPersona()
	models.CrearTablaDistrito()
	models.CrearTablaZonaVerde()
	models.CrearTablaArbol()
	models.CrearTablaEstadoArbol()

}
