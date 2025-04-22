package grpc

import (
	"context"

	"github.com/almat-kst10/message-service/internal/service"
	"github.com/almat-kst10/message-service/proto"
)

type Server struct {
	// proto.UnimplementedMessageServiceServer
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

// func (s *Server) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
// 	message := &models.Message{
// 		SenderId:   int(req.SenderId),
// 		ReceiverId: int(req.ReceiverId),
// 		Text:       req.Text,
// 	}

// 	success, err := s.service.SaveMessage(ctx, message)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &proto.SendMessageResponse{Success: success}, nil
// }

// func (s *Server) GetMessage(ctx context.Context, req *proto.GetMessageRequest) (*proto.GetMessageResponse, error) {
// 	messages, err := s.service.GetMessage(ctx, int(req.User1Id), int(req.User2Id))
// 	if err != nil {
// 		return nil, err
// 	}

// 	var protoMessages []*proto.Message
// 	for _, msg := range messages {
// 		message := &proto.Message{
// 			Id:         int32(msg.Id),
// 			SenderId:   int32(msg.SenderId),
// 			ReceiverId: int32(msg.ReceiverId),
// 			Text:       msg.Text,
// 			Timestamp:  msg.Timestamp,
// 		}
// 		protoMessages = append(protoMessages, message)
// 	}

// 	return &proto.GetMessageResponse{Messages: protoMessages}, nil
// }
