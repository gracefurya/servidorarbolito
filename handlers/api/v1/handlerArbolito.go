package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../../../handlers"
	"../../../models"
)

//AddArbolito agrega un arbol
func AddArbolito(w http.ResponseWriter, r *http.Request) {
	arbol := models.Arbol{}

	if err := json.Unmarshal([]byte(r.FormValue("arbol")), &arbol); err != nil {
		fmt.Println(err, "  ", arbol)
		models.SendUnprocessableEntity(w)

	} else {
		if idarbol, err2 := arbol.AgregarArbol(); err2 != nil {
			models.SendUnprocessableEntity(w)
		} else {
			err3 := handlers.RecibirArchivo(r, "./recursos/", strconv.Itoa(int(idarbol)))
			if err3 != nil {
				fmt.Println(err3)
			}
		}

	}

}
