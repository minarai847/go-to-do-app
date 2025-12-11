package main

import (
	"os"

	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	database := db.NewDB()
	
	// データベースマイグレーションを実行
	if err := database.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
	
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(database)
	taskRepository := repository.NewTaskRepository(database)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)

	// Render では PORT が環境変数として渡される
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // ローカル開発用
	}

	// サーバー起動
	e.Logger.Fatal(e.Start(":" + port))
}
