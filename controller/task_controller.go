package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"reflect"
	"strconv"

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

// JWTトークンからユーザーIDを取得するヘルパー関数
// echo-jwtが設定するuserの型の問題を回避するため
func getUserID(c echo.Context) (uint, error) {
	user := c.Get("user")
	if user == nil {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "User not found")
	}

	// リフレクションを使ってClaimsフィールドにアクセス
	userValue := reflect.ValueOf(user)

	// ポインタの場合は間接参照
	if userValue.Kind() == reflect.Ptr {
		if userValue.IsNil() {
			return 0, echo.NewHTTPError(http.StatusUnauthorized, "User token is nil")
		}
		userValue = userValue.Elem()
	}

	// Claimsフィールドを取得
	claimsField := userValue.FieldByName("Claims")
	if !claimsField.IsValid() {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "Claims field not found in token")
	}

	// Claimsの値を取得
	claimsValue := claimsField.Interface()

	// map[string]interface{}に変換を試みる
	var claims map[string]interface{}

	// まず直接型アサーションを試す
	if claimsMap, ok := claimsValue.(map[string]interface{}); ok {
		claims = claimsMap
	} else {
		// 型アサーションが失敗した場合、リフレクションで直接アクセス
		claimsReflect := reflect.ValueOf(claimsValue)
		if claimsReflect.Kind() == reflect.Map {
			claims = make(map[string]interface{})
			for _, key := range claimsReflect.MapKeys() {
				claims[key.String()] = claimsReflect.MapIndex(key).Interface()
			}
		} else {
			return 0, echo.NewHTTPError(http.StatusUnauthorized, "Invalid claims type")
		}
	}

	// user_idを取得
	userIDValue, exists := claims["user_id"]
	if !exists {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "User ID not found in token")
	}

	// float64に変換（JSONの数値はfloat64として扱われる）
	userID, ok := userIDValue.(float64)
	if !ok {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "User ID is not a number")
	}

	return uint(userID), nil
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	userId, err := getUserID(c)
	if err != nil {
		return err
	}

	tasksRes, err := tc.tu.GetAllTasks(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	userId, err := getUserID(c)
	if err != nil {
		return err
	}

	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.GetTaskById(userId, uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	userId, err := getUserID(c)
	if err != nil {
		return err
	}

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserID = userId
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	userId, err := getUserID(c)
	if err != nil {
		return err
	}

	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	taskRes, err := tc.tu.UpdateTask(task, userId, uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	userId, err := getUserID(c)
	if err != nil {
		return err
	}

	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	err = tc.tu.DeleteTask(userId, uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
