package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //Drivers que ayuda a la conexion con la base de datos
)

var username = "root"
var contrasenia = ""
var host = "localhost"
var puerto = 3306
var dbname = "Arbolito"

var db *sql.DB

func init() {
	conectar()
}

func conectar() {
	if GetConnection() != nil {
		return
	}
	coneccion, err := sql.Open("mysql", generarURL())
	if err != nil {
		panic(err)
	} else {
		db = coneccion
	}
}

//EjecutarExec ejecuta una sentencia
func EjecutarExec(query string, args ...interface{}) (sql.Result, error) {
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}

//EjecutarQuery ejecuta un query y devuelve filas de una tabla
func EjecutarQuery(query string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rows
}

//Ping hace ping a la base de datos
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//Cerrar cierra la conexion de la base de datos
func Cerrar() {
	db.Close()
}

func generarURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, contrasenia, host, puerto, dbname)
}

//GetConnection obtiene la conexion
func GetConnection() *sql.DB {
	return db
}
