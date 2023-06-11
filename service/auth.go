package service

import (
	"fmt"
	"os"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/bogdanvv/master-app-be/repo"
	"github.com/bogdanvv/master-app-be/utils"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	repo *repo.Repo
}

func NewAuth(repository *repo.Repo) *Auth {
	return &Auth{
		repo: repository,
	}
}

func (s *Auth) Signup(name, email, password string) (models.UserResponse, error) {
	hashedPasswordBytes, err := utils.GenerateHashedPassword(password)
	if err != nil {
		return models.UserResponse{}, err
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
		User: models.UserResponse{
			Id:              user.Id,
			Name:            user.Name,
			Email:           user.Email,
			ProfileImageURL: user.ProfileImageURL,
			CreatedAt:       user.CreatedAt,
			UpdatedAt:       user.UpdatedAt,
		},
		AccessToken: tokenString,
	}, nil
}

func (c *Auth) RefreshAccessTokenToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("invalid token")
		}

		return []byte(os.Getenv("ACESS_TOKEN_SECRET")), nil
	})
	if err != nil {
		return "", fmt.Errorf("could not parse the token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := claims.GetSubject()
		if err != nil {
			return "", fmt.Errorf("the token is mailformed")
		}
		newToken, err := utils.GenerateAccessToken(userId)
		if err != nil {
			return "", fmt.Errorf("could not generate a token")
		}
		return newToken, nil
	}

	return "", fmt.Errorf("unknown error occured")
}
