package controller

import (
	usecase "go-rest-api"
	"go-rest-api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
}

type userContoroller struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userContoroller{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
}
