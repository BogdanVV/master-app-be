package repo

import (
	"fmt"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(name, email, password string) (models.UserResponse, error) {
	createUserQuery := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	rows := r.db.QueryRow(createUserQuery, name, email, password)

	var user models.UserResponse
	var newUserId string
	err := rows.Scan(&newUserId)
	if err != nil {
		return user, err
	}

	query := "SELECT id, name, email, profile_image_url, created_at, updated_at FROM users WHERE id=$1"
	if err := r.db.Get(&user, query, newUserId); err != nil {
		return user, fmt.Errorf("failed to fetch data about the user")
	}

	return user, nil
}

func (r *Auth) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	queryString := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email=$1"
	if err := r.db.Get(&user, queryString, email); err != nil {
		return user, fmt.Errorf("user with email %s does not exist", email)
	}

	return user, nil
}
