package repo

import (
	"fmt"
	"os"

	"github.com/bogdanvv/master-app-be/models"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Repo struct {
	AuthRepo
	UsersRepo
	TodosRepo
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		AuthRepo:  NewAuth(db),
		UsersRepo: NewUsers(db),
		TodosRepo: NewTodos(db),
	}
}

type AuthRepo interface {
	CreateUser(name, email, password string) (models.UserResponse, error)
	GetUserByEmail(password string) (models.User, error)
}

type UsersRepo interface {
	GetUserById(id string) (models.UserResponse, error)
	UpdateUser(id string, updateBody models.UserUpdateBody) (models.UserResponse, error)
}

type TodosRepo interface {
	CreateTodo(input models.TodoCreateBody, userId string) (models.TodoResponseBody, error)
	GetAllTodos(userId string) ([]models.TodoResponseBody, error)
	GetTodoById(id int, userId string) (models.TodoResponseBody, error)
	DeleteTodoById(id int, userId string) error
	UpdateTodoById(id int, updateBody models.TodoUpdateBody, userId string) (models.TodoResponseBody, error)
}

func ConnectToDB() (*sqlx.DB, error) {
	user := viper.GetString("db.user")
	password := os.Getenv("DB_PASSWORD")
	host := viper.GetString("db.host")
	dbPort := viper.GetString("db.port")
	dbName := viper.GetString("db.name")
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, dbPort, user, dbName, password)

	return sqlx.Connect("postgres", connStr)
}
