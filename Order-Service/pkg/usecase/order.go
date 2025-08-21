package usecase

import (
	"errors"

	"github.com/jinzhu/copier"
	interfaceClient "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/client/interfaces"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/models"
	interfaceRepository "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/repository/interfaces"
	interfaceUseCase "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/usecase/interfaces"
)

type OrderUseCase struct {
	orderRepository   interfaceRepository.OrderRepository
	cartRepository    interfaceClient.CartClient
	productRepository interfaceClient.ProductClient
}

func NewOrderUseCase(repository interfaceRepository.OrderRepository, cartRepo interfaceClient.CartClient, productRepo interfaceClient.ProductClient) interfaceUseCase.OrderUseCase {
	return &OrderUseCase{
		orderRepository:   repository,
		cartRepository:    cartRepo,
		productRepository: productRepo,
	}
}

func (o *OrderUseCase) OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error) {
	var orderBody models.OrderIncoming
	err := copier.Copy(&orderBody, &orderFromCart)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	orderBody.UserID = userID
	cartExist, err := o.cartRepository.DoesCartExist(userID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	if !cartExist {
		return domain.OrderSuccessResponse{}, errors.New("cart empty can't order")
	}

	addressExist, err := o.orderRepository.AddressExist(orderBody)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if !addressExist {
		return domain.OrderSuccessResponse{}, errors.New("address does not exist")
	}

	paymentExist, err := o.orderRepository.PaymentExist(orderBody)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if !paymentExist {
		return domain.OrderSuccessResponse{}, errors.New("payment does not exist")
	}

	cartItems, err := o.cartRepository.GetAllItemsFromCart(orderBody.UserID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	total, err := o.cartRepository.TotalAmountInCart(orderBody.UserID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	FinalPrice := total
	order_id, err := o.orderRepository.OrderItems(orderBody, FinalPrice)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	if err := o.orderRepository.AddOrderProducts(order_id, cartItems); err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	orderSuccessResponse, err := o.orderRepository.GetBriefOrderDetails(order_id)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	var orderItemDetails domain.OrderItem
	for _, c := range cartItems {
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = c.Quantity
		err := o.cartRepository.UpdateCartAfterOrder(userID, int(orderItemDetails.ProductID), orderItemDetails.Quantity)
		if err != nil {
			return domain.OrderSuccessResponse{}, err
		}
		err = o.productRepository.ProductStockMinus(int(orderItemDetails.ProductID), int(orderItemDetails.Quantity))
		if err != nil {
			return domain.OrderSuccessResponse{}, err
		}
	}
	return orderSuccessResponse, nil
}

func (o *OrderUseCase) GetOrderDetails(userID int) ([]models.FullOrderDetails, error) {
	fullOrderDetails, err := o.orderRepository.GetOrderDetails(userID)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}
	return fullOrderDetails, nil
}
