package components

import (
	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/raptor/v3"
)

func Services(utils *raptor.Utils) raptor.Services {
	return raptor.Services{
		services.NewWorkerService(4),
		services.NewNATSService(utils.Config),
		services.NewAuthService(utils.Config),
		&services.GroupsService{},
		&services.ClientsService{},
		&services.LogsService{},
	}
}
