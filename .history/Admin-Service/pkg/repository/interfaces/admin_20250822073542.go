package interfacesRepository

import (
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/models"
)

type AdminRepository interface {
	AdminSignp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	AdminLogin(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)
}
