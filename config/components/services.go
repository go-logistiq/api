package components

import (
	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/raptor/v4"
)

func Services(c *raptor.Config) raptor.Services {
	return raptor.Services{
		services.NewWorkerService(c),
		services.NewNATSService(c),
		services.NewAuthService(c),
		&services.GroupsService{},
		services.NewClientsService(),
		&services.LogsService{},
	}
}
