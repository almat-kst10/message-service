package service

import (
	"context"

	"github.com/almat-kst10/message-service/internal/models"
	"github.com/almat-kst10/message-service/internal/repository"
)

type IMessageService interface {
	SaveMessage(ctx context.Context, message *models.Message) (bool, error)
	GetMessage(ctx context.Context, user1Id, user2Id int) ([]models.Message, error)
	ChatsList(ctx context.Context, profiles_id int) ([]models.Chat, error)
}

type MessageService struct {
	repo repository.IMessageRepo
}

func NewServiceMessage(repo repository.IMessageRepo) IMessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) SaveMessage(ctx context.Context, message *models.Message) (bool, error) {
	return s.repo.SaveMessage(ctx, message)
}

func (s *MessageService) GetMessage(ctx context.Context, user1Id, user2Id int) ([]models.Message, error) {
	return s.repo.GetMessage(ctx, user1Id, user2Id)
}

func (s *MessageService) ChatsList(ctx context.Context, profiles_id int) ([]models.Chat, error) {
	return s.repo.ChatsList(ctx, )
}