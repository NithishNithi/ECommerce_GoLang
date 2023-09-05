package services

import (
	"context"

	"github.com/kishorens18/ecommerce/interfaces"
	"github.com/kishorens18/ecommerce/models"
	ecommerce "github.com/kishorens18/ecommerce/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	ProfileCollection *mongo.Collection
	tokenCollection   *mongo.Collection
	ctx               context.Context
}

func InitCustomerService(collection *mongo.Collection, tokenCollection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, tokenCollection, ctx}
}

func (p *CustomerService) CreateCustomer(user *models.Customer) (*models.CustomerDBResponse, error) {
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

