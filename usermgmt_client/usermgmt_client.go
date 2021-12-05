package main

import (
	pb "CourseProject/usermgmt/usermgmt"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err!=nil{
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newUsers = make(map[string]int32)
	newUsers["Andrew"] = 34
	newUsers["Ann"] = 33
	for name, age := range newUsers {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err!=nil{
			log.Fatalf("could not create User: %v", err)
		}
		log.Printf("User Details:" +
			" Name: %s" +
			" Age: %d " +
			"Id: %d",
			r.GetName(), r.GetAge(), r.GetId())
	}
	params := &pb.GetUsersParams{}
	r,err := c.GetUsers(ctx, params)
	if err != nil {
		log.Fatalf("Couldn't recieve users: %v", err)
	}
	log.Print("\nUserList: \n")
	log.Printf("r.GetUsers() %v", r.GetUsers())
}