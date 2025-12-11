package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// import (
// 	"go-rest-api/controller"
// 	"go-rest-api/db"
// 	"go-rest-api/repository"
// 	"go-rest-api/router"
// 	"go-rest-api/usecase"
// 	"go-rest-api/validator"
// )

func main() {
	// db := db.NewDB()
	// userValidator := validator.NewUserValidator()
	// taskValidator := validator.NewTaskValidator()
	// userRepository := repository.NewUserRepository(db)
	// taskRepository := repository.NewTaskRepository(db)
	// userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	// taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	// userController := controller.NewUserController(userUsecase)
	// taskController := controller.NewTaskController(taskUsecase)
	// e := router.NewRouter(userController, taskController)
	// e.Logger.Fatal(e.Start(":8080"))
	e := echo.New()

	// --- 動作確認用の仮エンドポイント ---
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running!")
	})

	// Render では PORT が環境変数として渡される
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // ローカル開発用
	}

	// サーバー起動 (Fatalで落とさない)
	e.Logger.Print("Starting server on port " + port)
	e.Start(":" + port)

}
