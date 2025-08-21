package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID         int     `json:"user_id" gorm:"not null"`
	AddressID      uint    `json:"address_id" gorm:"not null"`
	ShipmentStatus string  `json:"shipment_status" gorm:"default:'pending'"`
	PaymentStatus  string  `json:"payment_status" gorm:"default:'not paid'"`
	FinalPrice     float64 `json:"final_price"`
	Approval       bool    `json:"approval" gorm:"default:false"`
}

type OrderItem struct {
	ID         uint    `json:"id" gorm:"primaryKey;not null"`
	OrderID    uint    `json:"order_id"`
	ProductID  uint    `json:"product_id"`
	Quantity   float64 `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type OrderSuccessResponse struct {
	OrderID        uint   `json:"order_id"`
	ShipmentStatus string `json:"shipment_status" gorm:"default:delivered"`
}
type Address struct {
	Id        int    `json:"id" gorm:"unique;not null"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	District  string `json:"district" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}
type PaymentMethod struct {
	ID           uint   `json:"id" gorm:"primaryKey;not null"`
	Payment_Name string `json:"payment_name" gorm:"unique; not null"`
}
