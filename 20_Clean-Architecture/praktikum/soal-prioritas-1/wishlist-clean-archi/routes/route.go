package routes

import (
	"clean/constants"
	"clean/controllers/user"
	"clean/controllers/wishlist"

	myMiddleware "clean/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteController struct {
	wishlistController *wishlist.WishlistController
	userController *user.UserController
}

func NewRoute(wishlistController *wishlist.WishlistController, userController *user.UserController) *RouteController {
	return &RouteController{
		wishlistController: wishlistController,
		userController: userController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)
	e.POST("/register", r.userController.Register)
	e.POST("/login", r.userController.Login)
	eJwt := e.Group("/")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("wishlist", r.wishlistController.GetAll)
	eJwt.POST("wishlist", r.wishlistController.Create)
}