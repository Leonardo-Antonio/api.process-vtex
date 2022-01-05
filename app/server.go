package app

import (
	"fmt"
	"log"

	"github.com/Leonardo-Antonio/api.process-vtex/router"
	"github.com/Leonardo-Antonio/api.process-vtex/util/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	app *echo.Echo
}

func newServer() *server {
	return &server{app: echo.New()}
}

func (s *server) Middlewares() {
	s.app.Use(middleware.Logger())
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.CORS())
}

func (s *server) Routers() {
	router.Inventory(s.app)
}

func (s *server) Listeing() {
	if err := s.app.Start(fmt.Sprintf(":%s", env.PORT)); err != nil {
		log.Fatalln(err)
	}
}
