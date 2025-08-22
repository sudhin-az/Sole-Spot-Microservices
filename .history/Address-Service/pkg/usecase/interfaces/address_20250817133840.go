package interfaceUseCase

import "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/domain"

type AddressUseCase interface {
	AddAddress(address domain.Address) (domain.Address, error)
	GetAddressByID(id uint) (domain.Address, error)
	UpdateAddress(address domain.Address) (domain.Address, error)
	DeleteAddress(id uint) error
}
