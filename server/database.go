package server

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("cotacao.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	if err := db.AutoMigrate(&Cotacao{}); err != nil {
		log.Fatal("Erro ao migrar banco:", err)
	}

	return db
}
