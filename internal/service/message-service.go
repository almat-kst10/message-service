package service

import (
	"context"

	"github.com/almat-kst10/message-service/internal/models"
	repo "github.com/almat-kst10/message-service/internal/repository"
)

type IMessageClientService interface {
	SetMessage(ctx context.Context, message models.MessageClientRoom) error
	GetMessage(ctx context.Context, message models.MessageClientRoom) ([]*models.MessageClientRoom, error)
}

type MessageClientService struct {
	messageRepo repo.IMessageClientRepo
}

func NewMessageClientService(messageRepo repo.IMessageClientRepo) IMessageClientService {
	return &MessageClientService{
		messageRepo: messageRepo,
	}
}

func (s *MessageClientService) SetMessage(ctx context.Context, message models.MessageClientRoom) error {
	return s.messageRepo.SetMessage(ctx, message)
}

func (s *MessageClientService) GetMessage(ctx context.Context, message models.MessageClientRoom) ([]*models.MessageClientRoom, error) {
	return s.messageRepo.GetMessage(ctx, message)
}