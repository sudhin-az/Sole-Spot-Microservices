package repository

import (
	"errors"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/models"
	interfacesRepository "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfacesRepository.AdminRepository {
	return &AdminRepository{
		DB: db,
	}
}

func (ad *AdminRepository) AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error) {
	var model models.AdminDetailsResponse

	if err := ad.DB.Raw(`INSERT INTO admins(firstname, lastname, email, password)
	VALUES (?, ?, ?, ?) RETURNING id, firstname, lastname, email`, adminDetails.Firstname, adminDetails.Lastname, adminDetails.Email, adminDetails.Password).Scan(&model).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}
	return model, nil
}

func (ad *AdminRepository) CheckAdminExistsByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	res := ad.DB.Where("email = ?", email).First(&admin)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Admin{}, res.Error
	}
	return &admin, nil
}

func (ad *AdminRepository) AdminLogin(admin models.AdminLogin) (models.AdminSignUp, error) {
	var admins models.AdminSignUp

	err := ad.DB.Raw("SELECT * FROM admins WHERE email = ?", admin.Email).Scan(&admins).Error
	if err != nil {
		return models.AdminSignUp{}, errors.New("error checking user details")
	}
	return admins, nil
}
