package interfaceRepository

import "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/models"

type CartRepository interface {
	QuantityOfProductInCart(userId int, productId int) (int, error)
	AddItemIntoCart(userId int, productId int, Quantity int, productPrice float64) error
	TotalPriceForProductInCart(userID int, productID int) (float64, error)
	UpdateCart(quantity int, price float64, userID int, productID int) error
	DisplayCart(userID int) ([]models.Cart, error)
	GetTotalPrice(userID int) (models.CartTotal, error)
	EmptyCart(userID int) error
	ProductExist(userID int, productID int) (bool, error)
	GetAllItemsFromCart(userID int) ([]models.Cart, error)
	DoesCartExist(userID int) (bool, error)
	TotalAmountInCart(userID int) (float64, error)
	UpdateCartAfterOrder(userID int, productID int, quantity float64) error
}
