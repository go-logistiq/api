package main

import (
	"github.com/go-logistiq/api/config"
	"github.com/go-logistiq/api/config/components"
	"github.com/go-raptor/raptor/v4"
)

func main() {
	app := raptor.New()

	app.Configure(components.New(app.Core.Resources.Config))
	app.RegisterRoutes(config.Routes())
	app.Run()
}
