package interfacesUseCase

import (
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/models"
)

type AdminUseCase interface {
	AdminSignup(adminDetails models.AdminSignUp) (*domain.TokenAdmin, error)
	AdminLogin(adminDetails models.AdminLogin) (*domain.TokenAdmin, error)
}