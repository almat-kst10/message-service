package service

import (
	"context"

	"github.com/almat-kst10/message-service/internal/models"
	"github.com/almat-kst10/message-service/internal/repository"
)

type IMessageService interface {
	RoomList(ctx context.Context, profiles_id int) ([]models.RoomGeneralInfo, error)
}

type MessageService struct {
	repo repository.IMessageRepo
}

func NewServiceMessage(repo repository.IMessageRepo) IMessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) RoomList(ctx context.Context, profiles_id int) ([]models.RoomGeneralInfo, error) {
	return s.repo.RoomList(ctx, profiles_id)
}