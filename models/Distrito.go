package models

//Distrito modelo distrito
type Distrito struct {
	IDDistrito int    `json:"iddistrito"`
	Nombre     string `json:"nombre"`
	Codigo     int    `json:"codigo"`
}

//Distritos Array de distritos
type Distritos []Distrito

var queryDistrito = `CREATE TABLE if NOT EXISTS distrito(
	iddistrito int primary KEY not null AUTO_INCREMENT,
	nombre varchar(45),
	codigo int
)`

//CrearTablaDistrito Crea la tabla de distrito en la base de datos
func CrearTablaDistrito() {
	EjecutarExec(queryDistrito)
}

//AgregarDistrito Agrega un distrito en la base de datos
func (d *Distrito) AgregarDistrito() error {
	query := `INSERT INTO distrito(iddistrito,nombre,codigo) VALUES (?,?,?)`
	_, err := EjecutarExec(query, &d.IDDistrito, &d.Nombre, &d.Codigo)
	if err != nil {
		return err
	}
	return nil
}

//ObtenerDistrito devuelve un array de distritos
func ObtenerDistrito() (Distritos, error) {
	query := `SELECT iddistrito,nombre,codigo FROM distrito`
	rows, err := EjecutarQuery(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	distritos := Distritos{}
	for rows.Next() {
		distrito := Distrito{}
		rows.Scan(&distrito.IDDistrito, &distrito.Nombre, &distrito.Codigo)
		distritos = append(distritos, distrito)
	}
	return distritos, nil
}
