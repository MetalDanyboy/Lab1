package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MetalDanyboy/Lab1/protos"
	//amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)


func main() {

	//
	//"host.docker.internal:50052"
	conn, err := grpc.Dial("dist106:50052", grpc.WithInsecure())
	fmt.Printf("Esperando\n")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)


	response, err := c.SayHello(context.Background(), &pb.Message{Body: "Tengo llaves"})
	if err != nil {
		log.Fatalf("Error calling SendMessage: %v", err)
	}

	log.Printf("Response from server: %s", response.Body)


}

