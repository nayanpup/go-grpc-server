package main

import (
	"context"
	"log"
	"net"

	pb "golib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port        = ":50051"
	placeholder = "How much does a hipster weigh? An instagram."
)

type server struct{}

func (s *server) GetGreetings(c context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	msg := &pb.HelloRequest{msg: placeholder}
	return msg, nil
}

func main() {
	log.Println("starting DadJoke server🤓")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
	}
	s := grpc.NewServer()
	pb.(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
