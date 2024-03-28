package main

import (
	"gorm-eksplorasi/config"
	"gorm-eksplorasi/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	e.GET("/api/v1/posts", controller.GetPostsController)
	e.GET("/api/v1/posts/:id", controller.GetPostByIdController)
	e.POST("/api/v1/posts", controller.AddPostController)
	e.PUT("/api/v1/posts/:id", controller.UpdatePostController)
	e.DELETE("/api/v1/posts/:id", controller.DeletePostController)
	e.POST("/api/v1/comments/:post_id", controller.AddCommentController)
	e.PUT("/api/v1/comments/:id", controller.UpdateCommentController)
	e.DELETE("/api/v1/comments/:id", controller.DeleteCommentController)
	e.Start(":8080")
}