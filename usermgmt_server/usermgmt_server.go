package main

import (
	"context"
	pb "CourseProject/usermgmt/usermgmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"os"
	"io/ioutil"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{

	}
}
type UserManagementServer struct {
	pb.UnimplementedUserManagementServer

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
	readBytes, err := ioutil.ReadFile("users.json")
	var usersList *pb.UsersList = &pb.UsersList{}
	var userId int32 = int32(rand.Intn(1000));
	createdUser := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: userId}

	if err!=nil{
		if os.IsNotExist(err){
			log.Print("File not found! Creating new file!")
			usersList.Users = append(usersList.Users, createdUser)
			jsonBytes,err := protojson.Marshal(usersList)
			if err!=nil{
				log.Fatalf("JSON Marsharing failed %v", err)
			}
			if err := ioutil.WriteFile("users.json", jsonBytes, 0664); err!=nil{
				log.Fatalf("Failed write to file %v", err)
			}
			return createdUser, nil
		} else {
			log.Fatalln("Error reading file ", err)
		}

	}

	if err := protojson.Unmarshal(readBytes, usersList);err!=nil{
			log.Fatalf("Failed to parseuser list %v", err)
	}
	usersList.Users = append(usersList.Users, createdUser)
	jsonBytes,err := protojson.Marshal(usersList)
	if err!=nil{
		log.Fatalf("JSON Marsharing failed %v", err)
	}
	if err := ioutil.WriteFile("users.json", jsonBytes, 0664); err!=nil{
		log.Fatalf("Failed write to file %v", err)
	}

	return createdUser, nil
}

func(s *UserManagementServer) GetUsers(ctx context.Context, params *pb.GetUsersParams)(*pb.UsersList, error){
	jsonBytes, err := ioutil.ReadFile("users.json")
	if err!=nil{
		log.Fatalf("Failed read from file % v", err)
	}
	var usersList *pb.UsersList = &pb.UsersList{}
	if err := protojson.Unmarshal(jsonBytes, usersList); err != nil {
		log.Fatalf("Unmarshaling failed %v", err)
	}

	return usersList, nil
}

func main() {
	var userManagementServer *UserManagementServer = NewUserManagementServer()
	if err := userManagementServer.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}