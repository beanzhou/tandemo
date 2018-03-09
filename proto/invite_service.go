package main

import (
	"fmt"

	pb "github.com/beanzhou/tandemo/proto/v1"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

type InviteService struct {
}

func (self InviteService) InviteList(ctx context.Context, request *pb.InviteRequest) (*pb.InviteResponse, error) {
	return &pb.InviteResponse{
		Page: request.Page,
		Size: request.Size,
		Data: []*pb.User{
			{"1", "12"},
			{"1", "12"},
			{"1", "12"},
			{"1", "12"},
			{"1", "12"},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	server := grpc.NewServer()
	pb.RegisterInviteServiceServer(server, InviteService{})
	server.Serve(lis)
}
