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
	DirFoto     string    `json:"dirfoto"`
	IDPersona   int       `json:"idpersona"`
	IDZonaVerde string    `json:"idzonaverde"`
}

var queryArbol = `CREATE TABLE if NOT EXISTS arbol(
	idarbol int primary KEY not null AUTO_INCREMENT,
	nombre varchar(60) not null,
	tipo varchar(60) not null,
	fechanac date,
	dirfoto varchar(120),
	persona_id int not null,
	CONSTRAINT fk_arbol_persona
		FOREIGN KEY (persona_id) REFERENCE persona(idpersona),
	zonaverde_id int not null,
	CONSTRAINT fk_arbol_zonaverde
		FOREIGN KEY (zonaverde_id) REFERENCES zonaverde(idzonaverde)
)`
