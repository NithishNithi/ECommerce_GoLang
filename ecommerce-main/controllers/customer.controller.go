package controllers

import (
	"context"


	"github.com/gin-gonic/gin"
	"github.com/kishorens18/ecommerce/interfaces"
	"github.com/kishorens18/ecommerce/models"

	pro "github.com/kishorens18/ecommerce/proto"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	ctx             gin.Context
	CustomerService interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerDetails) (*pro.CustomerResponse, error) {
	var address models.Address
	if req != nil {
		address = models.Address{
			Country: req.Address[0].Country,
			Street1: req.Address[0].Street1,
			Street2: req.Address[0].Street2,
			City:    req.Address[0].City,
			State:   req.Address[0].State,
			Zip:     req.Address[0].Zip,
		}
	}

	var shippingAddress models.ShippingAddress
	if req != nil {
		shippingAddress = models.ShippingAddress{
			Street1: req.ShippingAddress[0].Street1,
			Street2: req.ShippingAddress[0].Street2,
			City:    req.ShippingAddress[0].City,
			State:   req.ShippingAddress[0].State,
		}
	}

	dbCustomer := models.Customer{
		CustomerId:              req.CustomerId,
		Firstname:               req.Firstname,
		Lastname:                req.Lastname,
		HashesAndSaltedPassword: req.HashesAndSaltedPassword,
		Email:                   req.Email,
		Address:                 []models.Address{address},
		ShippingAddress:         []models.ShippingAddress{shippingAddress},
	}
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

func (s *RPCServer) CreateTokens(ctx context.Context, req *pro.Token) (*pro.Empty, error) {

	dbCustomer := models.Token{Email: req.Email, Token: req.Token, CustomerId: req.Customerid}
	_, err := CustomerService.CreateTokens(&dbCustomer)
	if err != nil {
		return nil, err
	} else {

		return nil, nil
	}
}

func (s *RPCServer) UpdatePassword(ctx context.Context, req *pro.PasswordDetails) (*pro.CustomerResponse, error) {
	var pass models.UpdatePassword
	if req != nil {
		pass = models.UpdatePassword{
			Email:       req.Email,
			OldPassword: req.OldPassword,
			NewPassword: req.NewPassword,
		}
	}

	result, err := CustomerService.UpdatePassword(&pass)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			Customer_ID: result.Customer_id,
		}
		return responseCustomer, nil
	}
}

func (s *RPCServer) UpdateEmail(ctx context.Context, req *pro.EmailDetails) (*pro.CustomerResponse, error) {
	var mail models.UpdateEmail
	if req != nil {
		mail = models.UpdateEmail{
			CustomerId:       req.CustomerId,
			OldEmail: req.OldEmail,
			NewEmail: req.NewEmail,
		}
	}

	result, err := CustomerService.UpdateEmail(&mail)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			Customer_ID: result.Customer_id,
		}
		return responseCustomer, nil
	}
}
