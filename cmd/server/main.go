package main

import (
	api "CourseProject/pkg/api"
	"CourseProject/pkg/order"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	srv := &order.GRPCServer{}
	api.RegisterOrderingServer(grpcServer, srv)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Order: Ошибка прослушивания: %v", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Order: Ошибка запуска сервера: %v", err)
	}

}
