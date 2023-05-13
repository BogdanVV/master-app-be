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
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		AuthRepo: NewAuth(db),
	}
}

type AuthRepo interface {
	CreateUser(name, email, password string) (string, error)
	GetUserByEmail(password string) (models.User, error)
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
