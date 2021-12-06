package main

import (
	pb "CourseProject/usermgmt/usermgmt"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}

type UserManagementServer struct {
	conn *pgx.Conn
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

func (server *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Recieved: %v", in.GetName())

	createSql := `
		CREATE TABLE if not exists users(
			id SERIAL PRIMARY KEY,
			name text,
			age int
		);`
	_, err := server.conn.Exec(context.Background(), createSql)
	if err!= nil{
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}
	createdUser := &pb.User{Name: in.GetName(), Age: in.GetAge()}
	tx, err := server.conn.Begin(context.Background())
	if err!=nil{
		log.Fatalf("conn.Begin failed: %v", err)
	}
	_, err = tx.Exec(context.Background(), "INSERT INTO users(name, age) values ($1, $2)", createdUser.Name, createdUser.Age)
	if err != nil{
		log.Fatalf("tx.Exec failed %v", err)
	}
	tx.Commit(context.Background())
	defer tx.Rollback(context.Background())
	return createdUser, nil
}

func(server *UserManagementServer) GetUsers(ctx context.Context, params *pb.GetUsersParams)(*pb.UsersList, error){
	var usersList *pb.UsersList = &pb.UsersList{}

	rows, err := server.conn.Query(context.Background(), "SELECT * FROM users")
	if err!=nil{
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		user :=pb.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err !=nil{
			return nil, err
		}
		usersList.Users = append(usersList.Users, &user)
	}

	return usersList, nil
}

func main() {

	databaseUrl:="postgres://postgres:mysecretpassword@localhost:5432/postgres"
	conn, err:=pgx.Connect(context.Background(), databaseUrl)
	if err!=nil {
		log. Fatalf("Unable to establish connection: %v", err)
	}
	defer conn.Close(context.Background())
	var userManagementServer *UserManagementServer = NewUserManagementServer()
	userManagementServer.conn = conn
	if err := userManagementServer.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}