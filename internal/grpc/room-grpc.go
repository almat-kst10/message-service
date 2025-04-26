package grpc

import (
	"context"

	"github.com/almat-kst10/message-service/internal/service"
	"github.com/almat-kst10/message-service/proto"
)

type Server struct {
	proto.UnimplementedMessageServiceServer
	roomService       service.IRoomService
	roomClientService service.IClientRoomService
	messageClient     service.IMessageClientService
}

func NewMessageHandler(roomService service.IRoomService, roomClientService service.IClientRoomService, messageClient service.IMessageClientService) *Server {
	return &Server{
		roomService:       roomService,
		roomClientService: roomClientService,
		messageClient:     messageClient,
	}
}

func (s *Server) RoomList(ctx context.Context, req *proto.RoomListRequest) (*proto.RoomListResponse, error) {
	roomList, err := s.roomService.RoomList(ctx, int(req.ProfileId))
	if err != nil {
		return nil, err
	}

	var protoRoomList []*proto.RoomGeneral
	for _, room := range roomList {
		protoRoom := &proto.RoomGeneral{
			RoomId:         int32(room.RoomId),
			RoomTile:       room.RoomTitle,
			ClientId:       int32(room.ClientId),
			ProfileId:      int32(room.ProfileId),
			RoleId:         int32(room.RoleId),
			RoleName:       room.RoleName,
			IsMuted:        room.IsMuted,
			IsTyping:       room.IsTyping,
		}

		protoRoomList = append(protoRoomList, protoRoom)
	}

	return &proto.RoomListResponse{RoomGeneral: protoRoomList}, nil
}


func (s *Server) RoomCreateGroup(ctx context.Context, req *proto.RoomCreateGroupRequest) (*proto.RoomCreateResponse, error) {
	roomId, err := s.roomService.RoomCreate(ctx, req.RoomTitle, int(req.ProfileId), 0)
	if err != nil {
		return nil, err
	}

	return &proto.RoomCreateResponse{RoomId: int32(roomId)}, nil
}


func (s *Server) RoomCreatePerson(ctx context.Context, req *proto.RoomCreatePersonRequest) (*proto.RoomCreateResponse, error) {
	roomId, err := s.roomService.RoomCreate(ctx, req.RoomTitle, int(req.FirstProfileId), int(req.SecondProfileId))
	if err != nil {
		return nil, err
	}

	return &proto.RoomCreateResponse{RoomId: int32(roomId)}, nil
}