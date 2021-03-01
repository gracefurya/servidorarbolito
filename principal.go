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
	mux.HandleFunc("/api/v1/persona/ci/{cipersona}", v1.GetPersonaByCi).Methods("GET")
	mux.HandleFunc("/api/v1/persona", v1.AddPersona).Methods("POST")
	mux.HandleFunc("/api/v1/distrito", v1.GetDistritos).Methods("GET")
	mux.HandleFunc("/api/v1/zonaverde/distrito/{iddistrito:[0-9]+}", v1.GetZonasVerdesByDistrito).Methods("GET")
	mux.HandleFunc("/api/v1/arbolito", v1.AddArbolito).Methods("POST")
	fs := http.FileServer(http.Dir("recursos"))
	mux.PathPrefix("/recursos/").Handler(http.StripPrefix("/recursos/", fs))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func crearTablas() {
	models.CrearTablaPersona()
	models.CrearTablaDistrito()
	models.CrearTablaZonaVerde()
	models.CrearTablaArbol()
	models.CrearTablaEstadoArbol()

}
