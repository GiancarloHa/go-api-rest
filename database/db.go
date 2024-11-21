package database

import (
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type User struct {
	gorm.Model
	Name string
}

func ConectaDB() {
	dsn := os.Getenv("db-connection")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Erro ao conectar no banco de dados")
	}
}
