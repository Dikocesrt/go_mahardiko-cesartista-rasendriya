package controllers

import (
	"middleware-prioritas-2/config"
	"middleware-prioritas-2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllPostsController(c echo.Context) error {
	var posts []models.Post

	err := config.DB.Find(&posts).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Get All Posts",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get All Posts",
		"data": posts,
	})
}

func GetPostByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var post models.Post
	post.Id = id

	err := config.DB.First(&post).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Get Post By Id",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Post By Id",
		"data": post,
	})
}

func AddPostController(c echo.Context) error {
	var post models.Post
	c.Bind(&post)

	err := config.DB.Create(&post).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Add Post",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Add Post",
	})
}

func UpdatePostController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var post models.Post
	c.Bind(&post)
	post.Id = id

	err := config.DB.Save(&post).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Update Post",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update Post",
	})
}

func DeletePostController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var post models.Post
	post.Id = id

	err := config.DB.Delete(&post).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Delete Post",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Post",
	})
}