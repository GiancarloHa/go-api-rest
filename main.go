package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/giancarloha/go-rest-api/database"
	"github.com/giancarloha/go-rest-api/routes"
	"github.com/giancarloha/go-rest-api/telemetry"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
	telemetry.InitTelemetry()
	database.ConectaDB()
	fmt.Println("Iniciando o servidor")
	routes.HandleRequest()
}
