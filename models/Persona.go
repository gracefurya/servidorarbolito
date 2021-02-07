package models

//Persona modelo de persona
type Persona struct {
	IDPersona       int    `json:"idpersona"`
	Nombre          string `json:"nombre"`
	Paterno         string `json:"paterno"`
	Materno         string `json:"materno"`
	CarnetIdentidad string `json:"carenetidentidad"`
	Telefono        string `json:"telefono"`
	Celular         string `json:"celular"`
}

var queryPersona = `CREATE TABLE if NOT EXISTS persona(
	idpersona int primary KEY not null AUTO_INCREMENT,
	nombre varchar(45),
	paterno varchar(45),
	materno varchar(45),
	carnetidentidad varchar(12),
	telefono varchar(12),
	celular varchar(12)
)`

func CrearTablaPersona() {
	EjecutarExec(queryPersona)
}
