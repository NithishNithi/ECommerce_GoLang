package main

import (
	"context"
	"fmt"
	"net"

	"github.com/kishorens18/ecommerce/config"
	"github.com/kishorens18/ecommerce/constants"
	"github.com/kishorens18/ecommerce/controllers"
	"github.com/kishorens18/ecommerce/services"

	pro "github.com/kishorens18/ecommerce/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	profileCollection := config.GetCollection(client, "Ecommerce", "CustomerProfile")
	tokenCollection := config.GetCollection(client, "Ecommerce", "Tokens")
	controllers.CustomerService = services.InitCustomerService(profileCollection, tokenCollection, context.Background())

}

func main() {
	fmt.Println("hello")
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &controllers.RPCServer{})
	fmt.Println("Server listening on", constants.Port)
	fmt.Println("1")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
