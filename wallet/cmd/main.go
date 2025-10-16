package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"wallet/internal/domain/service"
	"wallet/internal/infrastructure/repository"
	"wallet/internal/server/gen"
	"wallet/internal/server/handlers"
	"wallet/pkg/connectors"
)

func main() {
	db := connectors.Postgres{DSN: os.Getenv("DATABASE_URL")}
	ctx := context.Background()
	client := db.Client(ctx)

	walletRepo := repository.NewWalletRepository(client)

	walletService := service.NewWalletService(&walletRepo)

	handler := handlers.NewWalletHandler(walletService)

	e := echo.New()
	gen.RegisterHandlers(e, handler)

	log.Println("Server starting on :8080")
	e.Start(":8080")
}
