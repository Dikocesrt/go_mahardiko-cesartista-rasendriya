package chatbot

import (
	chatbotEntities "ai/entities/chatbot"
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/sashabaranov/go-openai"
)

type ChatbotUseCase struct {
}

func NewChatbotUseCase() *ChatbotUseCase {
	return &ChatbotUseCase{}
}

func getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo // see another model options: https://platform.openai.com/docs/models
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}

func (chatbotUseCase *ChatbotUseCase) Chat(chat *chatbotEntities.Chatbot) (*chatbotEntities.Chatbot, error) {
	if chat.Budget == 0 || chat.Purpose == "" {
		return nil, errors.New("chat content cannot be empty")
	}

	ctx := context.Background()

	client := openai.NewClient(os.Getenv("KEY"))

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a friendly chatbot.",
		},
	}
	model := openai.GPT3Dot5Turbo

	content := "find me a perfect laptop with this budget: " + strconv.Itoa(chat.Budget) + " and this purpose: " + chat.Purpose

	if chat.Budget != 0 || chat.Purpose != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		})
		chat.Budget = 0
		chat.Purpose = ""
	}

	resp, err := getCompletionFromMessages(ctx, client, messages, model)

	if err != nil {
		return nil, err
	}

	answer := openai.ChatCompletionMessage{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}

	chatAnswerEnt := chatbotEntities.Chatbot{
		Role:    answer.Role,
		Answer: answer.Content,
	}

	return &chatAnswerEnt, nil
}