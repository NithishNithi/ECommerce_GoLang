package controllers

import (
	"context"

	"github.com/kishorens18/ecommerce/interfaces"
	"github.com/kishorens18/ecommerce/models"
	pro "github.com/kishorens18/ecommerce/proto"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerDetails) (*pro.CustomerResponse, error) {
	dbCustomer := models.Customer{Customer_id: req.Customer_ID}
	result, err := CustomerService.CreateCustomer(&dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			Customer_ID: result.Customer_id,
		}
		return responseCustomer, nil
	}
}
