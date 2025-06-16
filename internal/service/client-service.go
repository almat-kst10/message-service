package service

import (
	"context"

	"github.com/almat-kst10/message-service/internal/models"
	repo "github.com/almat-kst10/message-service/internal/repository"
)

type IClientRoomService interface {
	ClientJoin(ctx context.Context, client *models.ClientRoom) error
	ClientDelete(ctx context.Context, client *models.ClientRoom) error
	ClientEdit(ctx context.Context, client *models.ClientRoom) error
}

type ClientRoomService struct {
	clientRoomRepo repo.IClientRoomRepo
}

func NewClientRoomService(clientRoomRepo repo.IClientRoomRepo) IClientRoomService {
	return &ClientRoomService{
		clientRoomRepo: clientRoomRepo,
	}
}

func (s *ClientRoomService) ClientJoin(ctx context.Context, client *models.ClientRoom) error {
	return s.clientRoomRepo.ClientJoin(ctx, client)
}

func (s *ClientRoomService) ClientDelete(ctx context.Context, client *models.ClientRoom) error {
	return s.clientRoomRepo.ClientDelete(ctx, client)
}

func (s *ClientRoomService) ClientEdit(ctx context.Context, client *models.ClientRoom) error {
	return s.clientRoomRepo.ClientEdit(ctx, client)
}
