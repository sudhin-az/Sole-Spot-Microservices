package repository

import (
	"errors"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/domain"
	interfaceRepository "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(DB *gorm.DB) interfaceRepository.AddressRepository {
	return &AddressRepository{
		DB: DB,
	}
}

// AddAddress adds a new address.
func (ar *AddressRepository) AddAddress(address domain.Address) (domain.Address, error) {
	err := ar.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
		INSERT INTO addresses (user_id, street, city, state, district, zip_code, country)
		VALUES(?, ?, ?, ?, ?, ?, ?)
		`, address.UserID, address.Street, address.City, address.State, address.District, address.ZipCode, address.Country).Error; err != nil {
			return err
		}
		if err := tx.Raw(`
		SELECT id, user_id, street, city, state, district, zip_code, country FROM addresses 
		WHERE user_id = ? AND street = ? AND city = ? AND state = ? AND district = ? AND zip_code = ? AND country = ?
		`, address.UserID, address.Street, address.City, address.State, address.District, address.ZipCode, address.Country).Scan(&address).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return domain.Address{}, err
	}
	return address, nil
}

// GetAddressByID retrieves an address by ID.
func (ar *AddressRepository) GetAddressByID(id uint) (domain.Address, error) {
	var address domain.Address
	err := ar.DB.First(&address, id).Error
	if err != nil {
		return domain.Address{}, errors.New("Record not found")
	}
	return address, nil
}

// UpdateAddress updates an existing address.
func (ar *AddressRepository) UpdateAddress(address domain.Address) (domain.Address, error) {
	err := ar.DB.Model(&address).Where("id = ?", address.ID).Updates(domain.Address{
		Street:   address.Street,
		City:     address.City,
		State:    address.State,
		District: address.District,
		ZipCode:  address.ZipCode,
		Country:  address.Country,
	}).Error

	if err != nil {
		return domain.Address{}, err
	}
	return address, nil
}

// DeleteAddress deletes an address by ID.
func (ar *AddressRepository) DeleteAddress(id uint) error {
	//  Check if the address exists
	var address domain.Address
	err := ar.DB.First(&address, id).Error
	if err != nil {
		return errors.New("Address Not Found")
	}

	err = ar.DB.Delete(&domain.Address{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
