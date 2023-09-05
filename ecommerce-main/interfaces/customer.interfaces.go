package interfaces

import (
	"github.com/kishorens18/ecommerce/models"
	ecommerce "github.com/kishorens18/ecommerce/proto"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.CustomerDBResponse, error)
	CreateTokens(token *models.Token) (*ecommerce.Empty, error)
	UpdatePassword(Password *models.UpdatePassword) (*models.CustomerDBResponse, error)
}
