package models

//ZonaVerde modelo de la zona verde
type ZonaVerde struct {
	IDZonaVerde int    `json:"idzonaverde"`
	Nombre      string `json:"nombre"`
	Direccion   string `json:"direccion"`
	IDDistrito  int    `json:"iddistrito"`
}

var queryZonaVerde = `CREATE TABLE if NOT EXISTS zonaverde(
	idzonaverde int primary KEY not null AUTO_INCREMENT,
	nombre varchar(45),
	direccion varchar(45),
	distrito_id int,
	CONSTRAINT fk_zonaverde_distrito
		FOREIGN KEY (distrito_id) REFERENCES distrito(iddistrito)
		ON DELETE CASCADE
)`

//CrearTablaZonaVerde crea la tabla zona verde
func CrearTablaZonaVerde() {
	EjecutarExec(queryZonaVerde)
}
