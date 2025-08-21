package usecase

import (
	"errors"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/domain"
	interfaceRepository "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/repository/interfaces"
	interfaceUseCase "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/usecase/interfaces"
	"gorm.io/gorm"
)

type AddressUseCase struct {
	AddressRepository interfaceRepository.AddressRepository
}

func NewAddressUseCase(repository interfaceRepository.AddressRepository) interfaceUseCase.AddressUseCase {
	return &AddressUseCase{
		AddressRepository: repository,
	}
}

// AddAddress adds an address to the user's account.
func (ad *AddressUseCase) AddAddress(address domain.Address) (domain.Address, error) {
	newAddress, err := ad.AddressRepository.AddAddress(address)
	if err != nil {
		return domain.Address{}, errors.New("Could not add the address")
	}
	return newAddress, nil
}

// GetAddress retrieves an address by ID through the use case
func (ad *AddressUseCase) GetAddressByID(id uint) (domain.Address, error) {
	address, err := ad.AddressRepository.GetAddressByID(id)
	if err != nil {
		return domain.Address{}, errors.New("Could not retrieve the address: " + err.Error())
	}
	return address, nil
}

// UpdateAddress updates an address through the use case
func (ad *AddressUseCase) UpdateAddress(address domain.Address) (domain.Address, error) {
	updateAddress, err := ad.AddressRepository.UpdateAddress(address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Address{}, domain.ErrAddressNotFound
		}
		return domain.Address{}, errors.New("could not update the address: " + err.Error())
	}
	return updateAddress, nil
}

func (ad *AddressUseCase) DeleteAddress(id uint) error {
	err := ad.AddressRepository.DeleteAddress(id)
	if err != nil {
		return errors.New("could not delete address: " + err.Error())
	}
	return nil
}
