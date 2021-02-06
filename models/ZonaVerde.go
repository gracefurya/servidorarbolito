package models

//ZonaVerde modelo de la zona verde
type ZonaVerde struct {
	IDZonaVerde int    `json:"idzonaverde"`
	Nombre      string `json:"nombre"`
	Direccion   string `json:"direccion"`
	IDDistrito  int    `json:"iddistrito"`
}
