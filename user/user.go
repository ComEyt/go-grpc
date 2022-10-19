package main

import (
	"context"
	"fmt"
	"go-grpc/protoc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type User struct {
	protoc.UnimplementedUserServer
}

func (u *User) GetUser(ctx context.Context, r *protoc.GetUserRequest) (*protoc.GetUserResponse, error) {
	if r.Username != nil {
		fmt.Println(*r.Username)
		return &protoc.GetUserResponse{}, nil
	}
	fmt.Println("nil name")
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen: %v",err)
	}
	s := grpc.NewServer()
	protoc.RegisterUserServer(s,&User{})
	if err := s.Serve(lis);err != nil {
		log.Fatalf("failed to serve :%v",err)
	}
}
