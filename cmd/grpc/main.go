package main

import (
	"github.com/PandeKaustubhS/microservice-todo/config"
	logger "github.com/PandeKaustubhS/microservice-todo/middleware"
	model "github.com/PandeKaustubhS/microservice-todo/model"
	repo "github.com/PandeKaustubhS/microservice-todo/repository"
	"github.com/PandeKaustubhS/microservice-todo/server"
	usecase "github.com/PandeKaustubhS/microservice-todo/usecase/handler"
	// pb "github.com/PandeKaustubhS/microservice-todo/usecase/user"
	log "github.com/sirupsen/logrus"
)

func init() {
	config.Load()
	logger.Setup()
}

func main() {
	conf := config.Db()
	// Creates a database connection and handles
	// closing it again before exit.
	db, err := model.CreateConnection(conf)
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Todo{})
	repository := &repo.UserRepository{Db: db}
	tokenService := &usecase.TokenService{Repo: repository}
	// TodoRepo := repository
	userService := &usecase.Service{Repo: repository, TokenService: tokenService}
	todoService := &usecase.TodoService{TodoRepo: repository}
	// Run the server
	if err := server.StartGRPCServer(userService, todoService); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
