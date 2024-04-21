package routes

import (
	"clean2/controllers/fruit"
	myMiddleware "clean2/middleware"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	fruitController *fruit.FruitController
}

func NewRoute(fruitController *fruit.FruitController) *RouteController {
	return &RouteController{
		fruitController: fruitController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)

	fruitRoutes := e.Group("/api/v1")

	fruitRoutes.GET("/fruits", r.fruitController.GetAllFruits)
	fruitRoutes.GET("/fruits/:id", r.fruitController.GetFruitById)
	fruitRoutes.POST("/fruits", r.fruitController.AddFruit)
	fruitRoutes.DELETE("/fruits/:id", r.fruitController.DeleteFruitById)
}