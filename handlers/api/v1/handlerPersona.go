package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../../../models"
	"github.com/gorilla/mux"
)

//GetPersonaByID devuelve un dato del tipo persona
func GetPersonaByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		re, err := strconv.Atoi(vars["idpersona"])
		if err != nil {
			log.Println("Error en GetPersonaById: ", err)
			models.SendUnprocessableEntity(w)
			return
		}
		persona, err2 := models.ObtenerUsuario(re)
		if err2 != nil {
			log.Println("Error en GetPersonaById: ", err)
			models.SendUnprocessableEntity(w)
			return
		}

		if persona.IDPersona == 0 {
			models.SendNotFound(w)
			return
		}

		models.SendData(w, persona)

	}
}

//GetPersonaByCi devuleve una persona si encuentra su ci
func GetPersonaByCi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		re := vars["cipersona"]

		persona, err := models.ObtenerUsuarioByCI(re)
		if err != nil {
			models.SendUnprocessableEntity(w)
			return
		}
		if persona.IDPersona == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, persona)
	}
}

//AddPersona agregar persona
func AddPersona(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		persona := models.Persona{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&persona); err != nil {
			models.SendUnprocessableEntity(w)
		}

		err := persona.AgregarUsuario()

		if err != nil {
			models.SendUnprocessableEntity(w)
		}

	}
}
