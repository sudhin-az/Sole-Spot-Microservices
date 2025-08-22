package interfaces

import "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"

type UserClient interface {
	UserSignUp(user models.UserSignUp) (models.TokenUser, error)
	UserLogin(user models.UserLogin) (models.TokenUser, error)
}
