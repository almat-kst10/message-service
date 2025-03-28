package app

import (
	"fmt"
	"net"

	"github.com/almat-kst10/message-service/configs"
	"github.com/almat-kst10/message-service/internal/repository"
	"github.com/almat-kst10/message-service/internal/service"
	"github.com/almat-kst10/message-service/proto"
	"google.golang.org/grpc"

	msgGRPC "github.com/almat-kst10/message-service/internal/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(configs *configs.Configs) error {
	repo, err := repository.NewRepositoryMessage(configs)
	if err != nil {
		return err
	}
	defer repo.Close()

	service := service.NewServiceMessage(repo)
	grpcServer := msgGRPC.NewMessageHandler(service)

	grpcPort := fmt.Sprintf(":%s", configs.GRPC.Port)
	lis, err := net.Listen(configs.GRPC.ConnectType, grpcPort)
	if err != nil {
		return err
	}
	serv := grpc.NewServer()

	proto.RegisterMessageServiceServer(serv, grpcServer)
	reflection.Register(serv)
	return serv.Serve(lis)

	// redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	// wsServer := websocket.NewWebSocket(redisClient)

	// http.HandleFunc("/ws", wsServer.HandlerConn)

}
