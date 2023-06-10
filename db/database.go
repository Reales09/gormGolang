package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:123456@tcp(localhost:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexión", err)
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
		return db
	}
}()

//username: password@tcp(localhost:3306)/database

const url = "root:123456@tcp(localhost:3306)/goweb_db"

// Guarda la conexión
var db *sql.DB

// Realiza la conexión
func Connect() {
	conection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Conexion exitosa")
	db = conection
}

// Cerrar la conexión
func Close() {
	fmt.Println("Conexion cerrada exitosa")
	db.Close()

}

//Verificar la conexión

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//Verifica si una tabla existe o no

func ExisteTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()

}

// Crea una tabla

func CreateTable(schema string, name string) {

	if !ExisteTable(name) {
		_, err := db.Exec(schema)

		if err != nil {

			fmt.Println(err)
		}
	}

}

//Reiniciar el registro de una tabla

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return result, err

}

// Polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()

	if err != nil {
		fmt.Println(err)
	}

	return rows, err

}
