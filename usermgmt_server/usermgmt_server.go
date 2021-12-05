package main

import (
	"context"
	pb "CourseProject/usermgmt/usermgmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		userList:&pb.UserList{},
	}
}
type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	userList *pb.UserList
}

func (server *UserManagementServer) Run() error{
	lis,err :=net.Listen("tcp",port)
	if err!=nil{
		log.Fatalf("failed to listen %v",err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Recieved: %v", in.GetName())
	var userId int32 = int32(rand.Intn(1000));
	createdUser := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: userId}
	s.userList.Users = append(s.userList.Users, createdUser)
	return createdUser, nil
}

func(s *UserManagementServer) GetUsers(ctx context.Context, params *pb.GetUsersParams)(*pb.UserList, error){
	return s.userList, nil
}

func main() {
	var userManagementServer *UserManagementServer = NewUserManagementServer()
	if err := userManagementServer.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}