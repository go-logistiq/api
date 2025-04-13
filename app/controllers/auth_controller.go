package controllers

import (
	"github.com/go-logistiq/api/app/models"
	"github.com/go-logistiq/api/app/services"
	"github.com/go-raptor/raptor/v3"
)

type AuthController struct {
	raptor.Controller

	Auth *services.AuthService
}

func (ac *AuthController) Login(s raptor.State) error {
	var user models.User
	if err := s.Bind(&user); err != nil {
		return s.JSONError(err, 400)
	}

	var err error
	if user, err = ac.Auth.Login(user); err != nil {
		return s.JSONError(err, 401)
	}

	return s.JSONResponse(user)
}
