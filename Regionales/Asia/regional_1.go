package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/MetalDanyboy/Lab1/protos"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)
var server_name string
var cant_registrados int
var cant_llaves_pedidas int

type Server struct {
	pb.UnimplementedChatServiceServer
	channel *amqp.Channel // Agregamos un campo para el canal de RabbitMQ
}

func Pedir_LLaves(cant_inicial int, cant_llaves_pedidas int)(int){

	if cant_inicial != 0 {
		if cant_llaves_pedidas == 0{
			num := int(cant_inicial/2)
			p := int(num/5)

			log.Printf("num: %d, p: %d", num, p)
			llaves_a_pedir := rand.Intn((num+p)-(num-p)) + (num - p)
	
			cant_llaves_pedidas=llaves_a_pedir
			log.Printf("llaves_a_pedir: %d, cant_llaves_pedidas: %d", llaves_a_pedir, cant_llaves_pedidas)
			return cant_llaves_pedidas
		}else{
			num := cant_inicial/2
			p := int(num* (1/5))
			llaves_a_pedir := rand.Intn((num+p)-(num-p)) + (num - p)
			cant_llaves_pedidas-=llaves_a_pedir
			return cant_llaves_pedidas
		}
	}else{
		return 0
	}
}


func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)

	// Enviamos un mensaje a RabbitMQ
	inMessage:=string(in.Body)
	fmt.Println("inMessage-->"+inMessage)
	if inMessage == "LLaves Disponibles"{
		cant_llaves_pedidas=Pedir_LLaves(cant_registrados,0)
		fmt.Println("cant_llaves_pedidas-->"+string(cant_llaves_pedidas))
		err := s.channel.Publish(
			"",        // exchange
			"testing", // key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				//Body:        []byte(server_name+"-"+string(cant_llaves_pedidas)),
				Body:       []byte(server_name+"-"+string(cant_llaves_pedidas)), // Enviamos el cuerpo del mensaje gRPC a RabbitMQ
			},
		)
		fmt.Println("Mande "+string(cant_llaves_pedidas)+" llaves")
		if err != nil {
			log.Printf("Error al publicar en RabbitMQ: %s", err)
		}
	}else if inMessage != "LLaves Disponibles"{
		registrados , _ := strconv.Atoi(in.Body)
		cant_registrados-=registrados
		if cant_registrados <= 0{
			return &pb.Message{Body: "OK"}, nil
		}else{
			cant_llaves_pedidas=Pedir_LLaves(cant_registrados,cant_llaves_pedidas)
			err := s.channel.Publish(
				"",        // exchange
				"testing", // key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(server_name+"-"+string(cant_llaves_pedidas)), // Enviamos el cuerpo del mensaje gRPC a RabbitMQ
				},
			)
			fmt.Println("Mande "+string(cant_llaves_pedidas)+" llaves")
			if err != nil {
				log.Printf("Error al publicar en RabbitMQ: %s", err)
			}
		}
	}

	return &pb.Message{Body: "OK"}, nil
}

func main() {
	

	directorioActual, err := os.Getwd()
    if err != nil {
        fmt.Println("Error al obtener el directorio actual:", err)
        return
    }

    fmt.Println("Directorio actual:", directorioActual)
	

	content, err := os.ReadFile(directorioActual+"/Regionales/Asia/parametros_de_inicio.txt")
	if err != nil {
		log.Fatal(err)
	}
	lineas := strings.Split(string(content), "\n")
	cant_registrados, err= strconv.Atoi(lineas[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n--->Cantidad de registrados: %d\n",cant_registrados)
	cant_llaves_pedidas=0

	server_name = "Asia"
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
