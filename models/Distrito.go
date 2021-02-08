package models

//Distrito modelo distrito
type Distrito struct {
	IDDistrito int    `json:"iddistrito"`
	Nombre     string `json:"nombre"`
	Codigo     int    `json:"codigo"`
}

var queryDistrito = `CREATE TABLE if NOT EXISTS distrito(
	iddistrito int primary KEY not null AUTO_INCREMENT,
	nombre varchar(45),
	codigo int
)`

//CrearTablaDistrito Crea la tabla de distrito en la base de datos
func CrearTablaDistrito() {
	EjecutarExec(queryDistrito)
}
