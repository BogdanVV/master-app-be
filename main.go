package main

import (
	"log"

	"github.com/bogdanvv/master-app-be/controller"
	"github.com/bogdanvv/master-app-be/repo"
	"github.com/bogdanvv/master-app-be/routing"
	"github.com/bogdanvv/master-app-be/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	// load config
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env: %s", err.Error())
	}

	// connect to db
	db, err := repo.ConnectToDB()
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}

	// create controller, service and repo
	repos := repo.NewRepo(db)
	s := service.NewService(*repos)
	controller := controller.NewController(s)

	// init routing
	router := routing.InitRoutes(*controller)

	// run server
	port := viper.GetString("server.port")
	router.Run(port)
}
