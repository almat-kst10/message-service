package service

import (
	"context"

	"github.com/almat-kst10/message-service/internal/models"
	"github.com/almat-kst10/message-service/internal/repository"
)

type IRoomService interface {
	RoomList(ctx context.Context, profiles_id int) ([]*models.RoomGeneralInfo, error)
	RoomCreate(ctx context.Context, roomTitle string) (int, error)
	RoomDelete(ctx context.Context, roomId int) error
}

type RoomService struct {
	repo repo.IRoomRepo
}

func NewRoomService(repo repo.IRoomRepo) IRoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) RoomCreate(ctx context.Context, roomTitle string) (int, error) {
	return s.repo.RoomCreate(ctx, roomTitle)
}

func (s *RoomService) RoomList(ctx context.Context, profiles_id int) ([]*models.RoomGeneralInfo, error) {
	return s.repo.RoomList(ctx, profiles_id)
}

func (s *RoomService) RoomDelete(ctx context.Context, roomId int) error {
	return s.repo.RoomDelete(ctx, roomId)
}
