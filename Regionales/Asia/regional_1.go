package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/MetalDanyboy/Lab1/protos"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	channel *amqp.Channel // Agregamos un campo para el canal de RabbitMQ
}

func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)

	// Enviamos un mensaje a RabbitMQ
	err := s.channel.Publish(
		"",        // exchange
		"testing", // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("186"), // Enviamos el cuerpo del mensaje gRPC a RabbitMQ
		},
	)
	fmt.Println("Mande 186 llaves")
	if err != nil {
		log.Printf("Error al publicar en RabbitMQ: %s", err)
	}
	

	return &pb.Message{Body: "Hello From the Server!"}, nil
}

func main() {
	addr_Rabbit := "dist106.inf.santiago.usm.cl"
	connection, err := amqp.Dial("amqp://guest:guest@" + addr_Rabbit + ":5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"testing",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	puerto := ":50053"
	lis, err := net.Listen("tcp", puerto)
	fmt.Printf("Escuchando %s\n", puerto)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	server := &Server{channel: channel} // Pasamos el canal de RabbitMQ al servidor gRPC
	pb.RegisterChatServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	err = channel.Publish(
		"",
		"testing",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("186"),
		},
	)

	fmt.Println("Mande 186 llaves")
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Successfully published message")
}
