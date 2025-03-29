package services

import (
	"errors"

	"github.com/go-logistiq/api/app/models"
	"github.com/go-raptor/raptor/v3"
)

type AuthService struct {
	raptor.Service

	Username string
	Password string
	Token    string
}

func NewAuthService(c *raptor.Config) *AuthService {
	return &AuthService{
		Username: c.AppConfig["auth_username"],
		Password: c.AppConfig["auth_password"],
		Token:    c.AppConfig["auth_token"],
	}
}

func (as *AuthService) Login(user models.User) (models.User, error) {
	if user.Username == as.Username && user.Password == as.Password {
		user.Token = as.Token
		return user, nil
	}

	return models.User{}, errors.New("invalid username or password")
}
