package routes

import (
	"middleware-eksplorasi/constants"
	"middleware-eksplorasi/controllers"
	myMiddleware "middleware-eksplorasi/middleware"

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
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("api/v1/posts", controllers.GetAllPostsController)
	eJwt.GET("api/v1/posts/:id", controllers.GetPostByIdController)
	eJwt.POST("api/v1/posts", controllers.AddPostController)
	eJwt.PUT("api/v1/posts/:id", controllers.UpdatePostController)
	eJwt.DELETE("api/v1/posts/:id", controllers.DeletePostController)

	adminGroup := e.Group("/")
	adminGroup.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	adminGroup.Use(myMiddleware.AdminOnlyMiddleware)
	adminGroup.GET("api/v1/categories", controllers.GetAllCategoriesController)
	adminGroup.POST("api/v1/categories", controllers.AddCategoryController)
	adminGroup.DELETE("api/v1/categories/:id", controllers.DeleteCategoryByIdController)
	return e
}