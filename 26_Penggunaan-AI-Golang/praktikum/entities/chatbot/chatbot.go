package chatbot

type Chatbot struct {
	Role    string
	Budget  int
	Purpose string
	Answer  string
}

type UseCaseInterface interface {
	Chat(chat *Chatbot) (*Chatbot, error)
}