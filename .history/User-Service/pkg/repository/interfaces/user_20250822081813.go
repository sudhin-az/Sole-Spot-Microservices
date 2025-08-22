package interfaces

import (
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/models"
)

type UserRepository interface {
	CheckUserExistsByEmail(email string) (*domain.User, error)
	CheckUserExistsByPhone(phone string) (*domain.User, error)
	UserSignUp(user models.UserSignUp) (models.UserDetails, error)
	UserLogin(user models.UserLogin) (models.UserDetails, error)
}