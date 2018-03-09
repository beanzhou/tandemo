package main

import (
	"fmt"

	"context"
	pb "github.com/beanzhou/tandemo/proto/v1"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("connet err: %+v", err)
	}

	client := pb.NewInviteServiceClient(conn)
	list, err := client.InviteList(context.Background(), &pb.InviteRequest{UserId: "1", Page: 1, Size: 10})
	if err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Printf("%+v", list)

}
