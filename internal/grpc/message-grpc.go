package grpc

import (
	"context"

	"github.com/almat-kst10/message-service/internal/models"
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

func (s *Server) ChatsList(ctx context.Context, req *proto.ChatsListRequest) (*proto.ChatsListResponse, error) {
	chatsList, err := s.service.ChatsList(ctx, int(req.ProfileId))
	if err != nil {
		return nil, err
	}

	var protoChatsList []*proto.Chat
	for _, chat := range chatsList {
		protoChats := &proto.Chat{
			MyProfileId:    int32(chat.MyProfileId),
			User2ProfileId: int32(chat.User2ProfileId),
			LastMessage:    chat.LastMessage,
			IsRead:         chat.IsRead,
			CountNewMsg:    int32(chat.CountNewMsg),
			IsVisible:      chat.IsVisible,
		}
		protoChatsList = append(protoChatsList, protoChats)
	}

	return &proto.ChatsListResponse{Chats: protoChatsList}, nil
}

func (s *Server) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	message := &models.Message{
		SenderId:   int(req.SenderId),
		ReceiverId: int(req.ReceiverId),
		Text:       req.Text,
	}

	success, err := s.service.SaveMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return &proto.SendMessageResponse{Success: success}, nil
}

func (s *Server) GetMessage(ctx context.Context, req *proto.GetMessageRequest) (*proto.GetMessageResponse, error) {
	messages, err := s.service.GetMessage(ctx, int(req.User1Id), int(req.User2Id))
	if err != nil {
		return nil, err
	}

	var protoMessages []*proto.Message
	for _, msg := range messages {
		message := &proto.Message{
			Id:         int32(msg.Id),
			SenderId:   int32(msg.SenderId),
			ReceiverId: int32(msg.ReceiverId),
			Text:       msg.Text,
			Timestamp:  msg.Timestamp,
		}
		protoMessages = append(protoMessages, message)
	}

	return &proto.GetMessageResponse{Messages: protoMessages}, nil
}
