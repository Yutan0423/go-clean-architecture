package main

import (
	"go-clean-architecture/controller"
	"go-clean-architecture/db"
	"go-clean-architecture/repository"
	"go-clean-architecture/router"
	"go-clean-architecture/usecase"
	"go-clean-architecture/validator"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUseCase := usecase.NewUserUseCase(userRepository, userValidator)
	userController := controller.NewUserController(userUseCase)
	taskRepository := repository.NewTaskRepository(db)
	taskValidator := validator.NewTaskValidator()
	taskUseCase := usecase.NewTaskUseCase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUseCase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
