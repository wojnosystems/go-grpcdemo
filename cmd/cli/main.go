package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcdemo/pkg/proto"
	"log"
	"time"
)

func main() {
	fmt.Println("client start!")
	// Set up a connection to the server.
	conn, err := grpc.Dial("backend:8500", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cli := proto.NewDogSayServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := cli.Say(ctx, &proto.BarkRequest{Name: "Clifford", Age: 5})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	dogSay(r.GetOut())
}
