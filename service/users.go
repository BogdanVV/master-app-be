package service

import (
	"fmt"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/bogdanvv/master-app-be/repo"
)

type Users struct {
	repo *repo.Repo
}

func NewUsers(repository *repo.Repo) *Users {
	return &Users{repo: repository}
}

func (s *Users) UpdateUser(userId string, userUpdateBody models.UserUpdateBody) (models.UserResponse, error) {
	_, err := s.repo.GetUserById(userId)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("user not found")
	}

	return s.repo.UpdateUser(userId, userUpdateBody)
}
