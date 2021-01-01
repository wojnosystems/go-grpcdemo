package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcdemo/pkg/proto"
	"log"
	"net"
)

type dogSayService struct {
	proto.UnimplementedDogSayServiceServer
}

func (s *dogSayService) Say(ctx context.Context, in *proto.BarkRequest) (*proto.BarkResponse, error) {
	log.Printf("Received: %v", in.GetName())
	response := fmt.Sprintf("Hello %s, you're age %d!", in.GetName(), in.GetAge())
	return &proto.BarkResponse{Out: response}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8500")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterDogSayServiceServer(s, &dogSayService{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
