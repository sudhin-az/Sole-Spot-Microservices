package domain

import (
	"errors"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/models"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"lastame" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Phone     string `json:"phone" gorm:"not null"`
}

type TokenUser struct {
	User         models.UserDetails
	AccessToken  string
	RefreshToken string
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)
