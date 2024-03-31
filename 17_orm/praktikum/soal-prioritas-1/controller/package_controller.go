package controller

import (
	"gorm-soal-prioritas-1/config"
	"gorm-soal-prioritas-1/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeletePackageByIdController(c echo.Context) error {
	id := c.Param("id")

	var pkg model.Package

	pkg.Id, _ = strconv.Atoi(id)

	config.DB.Delete(&pkg)

	baseResponseSuccess := model.BaseResponseSuccess{
		Message: "Success Delete Package By Id",
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func UpdatePackageByIdController(c echo.Context) error {
	id := c.Param("id")

	var pkg model.Package
	c.Bind(&pkg)

	pkg.Id, _ = strconv.Atoi(id)

	config.DB.Save(&pkg)

	baseResponseSuccess := model.BaseResponseSuccess{
		Message: "Success Update Package By Id",
		Data:    pkg,
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func GetPackageByIdController(c echo.Context) error {
	id := c.Param("id")
	var pkg model.Package

	config.DB.First(&pkg, id)

	baseResponseSuccess := model.BaseResponseSuccess{
		Message: "Success Get Packages By Id",
		Data:    pkg,
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func GetPackagesController(c echo.Context) error {
	var pkgs []model.Package

	err := config.DB.Find(&pkgs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	baseResponseSuccess := model.BaseResponseSuccess{
		Message: "Success Get Packages",
		Data:    pkgs,
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func CreatePackageController(c echo.Context) error {
	var pkg model.Package

	c.Bind(&pkg)

	err := config.DB.Save(&pkg).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	baseResponseSuccess := model.BaseResponseSuccess{
		Message: "Success Create Package",
		Data:    pkg,
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}