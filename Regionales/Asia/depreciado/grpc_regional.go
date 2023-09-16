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

//var llaves int

type Server struct {
	pb.UnimplementedChatServiceServer
	channel *amqp.Channel // Agregamos un campo para el canal de RabbitMQ
	mensaje string
}

/*func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from Central: %s", in.Body)
	message := in.GetBody()

    // Almacena el mensaje en la variable miembro.
    s.mensaje = message
	return &pb.Message{Body: "Hello From the Server!"}, nil
}*/


func (s Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
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
    if err != nil {
        log.Printf("Error al publicar en RabbitMQ: %s", err)
    }

    return &pb.Message{Body: "Hello From the Server!"}, nil
}
var grpcServer *grpc.Server
var server *Server

func StartGrpcServer(channel ) (socket net.Listener) {
	//Grpc
	puerto := ":50053"
	lis, err := net.Listen("tcp", puerto)
	fmt.Printf("Escuchando %s\n", puerto)
	if err != nil {
		log.Fatalf("'failed to listen: %v", err)
	}

	grpcServer = grpc.NewServer()
	server = &Server{channel: channel}
	pb.RegisterChatServiceServer(grpcServer, server)

	return lis
}

func ListenGrpcServer(lis net.Listener) {
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("\nfailed to serve: %s", err)
	}
}

func StopGrpcServer() {
	grpcServer.Stop()
}

func ServidorGRPC()(string){
	//Grpc
	lis:=StartGrpcServer()
	go ListenGrpcServer(lis)
	StopGrpcServer()


	/*puerto := ":50053"
	lis, err := net.Listen("tcp", puerto)
	fmt.Printf("Escuchando %s\n", puerto)
	if err != nil {
		log.Fatalf("'failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server = &Server{}
	pb.RegisterChatServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("\nfailed to serve: %s", err)
	}

	lis.Close()
	grpcServer.Stop()*/
	return server.mensaje
}

var msj string
func main() {
	//Conexion Rabbit
	addr_Rabbit := "dist106.inf.santiago.usm.cl"
	connection, err := amqp.Dial("amqp://guest:guest@"+addr_Rabbit+":5672/")
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

	// declaring queue with its properties over the the channel opened
	queue, err := channel.QueueDeclare(
		"testing", // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		panic(err)
	}

	//msj=ServidorGRPC()
	msj=ServidorGRPC()
	if msj == "Hola desde el central"{
		//Mensaje Rabbit
		err= channel.PublishWithContext(
			context.Background(),
			 "",
			"testing",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("186"),
			})
		/*err = channel.Publish(
			"",        // exchange
			"testing", // key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("186"),
			},
		)*/
		
		fmt.Println("Mande 186 llaves")
		if err != nil {
			panic(err)
		}

		fmt.Println("Queue status:", queue)
		fmt.Println("Successfully published message")
	}else{
		log.Println("No se pudo conectar con RabbitMQ")
	}

}
