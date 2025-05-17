package components

import (
	"github.com/go-logistiq/api/app/middlewares"
	"github.com/go-raptor/raptor/v4"
	"github.com/go-raptor/raptor/v4/core"
)

func Middlewares(c *raptor.Config) raptor.Middlewares {
	return raptor.Middlewares{
		core.UseExcept(&middlewares.AuthMiddleware{}, "Auth.Login"),
	}
}
