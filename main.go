package main

import (
	"go-clean-architecture/controller"
	"go-clean-architecture/db"
	"go-clean-architecture/repository"
	"go-clean-architecture/router"
	"go-clean-architecture/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
