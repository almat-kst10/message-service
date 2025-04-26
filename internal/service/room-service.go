package service

import (
	"context"
	"fmt"

	"github.com/almat-kst10/message-service/internal/models"
	repo "github.com/almat-kst10/message-service/internal/repository"
)

type IRoomService interface {
	RoomList(ctx context.Context, profilesId int) ([]*models.RoomGeneralInfo, error)
	RoomCreate(ctx context.Context, roomTitle string, firstProfileId int, secondProfileId int) (int, error)
	RoomDelete(ctx context.Context, roomId int) error
}

type RoomService struct {
	txRepo     repo.ITxRepo
	roomRepo   repo.IRoomRepo
	clientRepo repo.IClientRoomRepo
}

func NewRoomService(txRepo repo.ITxRepo, roomRepo repo.IRoomRepo, clientRepo repo.IClientRoomRepo) IRoomService {
	return &RoomService{
		txRepo:     txRepo,
		roomRepo:   roomRepo,
		clientRepo: clientRepo,
	}
}

var roomRoleDefault = map[string]int{
	"Owner": 1,
	"Admin": 2,
	"Member": 3,
}

func (s *RoomService) RoomCreate(ctx context.Context, roomTitle string, firstProfileId int, secondProfileId int) (int, error) {
	tx, err := s.txRepo.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed transaction: %s", err)
	}
	defer tx.Rollback()

	roomId, err := s.roomRepo.RoomCreate(ctx, roomTitle)
	if err != nil {
		return 0, err
	}

	client := &models.ClientRoom{
		RoomId: roomId,
		ProfileId: firstProfileId,
		RoleId: roomRoleDefault["Member"],
	}
	err = s.clientRepo.ClientJoin(ctx, client)
	if err != nil {
		return 0, err
	}

	if secondProfileId > 0 {
		client.ProfileId = secondProfileId
		err = s.clientRepo.ClientJoin(ctx, client)
		if err != nil {
			return 0, err
		}
	}

	tx.Commit()
	return roomId, nil
}

func (s *RoomService) RoomList(ctx context.Context, profilesId int) ([]*models.RoomGeneralInfo, error) {
	return s.roomRepo.RoomList(ctx, profilesId)
}

func (s *RoomService) RoomDelete(ctx context.Context, roomId int) error {
	return s.roomRepo.RoomDelete(ctx, roomId)
}
