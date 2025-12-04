package chat

import (
	"postgres/internal/repository"
	"postgres/internal/service"
)

type chatService struct {
	chatRepository repository.ChatRepository
}

func NewChatService(chatRepository repository.ChatRepository) service.ChatService {
	return &chatService{chatRepository: chatRepository}
}
