package models

//ZonaVerde Modelo de la zona verde
type ZonaVerde struct {
	IDZonaVerde int    `json:"idzonaverde"`
	Nombre      string `json:"nombre"`
	Direccion   string `json:"direccion"`
	IDDistrito  int    `json:"iddistrito"`
}

//ZonasVerdes Array de zonas verdes
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

//CrearTablaZonaVerde Crea la tabla zona verde
func CrearTablaZonaVerde() {
	EjecutarExec(queryZonaVerde)
}

//AgregarZonaVerde Agrega una zona verde en la base de datos
func (z *ZonaVerde) AgregarZonaVerde() error {
	query := `INSERT INTO zonaverde(nombre,direccion,distrito_id) VALUES (?,?,?)`
	_, err := EjecutarExec(query, &z.Nombre, &z.Direccion, &z.IDDistrito)
	if err != nil {
		return err
	}
	return nil
}

//GetZonaVerdeByID Obtiene una zona verde mediante su ID
func GetZonaVerdeByID(id int) (*ZonaVerde, error) {
	query := `SELECT idzonaverde,nombre,direccion,distrito_id FROM zonaverde WHERE idzonaverde=?`
	rows, err := EjecutarQuery(query, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	zonaverde := &ZonaVerde{}
	if rows.Next() {
		rows.Scan(&zonaverde.IDZonaVerde, &zonaverde.Nombre, &zonaverde.Direccion, &zonaverde.IDDistrito)
	}

	return zonaverde, nil
}

//GetZonasVerdes Obtiene todas las zonas verdes
func GetZonasVerdes() (ZonasVerdes, error) {
	query := `SELECT idzonaverde,nombre,direccion,distrito_id FROM zonaverde`
	rows, err := EjecutarQuery(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	zonasverdes := ZonasVerdes{}
	for rows.Next() {
		zv := ZonaVerde{}
		rows.Scan(&zv.IDZonaVerde, &zv.Nombre, &zv.Direccion, &zv.IDDistrito)
		zonasverdes = append(zonasverdes, zv)
	}
	return zonasverdes, nil
}

//GetZonasVerdesByDistrito Obtiene todas las zonas verdes de un ditrito
func GetZonasVerdesByDistrito(iddistrito int) (ZonasVerdes, error) {
	query := `SELECT idzonaverde,nombre,distrito_id,direccion FROM zonaverde WHERE distrito_id=?`
	rows, err := EjecutarQuery(query, iddistrito)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	zonasverdes := ZonasVerdes{}
	for rows.Next() {
		zv := ZonaVerde{}
		rows.Scan(&zv.IDZonaVerde, &zv.Nombre, &zv.IDDistrito, &zv.Direccion)
		zonasverdes = append(zonasverdes, zv)
	}
	return zonasverdes, nil
}
