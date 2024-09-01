package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("banco.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados", err)
	}
}
