package models

//Persona modelo de persona
type Persona struct {
	IDPersona       int    `json:"idpersona"`
	Nombre          string `json:"nombre"`
	Paterno         string `json:"paterno"`
	Materno         string `json:"materno"`
	CarnetIdentidad string `json:"carnetidentidad"`
	Telefono        string `json:"telefono"`
	Celular         string `json:"celular"`
}

//Personas Array de personas
type Personas []Persona

var queryPersona = `CREATE TABLE if NOT EXISTS persona(
	idpersona int primary KEY not null AUTO_INCREMENT,
	nombre varchar(45),
	paterno varchar(45),
	materno varchar(45),
	carnetidentidad varchar(12),
	telefono varchar(12),
	celular varchar(12)
)`

//CrearTablaPersona crea una tabla persona
func CrearTablaPersona() {
	EjecutarExec(queryPersona)
}

//AgregarUsuario agrega un usuario a la tabla persona
func (p *Persona) AgregarUsuario() error {
	query := `INSERT INTO persona(nombre,paterno,materno,carnetidentidad,telefono,celular) VALUES(?,?,?,?,?,?)`
	if _, err := EjecutarExec(query, &p.Nombre, &p.Paterno, &p.Materno, &p.CarnetIdentidad, &p.Telefono, &p.Celular); err != nil {
		return err
	}
	return nil
}

//ObtenerUsuario devuelve un usuario por su id
func ObtenerUsuario(id int) (*Persona, error) {
	query := `SELECT idpersona,nombre,paterno,materno,carnetidentidad,telefono,celular FROM persona
		WHERE idpersona=?`
	rows, err := EjecutarQuery(query, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	persona := &Persona{}

	if rows.Next() {
		rows.Scan(&persona.IDPersona, &persona.Nombre, &persona.Paterno, &persona.Materno, &persona.CarnetIdentidad, &persona.Telefono, &persona.Celular)
	}
	return persona, nil
}

//ObtenerUsuarioByCI Obtiene un usuario por su ci
func ObtenerUsuarioByCI(ci string) (*Persona, error) {
	query := `SELECT idpersona,nombre,paterno,materno,carnetidentidad,telefono,celular FROM persona
		WHERE carnetidentidad=?`
	rows, err := EjecutarQuery(query, ci)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	persona := &Persona{}

	if rows.Next() {
		rows.Scan(&persona.IDPersona, &persona.Nombre, &persona.Paterno, &persona.Materno, &persona.CarnetIdentidad, &persona.Telefono, &persona.Celular)
	}
	return persona, nil
}

//ObtenerUsuarios devuleve todos los usuarios de una tabla
func ObtenerUsuarios() (Personas, error) {
	query := `SELECT idpersona,nombre,paterno,materno,carnetidentidad,telefono,celular FROM persona`

	rows, err := EjecutarQuery(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	personas := Personas{}
	for rows.Next() {
		persona := Persona{}
		rows.Scan(&persona.IDPersona, &persona.Nombre, &persona.Paterno, &persona.Materno, &persona.CarnetIdentidad, &persona.Telefono, &persona.Celular)
		personas = append(personas, persona)
	}
	return personas, nil

}

//ActualizarPersona actualiza los datos de una persona
func (p *Persona) ActualizarPersona() error {
	query := `UPDATE persona SET nombre=?,paterno=?,materno=?,carnetidentidad=?,telefono=?,celular=? WHERE idpersona=?`
	_, err := EjecutarExec(query, &p.Nombre, &p.Paterno, &p.Materno, &p.CarnetIdentidad, &p.Telefono, &p.Celular, &p.IDPersona)
	if err != nil {
		return err
	}
	return nil
}

//EliminarPersona elimina una persona de la base de datos
func EliminarPersona(id int) error {
	query := `DELETE FROM persona WHERE idpersona=?`
	_, err := EjecutarExec(query, id)
	if err != nil {
		return err
	}
	return nil
}
