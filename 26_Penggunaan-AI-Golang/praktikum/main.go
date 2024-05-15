package main

import (
	chatbotController "ai/controllers/chatbot"
	"ai/routes"
	chatbotUseCase "ai/usecases/chatbot"

	"github.com/labstack/echo/v4"
)

func main() {
	chatbotUC := chatbotUseCase.NewChatbotUseCase()
	chatbotCont := chatbotController.NewChatbotController(chatbotUC)

	route := routes.NewRoute(chatbotCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}