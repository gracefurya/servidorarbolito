package models

//EstadoArbol modelo de el estado arbol
type EstadoArbol struct {
	IDEstadoArbol int    `json:"idestadoarbol"`
	Descripcion   string `json:"descripcion"`
	DirFoto       string `json:"dirfoto"`
	IDArbol       string `json:"idarbol"`
}
