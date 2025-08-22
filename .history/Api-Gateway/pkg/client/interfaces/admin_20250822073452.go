package interfaces

import "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"

type AdminClient interface {
	AdminSignp(adminDetails models.AdminSignUp) (models.TokenAdmin, error)
	AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error)
}
