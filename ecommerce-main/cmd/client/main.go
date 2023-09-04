package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	pb "github.com/kishorens18/ecommerce/proto"
	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

const (
	secretKey = "your-secret-key"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"hashedandsaltedpassword"`
}

// func main() {
// 	r := gin.Default()

// 	// Connect to gRPC service
// 	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Failed to connect: %v", err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewCustomerServiceClient(conn)

// 	// response, err := client.CreateCustomer(context.Background(), &pb.CustomerDetails{})
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to call CreateTransaction: %v", err, response)
// 	// }

// 	// Define a POST route
// 	r.POST("/signin", func(c *gin.Context) {
// 		var request pb.CustomerDetails

// 		// Parse incoming JSON
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Call the gRPC service
// 		response, err := client.CreateCustomer(c.Request.Context(), &request)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"value": response})
// 	})

// 	r.POST("/login", func(c *gin.Context) {
// 		var request pb.CustomerLoginRequest

// 		// Parse incoming JSON
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Call the gRPC service for customer login
// 		response, err := client.CustomerLogin(c.Request.Context(), &request)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"customer_id": response.Customer_ID})
// 	})

// 	// Start the Gin server
// 	if err := r.Run(":8080"); err != nil {
// 		fmt.Println("Failed to start Gin server:", err)
// 	}
// }

func main() {
	r := gin.Default()
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	r.POST("/signin", func(c *gin.Context) {
		var request pb.CustomerDetails

		// Parse incoming JSON
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the gRPC service
		response, err := client.CreateCustomer(c.Request.Context(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"value": response})
	})

	r.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			fmt.Println("1")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		// Simulated authentication (replace with your actual authentication logic)
		if isValidUser(user) {
			token, err := createToken(user.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed"})
				return
			}
			response1, err := client.CreateTokens(c.Request.Context(), &pb.Token{Email: user.Email, Token: token})
			fmt.Println(response1)
			c.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	})

	r.Run(":8080")
}

func isValidUser(user User) bool {
	// Simulated user validation (replace with your actual validation logic)
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	return user.Email == "jp@gmail.com" && user.Password == "123"
}

func createToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
