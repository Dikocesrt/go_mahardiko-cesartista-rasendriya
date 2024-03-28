package main

import (
	"gorm-soal-prioritas-2/config"
	"gorm-soal-prioritas-2/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	e.GET("/api/v1/fruits", controller.GetAllFruitsController)
	e.GET("/api/v1/fruits/:id", controller.GetFruitByIdController)
	e.POST("/api/v1/fruits", controller.AddFruitController)
	e.DELETE("/api/v1/fruits/:id", controller.DeleteFruitByIdController)
	e.Start(":8080")
}