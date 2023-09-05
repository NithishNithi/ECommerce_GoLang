package services

import (
	"context"
	"fmt"
	"log"

	"github.com/kishorens18/ecommerce/interfaces"
	"github.com/kishorens18/ecommerce/models"
	ecommerce "github.com/kishorens18/ecommerce/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService struct {
	ProfileCollection *mongo.Collection
	tokenCollection   *mongo.Collection
	ctx               context.Context
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Verify(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err != nil
}

func InitCustomerService(collection *mongo.Collection, tokenCollection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, tokenCollection, ctx}
}

func (p *CustomerService) CreateCustomer(user *models.Customer) (*models.CustomerDBResponse, error) {
	user.HashesAndSaltedPassword, _ = HashPassword(user.HashesAndSaltedPassword)
	res, err := p.ProfileCollection.InsertOne(p.ctx, &user)

	if err != nil {
		return nil, err
	}

	var newUser *models.CustomerDBResponse
	query := bson.M{"_id": res.InsertedID}

	err = p.ProfileCollection.FindOne(p.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (p *CustomerService) UpdatePassword(user *models.UpdatePassword) (*models.CustomerDBResponse, error) {
	var customer *models.Customer
	fmt.Println(user.NewPassword)
	fmt.Println(user.OldPassword)
	if user.OldPassword != user.NewPassword {
		query := bson.M{"email": user.Email}
		err := p.ProfileCollection.FindOne(p.ctx, query).Decode(&customer)

		if err != nil {
			fmt.Println("11")

			return nil, err // Authentication failed
		}
		fmt.Println("21")
		if Verify(customer.HashesAndSaltedPassword, user.OldPassword) == true {
			fmt.Println("12")

			user.NewPassword, _ = HashPassword(user.NewPassword)
			filter := bson.M{"email": user.Email}
			update := bson.M{"$set": bson.M{"hashedandsaltedpassword": user.NewPassword}}
			fmt.Println("13")

			_, err := p.ProfileCollection.UpdateOne(context.Background(), filter, update)
			fmt.Println("14")

			if err != nil {
				fmt.Println("15")

				log.Fatal(err)
			}

		} else {
			fmt.Println("16")

			return nil, nil
		}

	}
	response := models.CustomerDBResponse{
		Customer_id: customer.CustomerId,
	}
	return &response, nil
}

func (p *CustomerService) CustomerLogin(email string, password string) (*models.CustomerDBResponse, error) {
	// Validate the login credentials (email and password)
	// Perform the login logic here, checking if the email and password match
	// You can use your existing MongoDB collection to look up the user by email and compare passwords

	// For example:
	query := bson.M{"email": email, "hashesandsaltedpassword": password}

	var customer *models.Customer
	err := p.ProfileCollection.FindOne(p.ctx, query).Decode(&customer)
	if err != nil {
		return nil, err // Authentication failed
	}

	// If authentication succeeds, you can return the customer information or an authentication token
	// For simplicity, we're returning the customer ID here
	responseCustomer := &models.CustomerDBResponse{
		Customer_id: customer.CustomerId,
	}

	return responseCustomer, nil
}

func (p *CustomerService) CreateTokens(user *models.Token) (*ecommerce.Empty, error) {
	res, err := p.tokenCollection.InsertOne(p.ctx, &user)

	if err != nil {
		return nil, err
	}

	var newUser *models.Token
	query := bson.M{"_id": res.InsertedID}

	err = p.ProfileCollection.FindOne(p.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
