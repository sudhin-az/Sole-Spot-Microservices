package interfaces

import (
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserSignUp) (domain.TokenUser, error)
	UserLogin(user models.UserLogin) (domain.TokenUser, error)
}