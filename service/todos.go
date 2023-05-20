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

func (s *Todos) CreateTodo(input models.TodoCreateBody) (models.TodoResponseBody, error) {
	return s.repo.CreateTodo(input)
}

func (s *Todos) GetAllTodos() ([]models.TodoResponseBody, error) {
	return s.repo.GetAllTodos()
}

func (s *Todos) GetTodoById(id int) (models.TodoResponseBody, error) {
	return s.repo.GetTodoById(id)
}

func (s *Todos) DeleteTodoById(id int) error {
	return s.repo.DeleteTodoById(id)
}

func (s *Todos) UpdateTodoById(id int, input models.TodoUpdateBody) (models.TodoResponseBody, error) {
	return s.repo.UpdateTodoById(id, input)
}
