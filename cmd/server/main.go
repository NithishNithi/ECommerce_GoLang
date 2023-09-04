package main

import (
	"context"
	"fmt"
	"net"

	"github.com/kishorens18/ecommerce/config"
	"github.com/kishorens18/ecommerce/constants"
	"github.com/kishorens18/ecommerce/controllers"
	pro "github.com/kishorens18/ecommerce/proto"
	"github.com/kishorens18/ecommerce/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	profileCollection := config.GetCollection(client, "bankdb", "profiles")
	controllers.CustomerService = services.InitCustomerService(profileCollection, context.Background())
}

func main() {
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
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
