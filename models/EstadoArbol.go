package models

//EstadoArbol modelo de el estado arbol
type EstadoArbol struct {
	IDEstadoArbol int    `json:"idestadoarbol"`
	Descripcion   string `json:"descripcion"`
	DirFoto       string `json:"dirfoto"`
	IDArbol       string `json:"idarbol"`
}

var queryEstadoArbol = `CREATE TABLE if NOT EXISTS EstadoArbol(
	idestadoarbol int primary KEY not null AUTO_INCREMENT,
	descripcion varchar(45),
	dirfoto varchar(60),
	arbol_id int,
	CONSTRAINT fk_EstadoArbol_arbol
		FOREIGN KEY (arbol_id) REFERENCES EstadoArbol(idarbol)
		ON DELETE CASCADE
	)`
