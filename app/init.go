package app

import (
	"github.com/Leonardo-Antonio/api.process-vtex/util/env"
	"github.com/joho/godotenv"
)

func New() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	env.Load()

	app := newServer()
	app.Middlewares()
	app.Routers()
	app.Listeing()
}
