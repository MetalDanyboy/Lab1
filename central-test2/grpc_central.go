package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"
	"github.com/streadway/amqp"

	//amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConexionGRPC(mensaje string, servidor string , wg *sync.WaitGroup){
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
			log.Println("Server "+nombre+" not responding: ")
			log.Println("Trying again in 10 seconds. . .")
			time.Sleep(10 * time.Second)
			continue
		}
		log.Printf("Response from server "+nombre+": "+"%s", response.Body)
		break
	}
	defer wg.Done()
}

func main() {
	log.Println("Starting Central. . .\n")
	//"localhost:50052"
	//"host.docker.internal:50052"
	//172.21.255.255:50052
	//regional:50052
	//172.21.0.1:50052
	//"dist106.inf.santiago.usm.cl:50052"
	var wg sync.WaitGroup
	wg.Add(1)
	go ConexionGRPC("Hola desde el central","America", &wg)
	wg.Add(1)
	go ConexionGRPC("Hola desde el central","Asia", &wg)
	wg.Wait()


	//...CONEXION RABBITMQ...
	addr := "dist106.inf.santiago.usm.cl"
	//addr :="localhost"
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
	// ...
	
	//Mensaje Rabbit
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
			fmt.Printf("antes de mensaje_Cola\n")
			fmt.Printf("despues de mensaje_Cola\n")
			subcadenas := strings.Split(string(msg.Body), "-")
			
			if  subcadenas[0] == "Asia" {
				fmt.Printf("Entre a Asia\n")
				wg.Add(1)
				ConexionGRPC("200","Asia", &wg)
				
			}else if subcadenas[0] == "America"{

				ConexionGRPC("200","America", &wg)
			} else if subcadenas[0] == "Europa"{

				ConexionGRPC("200","Europa", &wg)
			} else if subcadenas[0] == "Oceania"{

				ConexionGRPC("200","Oceania", &wg)
			}else{
				fmt.Printf("No entre a ningun if")
			}
		}
	}()
	fmt.Println("Waiting for messages...")
	<-forever

	
	
	//...

	/*wg.Add(1)
	go ConexionGRPC("200","Asia", &wg)

	wg.Wait()
	log.Println("\nFinishing Central. . .")*/

}

