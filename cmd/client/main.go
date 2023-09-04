package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	pb "github.com/kishorens18/ecommerce/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &pb.CustomerDetails{})

}
