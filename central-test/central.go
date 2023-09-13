package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)


func main() {

    //Conexion Grpc
	//##############################################
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	//##############################################


	//Enviar Mensaje 1 - Grpc (Tengo llaves)
	//##############################################
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Tengo llaves"})
	if err != nil {
		log.Fatalf("Error send msj: %s", err)
	}
	log.Printf("Response from regional: %s", response.Body)
	//##############################################


    //Conexion Rabbit (en la misma VM)
    //##############################################
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

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
	//##############################################


	//Recibir Mensaje  - Rabbit (Usuarios interesados)
	//##############################################
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
	//##############################################


	//Enviar Mensaje 2 - Grpc (n° de registrados)
	//##############################################
	response, err := c.SayHello(context.Background(), &chat.Message{Body: msg.Body+" Msg 2"})
	if err != nil {
		log.Fatalf("Error send msj: %s", err)
	}
	log.Printf("Response from regional: %s", response.Body)