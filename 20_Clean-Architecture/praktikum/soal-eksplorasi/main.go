package main

import (
	"clean2/configs/mysql"
	fruitController "clean2/controllers/fruit"
	fruitRepositories "clean2/repositories/fruit"
	"clean2/routes"
	fruitUseCase "clean2/usecases/fruit"

	"github.com/labstack/echo/v4"
)

func main() {
	mysql.LoadEnv()
	db := mysql.ConnectDB(mysql.InitConfigMySQL())

	fruitRepo := fruitRepositories.NewFruitRepo(db)
	fruitUC := fruitUseCase.NewFruitUseCase(fruitRepo)
	fruitCont := fruitController.NewFruitController(fruitUC)

	route := routes.NewRoute(fruitCont)
	
	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}