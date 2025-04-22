package grpc

import (
	"context"

	"github.com/almat-kst10/message-service/internal/service"
	"github.com/almat-kst10/message-service/proto"
)

type Server struct {
	proto.UnimplementedMessageServiceServer
	service service.IMessageService
}

func NewMessageHandler(service service.IMessageService) *Server {
	return &Server{service: service}
}

func (s *Server) RoomList(ctx context.Context, req *proto.RoomListRequest) (*proto.RoomListResponse, error) {
	roomList, err := s.service.RoomList(ctx, int(req.ProfileId))
	if err != nil {
		return nil, err
	}

	var protoRoomList []*proto.RoomGeneral
	for _, room := range roomList {
		protoRoom := &proto.RoomGeneral{
			RoomId: int32(room.RoomId),
			RoomTile: room.RoomTitle,
			ClientId: int32(room.ClientId),
			ProfileId: int32(room.ProfileId),
			ProfileName: room.ProfileName,
			ProfileSurname: room.ProfileSurname,
			RoleId: int32(room.RoleId),
			RoleName: room.RoleName,
			IsMuted: room.IsMuted,
			IsTyping: room.IsTyping,
		}

		protoRoomList = append(protoRoomList, protoRoom)
	}

	return &proto.RoomListResponse{RoomGeneral: protoRoomList}, nil
}