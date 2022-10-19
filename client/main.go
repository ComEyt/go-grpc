package main

import (
	"context"
	"fmt"
	"go-grpc/protoc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const defaultName = "world"

func main() {
	// 设置server的链接
	conn, err := grpc.Dial("127.0.0.1:50001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect:%v",err)
	}
	defer conn.Close()
	client2 := protoc.NewUserClient(conn)
	//client := pb.NewGreeterClient(conn)
	// 联系服务器并打印它的响应
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//r, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	s := "xxx"
	user, err := client2.GetUser(ctx, &protoc.GetUserRequest{Username: &s})
	if err != nil {
		log.Fatalf("could not greet : %v",err)
	}
	//ss := struct {
	//	Message string `json:"message"`
	//}{Message: r.Message}
	fmt.Println(user.Username)
	//log.Printf("Greeting: %s",ss)
	//fmt.Println(ss)

}
