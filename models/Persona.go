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
