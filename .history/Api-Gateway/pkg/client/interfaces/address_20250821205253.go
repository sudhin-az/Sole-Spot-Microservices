package interfaces

import "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"

type AddressClient interface {
	AddAddress(address models.Address) (models.Address, error)
	GetAddress(id uint) (models.Address, error)
	UpdateAddress(address models.Address) (models.Address, error)
	DeleteAddress(id int) error
}