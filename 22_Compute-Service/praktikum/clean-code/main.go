package main

import (
	"clean/configs"
	userController "clean/controllers/user"
	"clean/repositories/mysql"
	userRepositories "clean/repositories/mysql/user"
	"clean/routes"
	userUseCase "clean/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())

	userRepo := userRepositories.NewUserRepo(db)
	userUC := userUseCase.NewUserUseCase(userRepo)
	userCont := userController.NewUserController(userUC)

	routes := routes.NewRoute(userCont)

	e := echo.New()
	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}