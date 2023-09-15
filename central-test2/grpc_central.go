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

func ConexionGRPC(mensaje string, servidor string){
	var host string
	var puerto string
	var nombre string
	//Uno de estos debe cambiar quizas por "regional:50052" ya que estara en la misma VM que el central
	if servidor == "America"{
		host="dist105.inf.santiago.usm.cl"
		puerto="50052"
		nombre="America"
	}else if servidor == "Asia"{
		
		host="dist106.inf.santiago.usm.cl"
		puerto="50053"
		nombre="Asia"
	}else if servidor == "Europa"{

		host="dist107.inf.santiago.usm.cl"
		puerto="50054"
		nombre="Europa"
	}else if servidor == "Oceania"{
		
		host="dist108.inf.santiago.usm.cl"
		puerto="50055"
		nombre="Oceania"
	}
	log.Println("Connecting to server "+nombre+": "+host+":"+puerto+". . .")
	conn, err := grpc.Dial(host+":"+puerto,grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	fmt.Printf("Esperando\n")
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)
	for {
		log.Println("Sending message to server "+nombre+": "+mensaje)
		response, err := c.SayHello(context.Background(), &pb.Message{Body: mensaje})
		if err != nil {
			log.Println("Server "+nombre+" not responding: "+err.Error())
			log.Println("Trying again in 10 seconds. . .")
			time.Sleep(5 * time.Second)
			continue
		}
		log.Printf("Response from server "+nombre+": "+"%s", response.Body)
		break
	}
}

func main() {

	//"localhost:50052"
	//"host.docker.internal:50052"
	//172.21.255.255:50052
	//regional:50052
	//172.21.0.1:50052
	//"dist106.inf.santiago.usm.cl:50052"
	ConexionGRPC("Hola desde el central","America")
	ConexionGRPC("Hola desde el central","Asia")

}

