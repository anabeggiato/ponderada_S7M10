package config

import (
	"log"
	"os"

	"github.com/anabeggiato/pond3_M10S07/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "figurinhas.db"
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}

	err = db.AutoMigrate(&domain.Figurinha{})
	if err != nil {
		log.Fatal("erro ao criar tabelas:", err)
	}

	return db
}
