package controller

import (
	"gorm-eksplorasi/config"
	"gorm-eksplorasi/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPostsController(c echo.Context) error {
	var posts []model.Post
	var comments []model.Comment

	err := config.DB.Find(&posts).Error
	if(err != nil) {
		baseResponseError := model.BaseResponseError {
			Message: "Post not found",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	config.DB.Find(&comments)

	for i:=0;i<len(posts);i++{
		for j:=0;j<len(comments);j++{
			if posts[i].Id == comments[j].PostinganId {
				posts[i].Comments = append(posts[i].Comments, comments[j])
			}
		}
	}

	baseResponseSuccess := model.BaseResponseSuccess {
		Message: "Success Get All Posts",
		Data: posts,
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func GetPostByIdController(c echo.Context) error {
	id := c.Param("id")

	var post model.Post
	err := config.DB.First(&post, id).Error
	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Post not found",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	var comments []model.Comment
	config.DB.Where("postingan_id = ?", post.Id).Find(&comments)

	post.Comments = append(post.Comments, comments...)

	baseResponseSuccess := model.BaseResponseSuccess {
		Message: "Success Get Post By Id",
		Data: post,
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func AddPostController(c echo.Context) error {
	var post model.Post
	c.Bind(&post)

	// fmt.Println(post)

	err := config.DB.Create(&post).Error
	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Failed Create Post",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	baseResponseSuccess := model.BaseResponseSuccess {
		Message: "Success Create Post",
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func UpdatePostController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var post model.Post
	c.Bind(&post)

	post.Id = id
	err := config.DB.Save(&post).Error
	
	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Failed Update Post",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}
	
	baseResponseSuccess := model.BaseResponseSuccess {
		Message: "Success Update Post",
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func DeletePostController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var post model.Post
	post.Id = id
	err := config.DB.Delete(&post).Error
	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Failed Delete Post",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}
	baseResponseSuccess := model.BaseResponseSuccess {
		Message: "Success Delete Post",
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}