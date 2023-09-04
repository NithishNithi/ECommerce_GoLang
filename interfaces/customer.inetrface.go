package interfaces

import (
	"github.com/kishorens18/ecommerce/models"
)

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.CustomerDBResponse, error)
}
