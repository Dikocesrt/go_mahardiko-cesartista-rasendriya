package main

import (
	"clean/configs"
	userController "clean/controllers/user"
	wishlistController "clean/controllers/wishlist"
	"clean/repositories/mysql"
	userRepositories "clean/repositories/mysql/user"
	wishlistRepositories "clean/repositories/mysql/wishlist"
	"clean/routes"
	userUseCase "clean/usecases/user"
	wishlistUseCase "clean/usecases/wishlist"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())

	wishlistRepo := wishlistRepositories.NewWishlistRepo(db)
	wishlistUC := wishlistUseCase.NewWishlistUseCase(wishlistRepo)
	wishlistCont := wishlistController.NewWishlistController(wishlistUC)

	userRepo := userRepositories.NewUserRepo(db)
	userUC := userUseCase.NewUserUseCase(userRepo)
	userCont := userController.NewUserController(userUC)

	route := routes.NewRoute(wishlistCont, userCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}