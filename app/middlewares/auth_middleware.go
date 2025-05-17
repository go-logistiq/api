package middlewares

import (
	"strings"

	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/raptor/v4"
	"github.com/go-raptor/raptor/v4/errs"
)

type AuthMiddleware struct {
	raptor.Middleware

	Auth *services.AuthService
}

func (am *AuthMiddleware) Handle(c *raptor.Context, next func(*raptor.Context) error) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		am.Log.Debug("Missing auth header", "ip", c.RealIP())
		return errs.NewErrorUnauthorized("Missing auth header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		am.Log.Debug("Invalid auth header", "ip", c.RealIP())
		return errs.NewErrorUnauthorized("Invalid auth header")
	}

	if authKey := parts[1]; authKey != am.Auth.Token {
		am.Log.Debug("Invalid auth token", "ip", c.RealIP())
		return errs.NewErrorUnauthorized("Invalid auth token")
	}

	return next(c)
}
