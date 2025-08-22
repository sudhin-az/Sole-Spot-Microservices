package repository

import (
	"errors"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/models"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

// CheckUserExistsByEmail checks if a user exists by email.
func (u *userRepository) CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := u.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CheckUserExistsByPhone checks if a user exists by phone.
func (u *userRepository) CheckUserExistsByPhone(phone string) (*domain.User, error) {
	var user domain.User

	err := u.DB.Where("phone = ?", phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UserSignUp signs up a new user.
func (u *userRepository) UserSignUp(user models.UserSignUp) (models.UserDetails, error) {
	var signUpDetail models.UserDetails

	err := u.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
		INSERT INTO users (first_name, lastname, email, password, phone)
		VALUES(?, ?, ?, ?, ?)
		`, user.FirstName, user.LastName, user.Email, user.Password, user.Phone).Error; err != nil {
			return err
		}

		if err := tx.Raw(`
		SELECT id, firstname, lastname, email, phone FROM users WHERE email = ?
		`, user.Email).Scan(&signUpDetail).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return models.UserDetails{}, err
	}
	return signUpDetail, nil
}

func (u *userRepository) UserLogin(user models.UserLogin) (models.UserDetail, error) {
	var userDetails models.UserDetail
	query := `SELECT id, firstname, lastname, email, password, phone FROM users WHERE email = ?`
	err := u.DB.Raw(query, user.Email).Scan(&userDetails).Error
	if err != nil {
		return models.UserDetail{}, errors.New("user not found")
	}
	return userDetails, nil
}
