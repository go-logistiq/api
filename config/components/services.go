package components

import (
	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/raptor/v3"
)

func Services(utils *raptor.Utils) raptor.Services {
	return raptor.Services{
		services.NewWorkerService(utils.Config),
		services.NewNATSService(utils.Config),
		services.NewAuthService(utils.Config),
		&services.GroupsService{},
		&services.ClientsService{},
		&services.LogsService{},
	}
}
