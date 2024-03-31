package controllers

import (
	"middleware-prioritas-1/config"
	"middleware-prioritas-1/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllPackagesController(c echo.Context) error {
	var pkgs []models.Package

	err := config.DB.Find(&pkgs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Get All Packages",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get All Packages",
		"data":    pkgs,
	})
}

func GetPackageByIdController(c echo.Context) error {
	id := c.Param("id")

	var pkg models.Package

	err := config.DB.First(&pkg, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Get Package By Id",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Package By Id",
		"data":    pkg,
	})
}

func AddPackageController(c echo.Context) error {
	var pkg models.Package
	c.Bind(&pkg)

	err := config.DB.Create(&pkg).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Add Package",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Add Package",
	})
}

func UpdatePackageByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var pkg models.Package
	c.Bind(&pkg)
	pkg.Id = id

	err := config.DB.Save(&pkg).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Update Package",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update Package",
	})
}

func DeletePackageByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var pkg models.Package
	pkg.Id = id

	err := config.DB.Delete(&pkg).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Delete Package",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Package",
	})
}