package controller

import (
	"gorm-eksplorasi/config"
	"gorm-eksplorasi/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddCommentController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("post_id"))
	
	var comment model.Comment
	
	c.Bind(&comment)

	comment.PostinganId = id

	err := config.DB.Create(&comment).Error
	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Failed Create Comment",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	baseResponseSuccess := model.BaseResponseSuccess {
		Message: "Success Create Comment",
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func UpdateCommentController(c echo.Context) error {
	id := c.Param("id")

	var comment model.Comment
	c.Bind(&comment)

	err := config.DB.Model(&model.Comment{}).Where("id = ?", id).Update("isi_komentar", comment.IsiKomentar).Error
	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Failed Update Comment",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}

	baseResponseSuccess := model.BaseResponseSuccess {
		Message: "Success Update Comment",
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func DeleteCommentController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var comment model.Comment
	comment.Id = id
	err := config.DB.Delete(&comment).Error
	if err != nil {
		baseResponseError := model.BaseResponseError {
			Message: "Failed Delete Comment",
		}
		return c.JSON(http.StatusInternalServerError, baseResponseError)
	}
	baseResponseSuccess := model.BaseResponseSuccess {
		Message: "Success Delete Comment",
	}
	return c.JSON(http.StatusOK, baseResponseSuccess)
}