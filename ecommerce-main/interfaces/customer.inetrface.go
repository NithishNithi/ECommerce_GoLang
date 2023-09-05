package interfaces

import (
	"github.com/kishorens18/ecommerce/models"
	ecommerce "github.com/kishorens18/ecommerce/proto"
)

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.CustomerDBResponse, error)
	CustomerLogin(email string, password string) (*models.CustomerDBResponse, error)
	CreateTokens(token *models.Token) (*ecommerce.Empty, error)
	
}
