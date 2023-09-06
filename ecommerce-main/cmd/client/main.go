package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kishorens18/ecommerce/config"
	"github.com/kishorens18/ecommerce/models"
	pb "github.com/kishorens18/ecommerce/proto"
	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson"
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
	Email      string `json:"email"`
	Password   string `json:"hashedandsaltedpassword"`
	CustomerId string `json:"customerid"`
}

func main() {
	r := gin.Default()
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	r.POST("/signup", func(c *gin.Context) {
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

	r.POST("/signin", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		if isValidUser(user) {
			token, err := createToken(user.Email, user.CustomerId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed"})
				return
			}
			response1, err := client.CreateTokens(c.Request.Context(), &pb.Token{Email: user.Email, Token: token, Customerid: user.CustomerId})
			fmt.Println(response1)
			c.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	})

	r.POST("/updatecustomer", func(c *gin.Context) {
		var user models.UpdateRequest
		

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updatedUser, err := client.UpdateCustomer(c.Request.Context(),&pb.UpdateDetails{CustomerId: user.CustomerId,
		Field: user.Field, OldValue: user.OldValue,NewValue: user.NewValue,})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated", "data": updatedUser})

	})
	r.POST("/deletecustomer", func(c *gin.Context) {
		var user models.DeleteRequest

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		client.DeleteCustomer(c.Request.Context(), &pb.DeleteDetails{CustomerID: user.CustomerId})

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

	})


	r.POST("/resetpassword", func(c *gin.Context) {
		var user models.UpdatePassword
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		response, err := client.UpdatePassword(c.Request.Context(), &pb.PasswordDetails{Email: user.Email, OldPassword: user.OldPassword, NewPassword: user.NewPassword})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"value": response})
	})

	r.Run(":8080")
}

func isValidUser(user User) bool {
	// Simulated user validation (replace with your actual validation logic)
	mongoclient, _ := config.ConnectDataBase()
	collection := mongoclient.Database("Ecommerce").Collection("CustomerProfile")
	filter := bson.M{"email": user.Email, "hashedandsaltedpassword": user.Password, "customerid": user.CustomerId}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {

		fmt.Println("ERROR")
		return false
	}
	return count > 0
}

func createToken(email, customerid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email, "customerid": customerid,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
