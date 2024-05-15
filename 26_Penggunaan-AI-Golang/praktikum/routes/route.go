package routes

import (
	"ai/controllers/chatbot"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	chatbotController *chatbot.ChatbotController
}

func NewRoute(chatbotController *chatbot.ChatbotController) *RouteController {
	return &RouteController{
		chatbotController:  chatbotController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	e.POST("/chat", r.chatbotController.Chat)
}