package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	grpcclient "github.com/kishorens18/ecommerce/cmd/client/gRPC_client"
	"github.com/kishorens18/ecommerce/cmd/client/routes"
)

func main() {
	fmt.Println("Client Server is Running")
	_, conn := grpcclient.GetGrpcClientInstance()
	defer conn.Close()
	r := gin.Default()
	routes.AppRoutes(r)
	r.Run(":8080")
}
