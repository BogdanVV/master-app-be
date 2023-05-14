package repo

import (
	"fmt"
	"strings"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/jmoiron/sqlx"
)

type Users struct {
	db *sqlx.DB
}

func NewUsers(db *sqlx.DB) *Users {
	return &Users{db: db}
}

func (r *Users) GetUserById(id string) (models.UserResponse, error) {
	var user models.UserResponse
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id=$1"
	err := r.db.Get(&user, query, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return user, err
}

func (r *Users) UpdateUser(id string, updateBody models.UserUpdateBody) (models.UserResponse, error) {
	var userResponse models.UserResponse
	updateChunks := []string{}

	if updateBody.Email == "" && updateBody.Name == "" {
		return models.UserResponse{}, fmt.Errorf("empty values")
	}
	if updateBody.Email != "" {
		updateChunks = append(updateChunks, fmt.Sprintf("email='%s'", updateBody.Email))
	}
	if updateBody.Name != "" {
		updateChunks = append(updateChunks, fmt.Sprintf("name='%s'", updateBody.Name))
	}
	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$1 RETURNING id, name, email, created_at, updated_at", strings.Join(updateChunks, ", "))
	row := r.db.QueryRowx(query, id)
	err := row.StructScan(&userResponse)
	if err != nil {
		return models.UserResponse{}, err
	}

	return userResponse, nil
}
