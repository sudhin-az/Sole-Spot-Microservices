package interfaceUseCase

import (
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/models"
)

type ProductUseCase interface {
	AddProduct(product models.Product) (domain.Product, error)
	ListProducts(userId int) ([]models.ProductBrief, error)
	ProductAlreadyExist(Name string) bool
	StockInvalid(Name string) bool
	DeleteProducts(id int) error
	CheckProduct(id int) (bool, error)
	UpdateProducts(product models.Product) (domain.Product, error)
	GetQuantityFromProductID(id int) (int, error)
	GetPriceofProductFromID(product_id int) (float64, error)
	ProductStockMinus(productID, stock int) error
}
