package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func VerifyUserPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}

func GenerateAccessToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 720).Unix(), // 30 days
	})

	return token.SignedString([]byte(os.Getenv("ACESS_TOKEN_SECRET")))
}

func GenerateHashedPassword(rawPassword string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	return string(hashedPasswordBytes), err
}
