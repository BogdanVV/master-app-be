package service

import (
	"github.com/bogdanvv/master-app-be/models"
	"github.com/bogdanvv/master-app-be/repo"
)

type Todos struct {
	repo *repo.Repo
}

func NewTodos(r *repo.Repo) *Todos {
	return &Todos{repo: r}
}

func (s *Todos) CreateTodo(input models.TodoCreateBody, userId string) (models.TodoResponseBody, error) {
	return s.repo.CreateTodo(input, userId)
}

func (s *Todos) GetAllTodos(userId string) ([]models.TodoResponseBody, error) {
	return s.repo.GetAllTodos(userId)
}

func (s *Todos) GetTodoById(id int, userId string) (models.TodoResponseBody, error) {
	return s.repo.GetTodoById(id, userId)
}

func (s *Todos) DeleteTodoById(id int, userId string) error {
	return s.repo.DeleteTodoById(id, userId)
}

func (s *Todos) UpdateTodoById(id int, input models.TodoUpdateBody, userId string) (models.TodoResponseBody, error) {
	return s.repo.UpdateTodoById(id, input, userId)
}
