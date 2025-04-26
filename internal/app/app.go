package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/almat-kst10/message-service/configs"
	"github.com/almat-kst10/message-service/internal/repository"
	"github.com/almat-kst10/message-service/internal/service"
	"github.com/almat-kst10/message-service/pkg/db"
	"github.com/almat-kst10/message-service/proto"
	"google.golang.org/grpc"

	msgGRPC "github.com/almat-kst10/message-service/internal/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(configs *configs.Configs) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	// defer cancel()

	postgresClient, err := db.NewPostgresClient(configs)
	if err != nil {
		return err
	}
	defer postgresClient.Close()


	roomRepo := repo.NewRoomRepo(postgresClient)
	roomService := service.NewRoomService(roomRepo)
	
	clientRoomRepo := repo.NewClientRoomRepo(postgresClient)
	clientRoomService := service.NewClientRoomService(clientRoomRepo)

	messageClientRepo := repo.NewMessageClientRepo(postgresClient)
	messageClientService := service.NewMessageClientService(messageClientRepo)
	
	grpcServer := msgGRPC.NewMessageHandler(roomService, clientRoomService, messageClientService)

	grpcPort := fmt.Sprintf(":%s", configs.GRPC.Port)
	lis, err := net.Listen(configs.GRPC.ConnectType, grpcPort)
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	proto.RegisterMessageServiceServer(serv, grpcServer)
	reflection.Register(serv)

	// Канал для graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Запускаем сервер в горутине
	go func() {
		if err := serv.Serve(lis); err != nil {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	// Ожидание сигнала завершения
	<-stop
	log.Println("\n Получен сигнал завершения. Завершаем сервер...")

	// Завершаем сервер корректно
	serv.GracefulStop()

	return nil
}
