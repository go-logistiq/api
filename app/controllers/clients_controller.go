package controllers

import (
	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/raptor/v3"
)

type ClientsController struct {
	raptor.Controller

	Clients *services.ClientsService
}

func (gc *ClientsController) All(c *raptor.Context) error {
	clients, err := gc.Clients.All()
	if err != nil {
		return c.JSONError(err)
	}
	return c.JSONResponse(clients)
}

func (gc *ClientsController) GetBySlug(c *raptor.Context) error {
	slug := c.Param("slug")
	client, err := gc.Clients.GetBySlug(slug)
	if err != nil {
		return c.JSONError(err)
	}
	return c.JSONResponse(client)
}
