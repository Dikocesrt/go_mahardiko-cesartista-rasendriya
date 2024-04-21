package fruit

import (
	fruitRequest "clean2/controllers/fruit/request"
	fruitEntities "clean2/entities/fruit"
	"clean2/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FruitController struct {
	fruitUseCase fruitEntities.UseCaseInterface
}

func NewFruitController(fruitUseCase fruitEntities.UseCaseInterface) *FruitController {
	return &FruitController{
		fruitUseCase: fruitUseCase,
	}
}

func (fruitController *FruitController) GetAllFruits(c echo.Context) error {
	var fruitsEntities []fruitEntities.Fruit
	fruitsEntities, err := fruitController.fruitUseCase.GetAllFruits()
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Fruits", fruitsEntities))
}

func (fruitController *FruitController) GetFruitById(c echo.Context) error {
	id := c.Param("id")

	fruitEnt, err := fruitController.fruitUseCase.GetFruitById(id)

	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Fruit By Id", fruitEnt))
}

func (fruitController *FruitController) AddFruit(c echo.Context) error {
	var fruitReq fruitRequest.Fruit
	c.Bind(&fruitReq)

	fruitEnt := fruitEntities.Fruit{
		Name:  fruitReq.Name,
		Price: fruitReq.Price,
	}

	fruitEnt, err := fruitController.fruitUseCase.AddFruit(fruitEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Add Fruit", fruitEnt))
}

func (fruitController *FruitController) DeleteFruitById(c echo.Context) error {
	id := c.Param("id")

	err := fruitController.fruitUseCase.DeleteFruitById(id)
	
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Fruit", nil))
}