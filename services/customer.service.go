package services

import (
	"context"

	"github.com/kishorens18/ecommerce/interfaces"
	"github.com/kishorens18/ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	ProfileCollection *mongo.Collection
	ctx               context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (p *CustomerService) CreateCustomer(user *models.Customer) (*models.CustomerDBResponse, error) {
	user.Firstname = "123"
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
