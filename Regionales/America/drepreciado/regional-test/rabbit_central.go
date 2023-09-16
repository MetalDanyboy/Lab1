package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {

	//addr := "dist106.inf.santiago.usm.cl:50052"
	addr :="localhost"
    //Conexion rabbit
	connection, err := amqp.Dial("amqp://guest:guest@"+addr+":5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		"testing", // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       //args
	)
	if err != nil {
		panic(err)
	}


	//Grpc
	
	conn, err := grpc.Dial(addr+":50052",grpc.WithTransportCredentials(insecure.NewCredentials()))
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


	//Mensaje Rabbit
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
