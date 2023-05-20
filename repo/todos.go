package repo

import (
	"fmt"
	"strings"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/jmoiron/sqlx"
)

type Todos struct {
	db *sqlx.DB
}

func NewTodos(db *sqlx.DB) *Todos {
	return &Todos{db: db}
}

func (r *Todos) CreateTodo(input models.TodoCreateBody) (models.TodoResponseBody, error) {
	var newTodo models.TodoResponseBody
	fields := []string{"title", "user_id"}
	values := []string{
		fmt.Sprintf("'%s'", input.Title),
		fmt.Sprintf("'%s'", input.UserId),
	}
	if input.Description != "" {
		fields = append(fields, "description")
		values = append(values, fmt.Sprintf("'%s'", input.Description))
	}
	if input.Status != "" {
		fields = append(fields, "status")
		values = append(values, fmt.Sprintf("'%s'", input.Status))
	}
	if len(input.ActiveDays) > 0 {
		fields = append(fields, "active_days")
		activeDays := []string{}
		for _, day := range input.ActiveDays {
			activeDays = append(activeDays, fmt.Sprintf("'%s'", day))
		}
		values = append(
			values,
			fmt.Sprintf("ARRAY[%s]::text[]::day_of_week[]", strings.Join(activeDays, ",")),
		)
	}
	if input.Priority != "" {
		fields = append(fields, "priority")
		values = append(values, fmt.Sprintf("'%s'", input.Priority))
	}
	if input.IsDaily {
		fields = append(fields, "is_daily")
		values = append(values, "true")
	}

	query := fmt.Sprintf(
		"INSERT INTO todos (%s) VALUES (%s) RETURNING %s",
		strings.Join(fields, ","),
		strings.Join(values, ","),
		"id, title, description, status, active_days, priority, is_daily, user_id, created_at, updated_at",
	)
	row := r.db.QueryRowx(query)
	err := row.StructScan(&newTodo)

	return newTodo, err
}

func (r *Todos) GetAllTodos() ([]models.TodoResponseBody, error) {
	var response []models.TodoResponseBody
	query := "SELECT id, title, description, status, active_days, priority, is_daily, created_at, updated_at, user_id FROM todos"
	err := r.db.Select(&response, query)

	return response, err
}

func (r *Todos) GetTodoById(id int) (models.TodoResponseBody, error) {
	var wantedTodo models.TodoResponseBody
	query := "SELECT id, title, description, status, active_days, priority, is_daily, created_at, updated_at, user_id FROM todos WHERE id=$1"
	err := r.db.Get(&wantedTodo, query, id)

	return wantedTodo, err
}

func (r *Todos) DeleteTodoById(id int) error {
	res, err := r.db.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}

	return err
}

func (r *Todos) UpdateTodoById(id int, input models.TodoUpdateBody) (models.TodoResponseBody, error) {
	var response models.TodoResponseBody
	todo, err := r.GetTodoById(id)
	if err != nil {
		return response, err
	}

	updateCols := []string{}
	updateVals := []string{}
	if input.Title != "" {
		updateCols = append(updateCols, "title")
		updateVals = append(updateVals, fmt.Sprintf("'%s'", input.Title))
	}
	if input.Description != "" {
		updateCols = append(updateCols, "description")
		updateVals = append(updateVals, fmt.Sprintf("'%s'", input.Description))
	}
	if input.Status != "" {
		updateCols = append(updateCols, "status")
		updateVals = append(updateVals, fmt.Sprintf("'%s'", input.Status))
	}
	if len(input.ActiveDays) > 0 {
		updateCols = append(updateCols, "active_days")
		newDays := []string{}
		for _, day := range input.ActiveDays {
			newDays = append(newDays, fmt.Sprintf("'%s'", day))
		}
		updateVals = append(updateVals, fmt.Sprintf("ARRAY[%s]::text[]::day_of_week[]", strings.Join(newDays, ",")))
	}
	if input.Priority != "" {
		updateCols = append(updateCols, "priority")
		updateVals = append(updateVals, fmt.Sprintf("'%s'", input.Priority))
	}
	if input.IsDaily != todo.IsDaily {
		updateCols = append(updateCols, "is_daily")
		newValue := "false"
		if input.IsDaily {
			newValue = "true"
		}
		updateVals = append(updateVals, newValue)
	}

	if len(updateCols) != len(updateVals) {
		return response, fmt.Errorf("invalid data")
	}
	if len(updateCols) == 0 || len(updateVals) == 0 {
		return response, fmt.Errorf("empty body")
	}

	updateChunks := []string{}
	for i := 0; i < len(updateCols); i++ {
		updateChunks = append(updateChunks, fmt.Sprintf("%s=%s", updateCols[i], updateVals[i]))
	}

	query := fmt.Sprintf(
		"UPDATE todos SET %s WHERE id=$1 RETURNING %s",
		strings.Join(updateChunks, ","),
		"id, title, description, status, active_days, priority, is_daily, user_id, created_at, updated_at",
	)
	row := r.db.QueryRowx(query, id)
	err = row.StructScan(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}
