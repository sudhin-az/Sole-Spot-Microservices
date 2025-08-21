package usecase

import (
	"errors"

	interfaceClient "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/client/interfaces"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/models"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/repository"
)

type cartUseCase struct {
	CartRepository   repository.CartRepository
	ProducRepository interfaceClient.NewProductClient
}

func NewCartUseCase(repository repository.CartRepository, producRepository interfaceClient.NewProductClient) *cartUseCase {
	return &cartUseCase{
		CartRepository:   repository,
		ProducRepository: producRepository,
	}
}

func (cr *cartUseCase) AddToCart(product_id, user_id, quantity int) (models.CartResponse, error) {
	ok, err := cr.ProducRepository.CheckProduct(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, errors.New("product Does not exist")
	}
	QuantityOfProductInCart, err := cr.CartRepository.QuantityOfProductInCart(user_id, product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	quantityOfProduct, err := cr.ProducRepository.GetQuantityFromProductID(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	if quantityOfProduct <= 0 {
		return models.CartResponse{}, errors.New("out of stock")
	}
	if quantityOfProduct == QuantityOfProductInCart {
		return models.CartResponse{}, errors.New("stock limit exceeded")
	}
	ok, err = cr.CartRepository.ProductExist(user_id, product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, err
	}
	productPrice, err := cr.ProducRepository.GetPriceOfProductFromID(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, err
	}
	FinalPrice := productPrice * float64(quantity)
	if QuantityOfProductInCart == 0 {
		err := cr.CartRepository.AddItemIntoCart(user_id, product_id, quantity, FinalPrice)
		if err != nil {
			return models.CartResponse{}, err
		}
	} else {
		currentTotal, err := cr.CartRepository.TotalPriceForProductInCart(user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
		err = cr.CartRepository.UpdateCart(QuantityOfProductInCart, int(currentTotal)+int(productPrice), user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
	}
	cartDetails, err := cr.CartRepository.DisplayCart(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := cr.CartRepository.GetTotalPrice(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	err = cr.ProducRepository.ProductStockMinus(product_id, QuantityOfProductInCart)
	if err != nil {
		return models.CartResponse{}, err
	}
	return models.CartResponse{
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cartDetails,
	}, nil
}

func (cr *cartUseCase) DisplayCart(user_id int) (models.CartResponse, error) {
	cart, err := cr.CartRepository.DisplayCart(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := cr.CartRepository.GetTotalPrice(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	return models.CartResponse{
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cart,
	}, nil
}

func (cr *cartUseCase) GetAllItemsFromCart(userID int) ([]models.Cart, error) {
	res, err := cr.CartRepository.GetAllItemsFromCart(userID)
	if err != nil {
		return []models.Cart{}, err
	}
	return res, err
}

func (cr *cartUseCase) DoesCartExist(userID int) (bool, error) {
	res, err := cr.CartRepository.DoesCartExist(userID)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (cr *cartUseCase) TotalAmountInCart(userID int) (float64, error) {
	res, err := cr.CartRepository.TotalAmountInCart(userID)
	if err != nil {
		return 0.0, err
	}
	return res, nil
}

func (cr *cartUseCase) UpdateCartAfterOrder(userID int, productID int, quantity float64) error {
	err := cr.CartRepository.UpdateCartAfterOrder(userID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}
