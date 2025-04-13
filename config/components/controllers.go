package components

import (
	"github.com/go-logistiq/api/app/controllers"
	"github.com/go-raptor/raptor/v3"
)

func Controllers() raptor.Controllers {
	return raptor.Controllers{
		&controllers.AuthController{},
		&controllers.GroupsController{},
		&controllers.ClientsController{},
	}
}
