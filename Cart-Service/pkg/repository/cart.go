package repository

import (
	"errors"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		DB: db,
	}
}

func (cr *CartRepository) QuantityOfProductInCart(userId int, productId int) (int, error) {
	var productQty int
	err := cr.DB.Raw("SELECT quantity FROM carts WHERE user_id = ? AND product_id = ?", userId, productId).Scan(&productQty).Error
	if err != nil {
		return 0, err
	}
	return productQty, nil
}

func (cr *CartRepository) AddItemIntoCart(userId int, productId int, Quantity int, productPrice float64) error {
	err := cr.DB.Exec(`INSERT INTO carts(user_id, product_id, quantity, total_price) VALUES(?, ?, ?, ?)
	`, userId, productId, Quantity, productPrice).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *CartRepository) TotalPriceForProductInCart(userId, productId int) (float64, error) {
	var totalPrice float64

	err := cr.DB.Raw("SELECT SUM(total_price) as total_price FROM carts WHERE user_id = ? AND product_id = ?", userId, productId).Scan(&totalPrice).Error
	if err != nil {
		return 0.0, err
	}
	return totalPrice, nil
}

func (cr *CartRepository) UpdateCart(quantity int, price int, userId int, productId int) error {
	err := cr.DB.Exec(`UPDATE carts SET quantity = ?, total_price = ? WHERE user_id = ? AND product_id = ?`,
		quantity, price, userId, productId).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *CartRepository) DisplayCart(userID int) ([]models.Cart, error) {
	var count int
	err := cr.DB.Raw("SELECT COUNT(*) FROM carts WHERE user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return []models.Cart{}, err
	}
	if count == 0 {
		return []models.Cart{}, nil
	}
	var cartResponse []models.Cart

	if err := cr.DB.Raw("SELECT * FROM carts WHERE user_id = ?", userID).Scan(&cartResponse).Error; err != nil {
		return []models.Cart{}, err
	}
	return cartResponse, nil
}

func (cr *CartRepository) GetTotalPrice(userID int) (models.CartTotal, error) {
	var cartTotal models.CartTotal
	err := cr.DB.Raw("SELECT COALESCE(SUM(total_price), 0) FROM carts WHERE user_id = ?", userID).Scan(&cartTotal.TotalPrice).Error
	if err != nil {
		return models.CartTotal{}, nil
	}
	err = cr.DB.Raw("SELECT COALESCE(SUM(total_price), 0) FROM carts WHERE user_id = ?", userID).Scan(&cartTotal.FinalPrice).Error
	if err != nil {
		return models.CartTotal{}, err
	}
	return cartTotal, nil
}

func (cr *CartRepository) EmptyCart(userID int) error {
	err := cr.DB.Exec("DELETE FROM carts WHERE user_id = ?", userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *CartRepository) ProductExist(userID int, productID int) (bool, error) {
	var count int
	if err := cr.DB.Raw("SELECT count(*) FROM carts WHERE user_id = ? AND product_id = ?", userID, productID).Scan(&count).Error; err != nil {
		return false, err
	}
	if count == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (cr *CartRepository) GetAllItemsFromCart(userID int) ([]models.Cart, error) {
	var count int
	var cartResponse []models.Cart
	err := cr.DB.Raw("SELECT COUNT(*) FROM carts WHERE user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return []models.Cart{}, err
	}
	if count == 0 {
		return []models.Cart{}, nil
	}
	err = cr.DB.Raw("SELECT carts.user_id, carts.product_id, carts.quantity, carts.total_price FROM carts WHERE user_id = ?", userID).First(&cartResponse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if len(cartResponse) == 0 {
				return []models.Cart{}, nil
			}
			return []models.Cart{}, err
		}
		return []models.Cart{}, err
	}
	return cartResponse, nil
}

func (cr *CartRepository) DoesCartExist(userID int) (bool, error) {
	var exist bool
	err := cr.DB.Raw("SELECT exists(SELECT 1 FROM carts WHERE user_id = ?)", userID).Scan(&exist).Error
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (cr *CartRepository) TotalAmountInCart(userID int) (float64, error) {
	var price float64
	err := cr.DB.Raw("SELECT SUM(total_price) FROM carts WHERE user_id = ?", userID).Scan(&price).Error
	if err != nil {
		return 0.0, err
	}
	return price, nil
}

func (cr *CartRepository) UpdateCartAfterOrder(userID int, productID int, quantity float64) error {
	err := cr.DB.Raw("DELETE FROM carts WHERE user_id = ? AND product_id = ?", userID, productID).Error
	if err != nil {
		return err
	}
	return nil
}
