package main

import (
	"context"
	pb "go-grpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)


// server被用来实现helloworld.GreeterServer
type server struct {
	pb.UnimplementedGreeterServer
}

//SayHello 实现了helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v",in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()},nil
}

func main() {
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen: %v",err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})
	if err := s.Serve(lis);err != nil {
		log.Fatalf("failed to serve :%v",err)
	}

}
