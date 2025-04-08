package controllers

import (
	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/raptor/v3"
)

type ClientsController struct {
	raptor.Controller

	Clients *services.ClientsService
}

func (gc *ClientsController) All(s raptor.State) error {
	clients, err := gc.Clients.All()
	if err != nil {
		return s.JSONError(err)
	}
	return s.JSONResponse(clients)
}

func (gc *ClientsController) GetBySlug(s raptor.State) error {
	groupSlug := s.Param("group")
	clientSlug := s.Param("client")
	client, err := gc.Clients.GetBySlug(groupSlug, clientSlug)
	if err != nil {
		return s.JSONError(err)
	}
	return s.JSONResponse(client)
}
