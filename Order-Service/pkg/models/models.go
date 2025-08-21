package models

type OrderDetails struct {
	OrderId        int     `json:"order_id"`
	FinalPrice     float64 `json:"final_price"`
	ShipmentStatus string  `json:"shipment_status"`
	PaymentStatus  string  `json:"payment_status"`
}

type OrderProductDetails struct {
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type FullOrderDetails struct {
	OrderDetails        OrderDetails
	OrderProductDetails []OrderProductDetails
}
type OrderProducts struct {
	ProductId string `json:"id"`
	Stock     int    `json:"stock"`
}

type AddedOrderProductDetails struct {
	UserID          int `json:"user_id"`
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_id"`
}
type OrderResponse struct {
	AddedOrderProductDetails AddedOrderProductDetails
	OrderDetails             OrderDetails
}

type OrderFromCart struct {
	AddressID uint `json:"address_id" binding:"required"`
	PaymentID uint `json:"payment_id" binding:"required"`
}

type OrderIncoming struct {
	UserID    int `json:"user_id"`
	PaymentID int `json:"payment_id"`
	AddressID int `json:"address_id"`
}

type Cart struct {
	ProductID  uint    `json:"product_id"`
	Quantity   float64 `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
