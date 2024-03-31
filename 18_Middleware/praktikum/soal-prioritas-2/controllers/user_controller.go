package controllers

import (
	"middleware-prioritas-2/config"
	"middleware-prioritas-2/middleware"
	"middleware-prioritas-2/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	var user models.User

	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login User",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.Id, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login User",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{
		Id:    user.Id,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login User",
		"data":    userResponse,
	})
}

func RegisterUserController(c echo.Context) error {
	var user models.User

	c.Bind(&user)

	err := config.DB.Create(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Register User",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Register User",
	})
}