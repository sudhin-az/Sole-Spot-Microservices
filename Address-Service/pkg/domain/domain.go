package domain

import "errors"

type Address struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID   uint   `json:"user_id" gorm:"not null"`
	Street   string `json:"street" gorm:"not null"`
	City     string `json:"city" gorm:"not null"`
	State    string `json:"state" gorm:"not null"`
	District string `json:"district" gorm:"not null"`
	ZipCode  string `json:"zip_code" gorm:"not null"`
	Country  string `json:"country" gorm:"not null"`
}

var (
	ErrAddressNotFound    = errors.New("address not found")
	ErrInvalidCredentials = errors.New("Invalid credentials")
)
