package service

import (
	"github.com/bogdanvv/master-app-be/models"
	"github.com/bogdanvv/master-app-be/repo"
)

type Service struct {
	AuthService
	UsersService
}

func NewService(repository *repo.Repo) *Service {
	return &Service{
		AuthService:  NewAuth(repository),
		UsersService: NewUsers(repository),
	}
}

type AuthService interface {
	Signup(name, email, password string) (string, error)
	Login(email, password string) (models.LoginResponse, error)
}

type UsersService interface {
	UpdateUser(userId string, updateBody models.UserUpdateBody) (models.UserResponse, error)
}
