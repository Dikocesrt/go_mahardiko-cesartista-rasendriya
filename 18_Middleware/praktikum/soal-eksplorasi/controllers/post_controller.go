package controllers

import (
	"middleware-eksplorasi/config"
	"middleware-eksplorasi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllPostsController(c echo.Context) error {
	var posts []models.Post
	var categories []models.Category

	err := config.DB.Find(&posts).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Get All Posts",
			"error": err.Error(),
		})
	}

	err = config.DB.Find(&categories).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Get All Categories",
			"error": err.Error(),
		})
	}

	postResponses := make([]models.PostResponse, len(posts))

	for i:=0;i<len(posts);i++ {
		for _, category := range categories {
			if posts[i].CategoryId == category.Id {
				posts[i].Category = category
				break
			}
		}
		postResponses[i].Id = posts[i].Id
		postResponses[i].Judul = posts[i].Judul
		postResponses[i].Konten = posts[i].Konten
		postResponses[i].Category = posts[i].Category
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get All Posts",
		"data": postResponses,
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

	var category models.Category
	err = config.DB.Where("id = ?", post.CategoryId).First(&category).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Get Category By Id",
			"error": err.Error(),
		})
	}

	post.Category = category

	var postResponse models.PostResponse
	postResponse.Id = post.Id
	postResponse.Judul = post.Judul
	postResponse.Konten = post.Konten
	postResponse.Category = post.Category

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Post By Id",
		"data": postResponse,
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