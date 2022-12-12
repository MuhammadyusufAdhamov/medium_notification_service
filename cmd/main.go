package main

import (
	"github.com/MuhammadyusufAdhamov/medium_notification_service/config"
	pb "github.com/MuhammadyusufAdhamov/medium_notification_service/genproto/notification_service"
	"github.com/MuhammadyusufAdhamov/medium_notification_service/service"
	"google.golang.org/grpc"
	"log"
	"net"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")

	notificationService := service.NewNotificationService(&cfg)

	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		log.Fatalf("failed to server error: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, notificationService)

	log.Println("Grpc server started in port ", cfg.GrpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while listening: %v", err)
	}
}
