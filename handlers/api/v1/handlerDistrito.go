package v1

import (
	"net/http"
	"strconv"

	"../../../models"
	"github.com/gorilla/mux"
)

//GetDistritos obtiene todos los distritos
func GetDistritos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		distritos, err := models.ObtenerDistrito()
		if err != nil {
			models.SendUnprocessableEntity(w)
		}
		if len(distritos) == 0 {
			models.SendNoContent(w)

		} else {
			models.SendData(w, distritos)
		}
	}
}

//GetZonasVerdesByDistrito obtiene todas las zonas verdes relacionadas con un distrito
func GetZonasVerdesByDistrito(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		iddistrito, err := strconv.Atoi(vars["iddistrito"])
		if err != nil {
			models.SendUnprocessableEntity(w)
		}

		zonasverdes, err2 := models.GetZonasVerdesByDistrito(iddistrito)
		if err2 != nil {
			models.SendUnprocessableEntity(w)
		}
		models.SendData(w, zonasverdes)
	}
}
