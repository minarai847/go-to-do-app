package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	// JWTトークンからユーザー情報を取得
	user := c.Get("user").(*jwt.Token)
	// トークンのClaims（情報）を取得
	claims := user.Claims.(jwt.MapClaims)
	// user_idを取得してuint型に変換
	userId := uint(claims["user_id"].(float64))

	// UseCase層でタスク一覧を取得
	taskRes, err := tc.tu.GetAllTasks(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	// JWTトークンからユーザー情報を取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	// URLパラメータからタスクIDを取得
	id := c.Param("taskId")
	taskid, _ := strconv.Atoi(id)

	// UseCase層で特定のタスクを取得
	taskRes, err := tc.tu.GetTaskById(userId, uint(taskid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	// JWTトークンからユーザー情報を取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["user_id"]

	// リクエストボディからタスク情報を取得
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// タスクにユーザーIDを設定
	task.UserID = uint(user_id.(float64))

	// UseCase層でタスクを作成
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	// JWTトークンからユーザー情報を取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["user_id"]

	// URLパラメータからタスクIDを取得
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// リクエストボディからタスク情報を取得
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// UseCase層でタスクを更新
	taskRes, err := tc.tu.UpdateTask(task, uint(user_id.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	// JWTトークンからユーザー情報を取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["user_id"]

	// URLパラメータからタスクIDを取得
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// UseCase層でタスクを削除
	if err := tc.tu.DeleteTask(uint(user_id.(float64)), uint(taskId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
