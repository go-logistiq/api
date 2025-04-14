package middlewares

import (
	"strings"

	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/errs"
	"github.com/go-raptor/raptor/v3"
)

type AuthMiddleware struct {
	raptor.Middleware

	Auth *services.AuthService
}

func (am *AuthMiddleware) New(s raptor.State, next func(raptor.State) error) error {
	authHeader := s.Request().Header.Get("Authorization")
	if authHeader == "" {
		am.Log.Debug("Missing auth header", "ip", s.RealIP())
		return errs.NewErrorUnauthorized("Missing auth header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		am.Log.Debug("Invalid auth header", "ip", s.RealIP())
		return errs.NewErrorUnauthorized("Invalid auth header")
	}

	if authKey := parts[1]; authKey != am.Auth.Token {
		am.Log.Debug("Invalid auth token", "ip", s.RealIP())
		return errs.NewErrorUnauthorized("Invalid auth token")
	}

	return next(s)
}
