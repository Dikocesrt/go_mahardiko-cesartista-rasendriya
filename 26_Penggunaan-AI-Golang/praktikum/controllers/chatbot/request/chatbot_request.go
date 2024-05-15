package request

type ChatbotRequest struct {
	Budget  int    `json:"budget"`
	Purpose string `json:"purpose"`
}