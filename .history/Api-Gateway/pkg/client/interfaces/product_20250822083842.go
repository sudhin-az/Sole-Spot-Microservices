package interfaces

import "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"

type ProductClient interface {
	ListProducts(userID int) ([]models.ProductBrief, error)
	AddProduct(product models.Product) (models.Products, error)
	DeleteProduct(id int) error
	UpdateProducts(product models.Product) (models.Products, error)
}
