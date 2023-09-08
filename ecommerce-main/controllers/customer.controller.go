package controllers

import (
	"context"
	"fmt"

	"log"

	"github.com/kishorens18/ecommerce/constants"
	"github.com/kishorens18/ecommerce/interfaces"
	"github.com/kishorens18/ecommerce/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pro "github.com/kishorens18/ecommerce/proto"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	// ctx             gin.Context
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
		CustomerId:      req.CustomerId,
		Firstname:       req.Firstname,
		Lastname:        req.Lastname,
		Password:        req.Password,
		Email:           req.Email,
		Address:         []models.Address{address},
		ShippingAddress: []models.ShippingAddress{shippingAddress},
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
		fmt.Println(result.Customer_id)
		responseCustomer := &pro.CustomerResponse{
			Customer_ID: result.Customer_id,
		}
		return responseCustomer, nil
	}
}

func (s *RPCServer) UpdateCustomer(ctx context.Context, req *pro.UpdateDetails) (*pro.CustomerResponse, error) {
    // Check if the request is nil

    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "Request is nil")
    }

    // Validate the request fields

    if req.CustomerId == "" || req.Field == "" || req.OldValue == "" || req.NewValue == "" {
        fmt.Println("error fpound")
        return nil, status.Error(codes.InvalidArgument, "Missing required fields")
    }

    var cus models.UpdateRequest
    if req != nil {
        cus = models.UpdateRequest{
            CustomerId: req.CustomerId,
            Field:      req.Field,
            OldValue:   req.OldValue,
            NewValue:   req.NewValue,
        }
    }

    // Call the UpdateCustomer service function
    updatedUser, err := CustomerService.UpdateCustomer(&cus)
    if err != nil {
        fmt.Println("error from service", err)
        return nil, status.Error(codes.Internal, "Failed to update user")
    }

    // Create and return the response
    responseCustomer := &pro.CustomerResponse{
        Customer_ID: updatedUser.Customer_id,
    }

    return responseCustomer, err
}

func (s *RPCServer) DeleteCustomer(ctx context.Context, req *pro.DeleteDetails) (*pro.Empty, error) {
	t1 := req.Token
	customerID, err := ExtractCustomerID(t1, constants.SecretKey)

	if err != nil {
		return nil, err // Return the error from the gRPC function
	}
	err1 := CustomerService.DeleteCustomer(customerID)
	if err1 != nil {
		return nil, err1 // Return the error from the gRPC function
	}
	return &pro.Empty{}, nil // Return a valid response
}

func (s *RPCServer) GetByCustomerId(ctx context.Context, req *pro.GetbyId) (*pro.CustomerDetails, error) {
	t1 := req.Token
	// Extract customer ID from the token
	customerID, err := ExtractCustomerID(t1, constants.SecretKey)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Retrieve customer details from CustomerService
	result, err := CustomerService.GetByCustomerId(customerID)
	if err != nil {
		// Handle the error (e.g., log it or return an error response)
		return nil, err
	}

	// Create and populate the responseCustomer struct
	responseCustomer := &pro.CustomerDetails{
		CustomerId: result.CustomerId,
		Firstname:  result.Firstname,
		Lastname:   result.Lastname,
		Email:      result.Email,
	}

	// Check if there are addresses and populate them
	if len(result.Address) > 0 {
		responseCustomer.Address = []*pro.Address{
			{
				Country: result.Address[0].Country,
				Street1: result.Address[0].Street1,
				Street2: result.Address[0].Street2,
				City:    result.Address[0].City,
				State:   result.Address[0].State,
				Zip:     result.Address[0].Zip,
			},
			// Add more address entries if needed.
		}
	}

	// Check if there are shipping addresses and populate them
	if len(result.ShippingAddress) > 0 {
		responseCustomer.ShippingAddress = []*pro.ShippingAddress{
			{
				Street1: result.ShippingAddress[0].Street1,
				Street2: result.ShippingAddress[0].Street2,
				City:    result.ShippingAddress[0].City,
				State:   result.ShippingAddress[0].State,
			},
			// Add more shipping address entries if needed.
		}
	}

	return responseCustomer, nil
}
