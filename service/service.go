package service

import (
	"github.com/bogdanvv/master-app-be/models"
	"github.com/bogdanvv/master-app-be/repo"
)

type Service struct {
	AuthService
	UsersService
	TodosService
}

func NewService(repository *repo.Repo) *Service {
	return &Service{
		AuthService:  NewAuth(repository),
		UsersService: NewUsers(repository),
		TodosService: NewTodos(repository),
	}
}

type AuthService interface {
	Signup(name, email, password string) (models.UserResponse, error)
	Login(email, password string) (models.LoginResponse, error)
	RefreshAccessTokenToken(token string) (string, error)
}

type UsersService interface {
	UpdateUser(userId string, updateBody models.UserUpdateBody) (models.UserResponse, error)
	GetUserById(userId string) (models.UserResponse, error)
}

type TodosService interface {
	CreateTodo(input models.TodoCreateBody, userId string) (models.TodoResponseBody, error)
	GetAllTodos(userId string) ([]models.TodoResponseBody, error)
	GetTodoById(id int, userId string) (models.TodoResponseBody, error)
	DeleteTodoById(id int, userId string) error
	UpdateTodoById(id int, input models.TodoUpdateBody, userId string) (models.TodoResponseBody, error)
}
