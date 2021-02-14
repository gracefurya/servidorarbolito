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

//Arboles array de arbol
type Arboles []Arbol

var queryArbol = `CREATE TABLE if NOT EXISTS arbol(
	idarbol int primary KEY not null AUTO_INCREMENT,
	nombre varchar(60) not null,
	tipo varchar(60) not null,
	fechanac DATETIME(0),
	dirfoto varchar(120),
	persona_id int not null,
	CONSTRAINT fk_arbol_persona
		FOREIGN KEY (persona_id) REFERENCES persona(idpersona),
	zonaverde_id int not null,
	CONSTRAINT fk_arbol_zonaverde
		FOREIGN KEY (zonaverde_id) REFERENCES zonaverde(idzonaverde)
)`

//CrearTablaArbol se crea una tabla para el objeto arbol
func CrearTablaArbol() {
	EjecutarExec(queryArbol)
}

//AgregarArbol agregamos un arbol a la base de datos
func (a *Arbol) AgregarArbol() error {
	query := `INSERT INTO arbol(nombre,tipo,fechanac,dirfoto,persona_id,zonaverde_id) VALUE(?,?,?,?,?,?)`
	var fn = FechaAString(a.FechaNac)
	if _, err := EjecutarExec(query, &a.Nombre, &a.Tipo, fn, &a.DirFoto, &a.IDPersona, &a.IDZonaVerde); err != nil {
		return err
	}
	return nil
}

//ObtenerArbolByID crea un arbol relacionada a una persona y a su zona verde
func ObtenerArbolByID(id int) (*Arbol, error) {
	query := `SELECT idarbol,nombre,tipo,fechanac,dirfoto,persona_id,zonaverde_id FROM arbol
	WHERE idarbol=?`
	rows, err := EjecutarQuery(query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	arbol := &Arbol{}
	if rows.Next() {
		rows.Scan(&arbol.IDArbol, &arbol.Nombre, &arbol.Tipo, &arbol.FechaNac, &arbol.DirFoto, &arbol.IDPersona, &arbol.IDZonaVerde)
	}

	return arbol, nil
}

//ObtenerArbolesByPersonaID obtiene todos los arboles relacionadas a una persona
func ObtenerArbolesByPersonaID(idpersona int) (Arboles, error) {
	query := `SELECT idarbol,nombre,tipo,fechanac,dirfoto,persona_id,zonaverde_id FROM arbol
	WHERE persona_id=?`
	rows, err := EjecutarQuery(query, idpersona)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	arboles := Arboles{}
	for rows.Next() {
		arbol := &Arbol{}
		rows.Scan(&arbol.IDArbol, &arbol.Nombre, &arbol.Tipo, &arbol.FechaNac, &arbol.DirFoto, &arbol.IDPersona, &arbol.IDZonaVerde)
		arboles = append(arboles, *arbol)
	}

	return arboles, nil
}
