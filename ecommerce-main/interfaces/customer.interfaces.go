package interfaces

import (
	"github.com/kishorens18/ecommerce/models"
)

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.CustomerDBResponse, error)
	CreateTokens(token *models.Token) (*models.TokenResponse, error)
	UpdatePassword(Password *models.UpdatePassword) (*models.CustomerDBResponse, error)
	UpdateCustomer(cus *models.UpdateRequest) (*models.CustomerDBResponse, error)
	DeleteCustomer(res string) error
	GetByCustomerId(res string) (*models.Customer, error)
}
