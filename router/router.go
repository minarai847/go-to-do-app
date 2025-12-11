package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	// CORS設定: ローカル開発環境と本番環境のURLを許可
	allowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:3001",
		"https://go-to-do-app.vercel.app", // VercelのフロントエンドURL
	}
	// 環境変数から追加のURLを取得
	if feURL := os.Getenv("FE_URL"); feURL != "" {
		allowedOrigins = append(allowedOrigins, feURL)
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowedOrigins,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	// CSRFミドルウェアを全体に適用（GETリクエストは自動的に許可される）
	// signupとloginはCSRF保護の対象外
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		Skipper: func(c echo.Context) bool {
			// signupとloginはCSRF保護の対象外
			return c.Path() == "/signup" || c.Path() == "/login"
		},
	}))
	// 認証前のエンドポイント（signup, login, csrf）はCSRF保護の対象外（GETは自動的に許可される）
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.GET("/csrf", uc.GetCSRFToken)
	// 認証後のエンドポイント
	auth := e.Group("")
	auth.POST("/logout", uc.LogOut)
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
