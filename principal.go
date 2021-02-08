package main

import (
	"fmt"

	"./models"
)

func main() {
	crearTablas()
	fmt.Println("hola")
}

func crearTablas() {
	models.CrearTablaPersona()
	models.CrearTablaDistrito()
	models.CrearTablaZonaVerde()
	models.CrearTablaArbol()
	models.CrearTablaEstadoArbol()
}
