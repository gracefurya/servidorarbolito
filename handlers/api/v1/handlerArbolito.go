package v1

import (
	"fmt"
	"net/http"

	"../../../handlers"
)

//AddArbolito agrega un arbol
func AddArbolito(w http.ResponseWriter, r *http.Request) {
	err := handlers.RecibirArchivo(r, "./recursos/", "ses")
	fmt.Println(err)
}
