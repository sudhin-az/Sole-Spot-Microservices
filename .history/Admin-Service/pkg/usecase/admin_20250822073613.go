package usecase

import (
	"errors"

	"github.com/jinzhu/copier"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/helper"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/models"
	interfacesRepository "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/repository/interfaces"
	interfacesUseCase "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/usecase/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	AdminRepository interfacesRepository.AdminRepository
}

func NewAdminUseCase(repository interfacesRepository.AdminRepository) interfacesUseCase.AdminUseCase {
	return &AdminUseCase{
		AdminRepository: repository,
	}
}

func (ad *AdminUseCase) AdminSignup(admin models.AdminSignUp) (*domain.TokenAdmin, error) {
	email, err := ad.AdminRepository.CheckAdminExistsByEmail(admin.Email)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email != nil {
		return &domain.TokenAdmin{}, errors.New("admin with this email is already exists")
	}
	hashPassword, err := helper.PasswordHash(admin.Password)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error in hashing password")
	}
	admin.Password = hashPassword
	adminData, err := ad.AdminRepository.AdminSignUp(admin)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("could not add the user")
	}
	tokenString, err := helper.GenerateToken(adminData)

	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	return &domain.TokenAdmin{
		Admin: adminData,
		Token: tokenString,
	}, nil
}

func (ad *AdminUseCase) AdminLogin(admin models.AdminLogin) (*domain.TokenAdmin, error) {
	email, err := ad.AdminRepository.CheckAdminExistsByEmail(admin.Email)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email == nil {
		return &domain.TokenAdmin{}, errors.New("email doesn't exist")
	}

	adminDetails, err := ad.AdminRepository.AdminLogin(admin)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminDetails.Password), []byte(admin.Password))
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("password not matching")
	}
	var AdminDetailsResponse models.AdminDetailsResponse
	err = copier.Copy(&AdminDetailsResponse, &adminDetails)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateToken(AdminDetailsResponse)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	return &domain.TokenAdmin{
		Admin: AdminDetailsResponse,
		Token: tokenString,
	}, nil
}
