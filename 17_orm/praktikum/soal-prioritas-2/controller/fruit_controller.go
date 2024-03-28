package controller

import (
	"encoding/json"
	"gorm-soal-prioritas-2/config"
	"gorm-soal-prioritas-2/helper"
	"gorm-soal-prioritas-2/model"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func DeleteFruitByIdController(c echo.Context) error {
	id := c.Param("id")
	err := config.DB.Delete(&model.Fruit{}, id).Error

	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Error Query Get Fruit By Id",
			Data: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Fruit By Id",
	})
}

func GetFruitByIdController(c echo.Context) error {
	id := c.Param("id")

	var fruit model.Fruit

	err := config.DB.First(&fruit, id).Error;
	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Error Query Get Fruit By Id",
			Data: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	baseResponseSuccess := model.BaseResponseSuccess {
		Message: "Success Get Fruit By Id",
		Data: helper.ConvertFruitToJSONFormat(&fruit),
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func AddFruitController(c echo.Context) error {
	var fruit model.Fruit
	c.Bind(&fruit)

	var url = "https://www.fruityvice.com/api/fruit/" + fruit.Nama

	res, _ := http.Get(url)
	resData, _ := io.ReadAll(res.Body)

	var data map[string]interface{}

	_ = json.Unmarshal(resData, &data)

	fruit.Kalori = data["nutritions"].(map[string]interface{})["calories"].(float64)
	fruit.Lemak = data["nutritions"].(map[string]interface{})["fat"].(float64)
	fruit.Gula = data["nutritions"].(map[string]interface{})["sugar"].(float64)
	fruit.Karbohidrat = data["nutritions"].(map[string]interface{})["carbohydrates"].(float64)
	fruit.Protein = data["nutritions"].(map[string]interface{})["protein"].(float64)

	id := int(uuid.New().ID())

	fruit.ID = id

	err := config.DB.Save(&fruit).Error
	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Error Add Fruit",
			Data: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Add Fruit",
	})
}

func GetAllFruitsController(c echo.Context) error {
	var fruits []model.Fruit

	err := config.DB.Find(&fruits).Error;

	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Error Query Get All Fruits",
			Data: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	var fruitsJSON []interface{}
	for _, fruit := range fruits {
		fruitsJSON = append(fruitsJSON, helper.ConvertFruitToJSONFormat(&fruit))
	}

	baseResponseSuccess := model.BaseResponseSuccess {
		Message: "Success Get All Fruits",
		Data: fruitsJSON,
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}