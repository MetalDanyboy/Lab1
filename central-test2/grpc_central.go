package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MetalDanyboy/Lab1/protos"
	//amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {

	//"localhost:50052"
	//"host.docker.internal:50052"
	conn, err := grpc.Dial("regional:50052",grpc.WithTransportCredentials(insecure.NewCredentials()))
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

