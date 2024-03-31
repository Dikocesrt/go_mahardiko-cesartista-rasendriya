package controllers

import (
	"middleware-eksplorasi/config"
	"middleware-eksplorasi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllCategoriesController(c echo.Context) error {
	var categories []models.Category

	err := config.DB.Find(&categories).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Get All Categories",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get All Categories",
		"data": categories,
	})
}

func AddCategoryController(c echo.Context) error {
	var category models.Category
	c.Bind(&category)

	err := config.DB.Create(&category).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Add Category",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Add Category",
	})
}

func DeleteCategoryByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var category models.Category
	category.Id = id

	err := config.DB.Delete(&category).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Delete Category",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Category",
	})
}