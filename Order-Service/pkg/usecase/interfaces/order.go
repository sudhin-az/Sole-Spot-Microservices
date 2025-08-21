package interfaceUseCase

import (
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userID int) ([]models.FullOrderDetails, error)
}