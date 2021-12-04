package main

import (
	api "CourseProject/pkg/api"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	if flag.NArg()<2{
		log.Fatal("not enough arguments")
	}

	x, err:=strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err:=strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err:=grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c:=api.NewOrderingClient(conn)

	res, err := c.MakeOrder(context.Background(), &api.OrderInfo{GameName: int32(y), Fio: int32(x)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetStatus())
}
