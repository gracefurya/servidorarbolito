package models

//ZonaVerde modelo de la zona verde
type ZonaVerde struct {
	IDZonaVerde int    `json:"idzonaverde"`
	Nombre      string `json:"nombre"`
	Direccion   string `json:"direccion"`
	IDDistrito  int    `json:"iddistrito"`
}

//ZonasVerdes array de zonas verdes
type ZonasVerdes []ZonaVerde

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

//AgregarZonaVerde agrega una zona verde
func (z *ZonaVerde) AgregarZonaVerde() error {
	query := `INSERT INTO zonaverde(nombre,direccion,iddistrito) VALUE (?,?,?)`
	_, err := EjecutarExec(query, &z.Nombre, &z.Direccion, &z.IDZonaVerde)
	if err != nil {
		return err
	}
	return nil
}
