package service

import (
	"github.com/bogdanvv/master-app-be/models"
	"github.com/bogdanvv/master-app-be/repo"
)

type Service struct {
	AuthRepo
}

func NewService(repository repo.Repo) *Service {
	return &Service{
		AuthRepo: NewAuthService(repository),
	}
}

type AuthRepo interface {
	Signup(name, email, password string) (string, error)
	Login(email, password string) (models.LoginResponse, error)
}
