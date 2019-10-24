package main

import (
	context "context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/grpchw"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedOlapServer
}

func (s *server) GetMessage(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.MsgReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("start the server")

	s := grpc.NewServer()
	pb.RegisterOlapServer(s, &server{})

	log.Println("start grpc")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
