package routes

import (
	"middleware-prioritas-1/constant"
	"middleware-prioritas-2/controllers"
	myMiddleware "middleware-prioritas-2/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	myMiddleware.LogMiddleware(e)
	myMiddleware.RateLimitterMiddleware(e)
	e.POST("/api/v1/register", controllers.RegisterUserController)
	e.POST("/api/v1/login", controllers.LoginUserController)
	eJwt := e.Group("/")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.GET("api/v1/posts", controllers.GetAllPostsController)
	eJwt.GET("api/v1/posts/:id", controllers.GetPostByIdController)
	eJwt.POST("api/v1/posts", controllers.AddPostController)
	eJwt.PUT("api/v1/posts/:id", controllers.UpdatePostController)
	eJwt.DELETE("api/v1/posts/:id", controllers.DeletePostController)
	return e
}