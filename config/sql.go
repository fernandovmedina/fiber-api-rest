package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Variable de tipo *gorm.DB que guarde lo relacionado a la base de datos
var Database *gorm.DB = ConnectDB()

// Variable que guarde la url de la base de datos
var DatabaseURL string = "root:Medina_04@tcp(localhost:3306)/api_rest?charset=utf8mb4&parseTime=True&loc=Local"

// Funcion para conectarse a la base de datos
func ConnectDB() *gorm.DB {
	if Database, err := gorm.Open(mysql.Open(DatabaseURL), &gorm.Config{}); err != nil {
		panic(err.Error())
	} else {
		return Database
	}
}
