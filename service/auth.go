package service

import (
	"fmt"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/bogdanvv/master-app-be/repo"
	"github.com/bogdanvv/master-app-be/utils"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repo *repo.Repo
}

func NewAuth(repository *repo.Repo) *Auth {
	return &Auth{
		repo: repository,
	}
}

func (s *Auth) Signup(name, email, password string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return s.repo.CreateUser(name, email, string(hashedPasswordBytes))
}

func (s *Auth) Login(email, password string) (models.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return models.LoginResponse{}, err
	}

	if err := utils.VerifyUserPassword(user.Password, password); err != nil {
		return models.LoginResponse{}, fmt.Errorf("invalid password")
	}

	tokenString, err := utils.GenerateAccessToken(user.Id)
	if err != nil {
		return models.LoginResponse{}, fmt.Errorf("failed to create an accessToken")
	}

	return models.LoginResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		AccessToken: tokenString,
	}, nil
}
