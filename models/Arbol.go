package models

import (
	"time"
)

//Arbol modelo de arbol
type Arbol struct {
	IDArbol     int       `json:"idarbol"`
	Nombre      string    `json:"nombre"`
	Tipo        string    `json:"tipo"`
	FechaNac    time.Time `json:"fechanac"`
	Celular     string    `json:"celular"`
	IDPersona   int       `json:"idpersona"`
	IDZonaVerde string    `json:"idzonaverde"`
}
