package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"
	//amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {

	//"localhost:50052"
	//"host.docker.internal:50052"
	//172.21.255.255:50052
	//regional:50052
	conn, err := grpc.Dial("172.21.255.255:50052",grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Printf("Esperando\n")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	for {

		response, err := c.SayHello(context.Background(), &pb.Message{Body: "Tengo llaves"})
		if err != nil {
			log.Panic("Error calling SendMessage: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}
		log.Printf("Response from server: %s", response.Body)
		break
	}	

}

