package routes

import (
	"middleware-prioritas-1/constant"
	"middleware-prioritas-1/controllers"
	myMiddleware "middleware-prioritas-1/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	myMiddleware.LogMiddleware(e)
	e.POST("/api/v1/register", controllers.RegisterUserController)
	e.POST("/api/v1/login", controllers.LoginUserController)
	eJwt := e.Group("/")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.GET("api/v1/packages", controllers.GetAllPackagesController)
	eJwt.GET("api/v1/packages/:id", controllers.GetPackageByIdController)
	eJwt.POST("api/v1/packages", controllers.AddPackageController)
	eJwt.PUT("api/v1/packages/:id", controllers.UpdatePackageByIdController)
	eJwt.DELETE("api/v1/packages/:id", controllers.DeletePackageByIdController)
	return e
}