package services

import (
	"github.com/kishorens18/ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p *CustomerService) UpdateCustomer(user *models.UpdateRequest) (*models.Customer, error) {
	filter := bson.D{{Key: "customerid", Value: user.CustomerId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: user.Field, Value: user.Value}}}}
	options := options.Update()

	_, err := p.ProfileCollection.UpdateOne(p.ctx, filter, update, options)
	if err != nil {
		return nil, err
	}

	// Fetch the updated user document
	var updatedUser models.Customer
	err = p.ProfileCollection.FindOne(p.ctx, filter).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
