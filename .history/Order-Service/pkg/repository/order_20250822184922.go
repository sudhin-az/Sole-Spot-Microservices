package repository

import (
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/models"
	interfaceRepository "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaceRepository.OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (or *OrderRepository) AddressExist(orderBody models.OrderIncoming) (bool, error) {
	
	var count int
	err := or.DB.Raw("SELECT COUNT(*) FROM addresses WHERE user_id = ? AND id = ?", orderBody.UserID, orderBody.AddressID).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (or *OrderRepository) PaymentExist(orderBody models.OrderIncoming) (bool, error) {
	var count int
	err := or.DB.Raw("SELECT COUNT(*) FROM payment_methods WHERE id = ?", orderBody.PaymentID).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (or *OrderRepository) PaymentStatus(orderID int) (string, error) {
	var PaymentStatus string
	err := or.DB.Raw("SELECT payment_status FROM orders WHERE id = ?", orderID).Scan(&PaymentStatus).Error
	if err != nil {
		return "", err
	}
	return PaymentStatus, nil
}

func (or *OrderRepository) OrderItems(ob models.OrderIncoming, price float64) (int, error) {
	shipment_status := "Pending"
	PaymentMethod := "Paid"
	var id int
	query := `INSERT INTO orders (created_at, user_id, address_id, shipment_status, payment_status, final_price)
	VALUES (NOW(), ?, ?, ?, ?, ?)
	RETURNING id`
	or.DB.Raw(query, ob.UserID, ob.AddressID, shipment_status, PaymentMethod, price).Scan(&id)
	return id, nil
}

func (or *OrderRepository) AddOrderProducts(order_id int, cart []models.Cart) error {
	query := `INSERT INTO order_items(order_id, product_id, quantity, total_price) VALUES(?, ?, ?, ?)`
	for _, v := range cart {
		
		if err := or.DB.Exec(query, order_id, v.ProductID, v.Quantity, v.TotalPrice).Error; err != nil {
			return err
		}
	}
	return nil
}

func (or *OrderRepository) GetBriefOrderDetails(orderID int) (domain.OrderSuccessResponse, error) {
	var OrderSuccessResponse domain.OrderSuccessResponse
	err := or.DB.Raw(`SELECT id as order_id, shipment_status FROM orders WHERE id = ?`, orderID).Scan(&OrderSuccessResponse).Error
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	return OrderSuccessResponse, nil
}

func (or *OrderRepository) GetOrderDetails(userID int) ([]models.FullOrderDetails, error) {
	var orders []models.FullOrderDetails
	
	//First fetch orders
	err := or.DB.Table("orders").
	Select("id as order_id, final_price, shipment_status, payment_status").
	Where("user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	var fullOrders []models.FullOrderDetails
	for _, order := range orders {
		var products []models.OrderProductDetails

		//Fetch products for each order
		err := or.DB.Table("order_items").
		Select("product_id, quantity, total_price").
		Where("order_id = ?", order.or)
	}
}