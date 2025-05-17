package controllers

import (
	"github.com/go-logistiq/api/app/models"
	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/raptor/v4"
	"github.com/go-raptor/raptor/v4/errs"
)

type AuthController struct {
	raptor.Controller

	Auth *services.AuthService
}

func (ac *AuthController) Login(c *raptor.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return errs.NewErrorBadRequest("Invalid request body")
	}

	var err error
	if user, err = ac.Auth.Login(user); err != nil {
		return errs.NewErrorUnauthorized("Invalid credentials")
	}

	return c.Data(user)
}
