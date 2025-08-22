package usecase

import (
	"errors"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/models"
	interfaceRepository "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/repository/interfaces"
)

type ProductUseCase struct {
	ProductRepository interfaceRepository.ProductRepository
}

func NewProductUseCase(repository interfaceRepository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		ProductRepository: repository,
	}
}

func (pu *ProductUseCase) AddProduct(product models.Product) (domain.Product, error) {
	exist := pu.ProductRepository.ProductAlreadyExist(product.Name)
	if exist {
		return domain.Product{}, errors.New("product already exits")
	}
	productResponse, err := pu.ProductRepository.AddProduct(product)
	if err != nil {
		return domain.Product{}, err
	}
	stock := pu.ProductRepository.StockInvalid(productResponse.Name)
	if !stock {
		return domain.Product{}, errors.New("stock is invalid input")
	}
	return productResponse, nil
}

func (pu *ProductUseCase) ListProducts(userId int) ([]models.ProductBrief, error) {

	products, err := pu.ProductRepository.ListProducts()
	if err != nil {
		return []models.ProductBrief{}, err
	}
	return products, nil
}

func (pu *ProductUseCase) DeleteProducts(id int) error {
	err := pu.ProductRepository.DeleteProducts(id)
	if err != nil {
		return err
	}
	return nil
}

func (pu *ProductUseCase) UpdateProducts(product models.Product) (domain.Product, error) {

	return pu.ProductRepository.UpdateProducts(product)
}

func (pu *ProductUseCase) GetQuantityFromProductID(id int) (int, error) {
	quantity, err := pu.ProductRepository.GetQuantityFromProductID(id)
	if err != nil {
		return 0.0, err
	}
	return quantity, nil
}

func (pu *ProductUseCase) GetPriceofProductFromID(id int) (float64, error) {
	price, err := pu.ProductRepository.GetPriceofProductFromID(id)
	if err != nil {
		return 0.0, err
	}
	return price, nil
}

func (pu *ProductUseCase) ProductStockMinus(productID, stock int) error {
	err := pu.ProductRepository.ProductStockMinus(productID, stock)
	if err != nil {
		return err
	}
	return nil
}

func (pu *ProductUseCase) CheckProduct(product_id int) (bool, error) {
	ok, err := pu.ProductRepository.CheckProduct(product_id)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, err
	}
	return true, nil
}
