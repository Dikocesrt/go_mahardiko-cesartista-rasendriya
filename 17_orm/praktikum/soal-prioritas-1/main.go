package main

import (
	"gorm-soal-prioritas-1/config"
	"gorm-soal-prioritas-1/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	e.GET("/api/v1/packages", controller.GetPackagesController)
	e.GET("/api/v1/packages/:id", controller.GetPackageByIdController)
	e.POST("/api/v1/packages", controller.CreatePackageController)
	e.PUT("/api/v1/packages/:id", controller.UpdatePackageByIdController)
	e.DELETE("/api/v1/packages/:id", controller.DeletePackageByIdController)
	e.Start(":8080")
}