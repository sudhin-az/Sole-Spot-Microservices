package usecase

import (
	"errors"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/helper"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/models"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/repository/interfaces"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/copier"
)

type userUseCase struct {
	userRepository interfaces.UserRepository
}

func NewUserUseCase(repository interfaces.UserRepository) *userUseCase {
	return &userUseCase{
		userRepository: repository,
	}
}

// UsersSignUp handles user registration.
func (uc *userUseCase) UserSignUp(user models.UserSignUp) (domain.TokenUser, error) {
	//Check if email exists
	existingUser, err := uc.userRepository.CheckUserExistsByEmail(user.Email)
	if err != nil {
		return domain.TokenUser{}, errors.New("server error during email check")
	} else if existingUser != nil {
		return domain.TokenUser{}, errors.New("user with this email already exists")
	}

	//Check if phone exists
	existingPhone, err := uc.userRepository.CheckUserExistsByPhone(user.Phone)
	if err != nil {
		return domain.TokenUser{}, errors.New("server error during phone check")
	} else if existingPhone != nil {
		return domain.TokenUser{}, errors.New("user with this phone already exists")
	}

	//Hash password
	hahedPassword, err := helper.PasswordHash(user.Password)
	if err != nil {
		return domain.TokenUser{}, errors.New("failed to hash password")
	}
	user.Password = hahedPassword

	//Create user
	userData, err := uc.userRepository.UserSignUp(user)
	if err != nil {
		return domain.TokenUser{}, errors.New("failed to sign up user")
	}

	//Generate Tokens
	accessToken, err := helper.GenerateAccessToken(userData)
	if err != nil {
		return domain.TokenUser{}, errors.New("failed to generate access token")
	}
	refreshToken, err := helper.GenerateRefreshToken(userData)
	if err != nil {
		return domain.TokenUser{}, errors.New("failed to generate refresh token")
	}

	return domain.TokenUser{
		User: userData,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *userUseCase) UserLogin(user models.UserLogin) (domain.TokenUser, error) {
	email, err := uc.userRepository.CheckUserExistsByEmail(user.Email)
	if err != nil {
		return domain.TokenUser{}, errors.New("error with server")
	}
	if email == nil {
		return domain.TokenUser{}, errors.New("email doesn't exist")
	}
	userDetails, err := uc.userRepository.UserLogin(user)
	if err != nil {
		return domain.TokenUser{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password))
	if err != nil {
		return domain.TokenUser{}, errors.New("password not matching")
	}
	// var userdetails models.UserDetails
	// err = copier.Copy(&userDetails, &userdetails)
	// if err != nil {
	// 	return domain.TokenUser{}, err
	// }
	accessToken, err := helper.GenerateAccessToken(userDetails)
	if err != nil {
		return domain.TokenUser{}, errors.New("couldn't create accessToken due to internal error")
	}
	refreshToken, err := helper.GenerateRefreshToken(userdetails)
	if err != nil {
		return domain.TokenUser{}, errors.New("counldn't create refreshtoken due to internal error")
	}
	return domain.TokenUser{
		User: userdetails,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}
