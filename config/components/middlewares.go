package components

import (
	"github.com/go-logistiq/api/app/middlewares"
	"github.com/go-raptor/raptor/v3"
	"github.com/go-raptor/raptor/v3/core"
)

func Middlewares(utils *raptor.Utils) raptor.Middlewares {
	return raptor.Middlewares{
		core.UseExcept(&middlewares.AuthMiddleware{}, "Auth.Login"),
	}
}
