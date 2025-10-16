package server

import (
	"log"
	"wallet/internal/domain/service"
	"wallet/internal/server/gen"
	"wallet/internal/server/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
	port string
}

func NewServer(walletService service.WalletService, port string) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	walletHandler := handlers.NewWalletHandler(&walletService)

	gen.RegisterHandlers(e, walletHandler)

	return &Server{
		echo: e,
		port: port,
	}
}

func (s *Server) Start() error {
	log.Printf("Server starting on port %s", s.port)
	return s.echo.Start(":" + s.port)
}

func (s *Server) Shutdown() error {
	return s.echo.Close()
}
