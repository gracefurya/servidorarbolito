package v1

import (
	"net/http"

	"../../../handlers"
)

//AddArbolito agrega un arbol
func AddArbolito(w http.ResponseWriter, r *http.Request) {
	handlers.RecibirArchivo(r, "./recursos/", "ses")
}
