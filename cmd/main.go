package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/Regular-Pashka/todo-app"
	"github.com/Regular-Pashka/todo-app/pkg/handler"
	"github.com/Regular-Pashka/todo-app/pkg/repository"
	"github.com/Regular-Pashka/todo-app/pkg/service"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs $s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}	
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

//migrate -path ./schema -database 'postgres://postgres:qwert@localhost:5436/postgres?sslmode=disable' up