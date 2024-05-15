package chatbot

import (
	"ai/controllers/chatbot/request"
	"ai/controllers/chatbot/response"
	chatbotEntities "ai/entities/chatbot"
	"net/http"

	"github.com/labstack/echo/v4"
	openai "github.com/sashabaranov/go-openai"
)

type ChatbotController struct {
	chatbotUseCase chatbotEntities.UseCaseInterface
}

func NewChatbotController(chatbotUseCase chatbotEntities.UseCaseInterface) *ChatbotController {
	return &ChatbotController{
		chatbotUseCase: chatbotUseCase,
	}
}

func (chatbotController *ChatbotController) Chat(c echo.Context) error {
	var chatbotRequest request.ChatbotRequest
	c.Bind(&chatbotRequest)

	chatbotEnt := chatbotEntities.Chatbot{
		Budget:  chatbotRequest.Budget,
		Purpose: chatbotRequest.Purpose,
	}

	chatbotEnt.Role = openai.ChatMessageRoleUser

	chatbot, err := chatbotController.chatbotUseCase.Chat(&chatbotEnt)

	chatbotResp := response.ChatbotResponse{
		Role:    chatbot.Role,
		Data: chatbot.Answer,
	}
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, chatbotResp)
}

