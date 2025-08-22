package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
)

// AuthUserClaims represents the claims for user tokens
type AuthUserClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

//PasswordHash hashes the password
func PasswordHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("internal server error")
	}
	return string(hashPassword), nil
}

// GenerateTokenUsers generates a JWT token for the user
func GenerateTokenUsers(userID int, userEmail string, expirationTime time.Time) (string, error) {
	claims := &AuthUserClaims{
		Id: userID,
		Email: userEmail,
		Role: "user", // Add roles or other claims as needed
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(),
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("123456789")) // Use a secure key in production
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// GenerateAccessToken generates an access token for the user with a short expiration time
func GenerateAccessToken(user models.UserDetailsResponse) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	return GenerateTokenUsers(user.Id, user.Email, expirationTime)
}

// GenerateRefreshToken generates a refresh token for the user with a long expiration time
func GenerateRefreshToken(user models.UserDetails) (string, error) {
	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	return GenerateTokenUsers(int(user.ID), user.Email, expirationTime)
}

// CompareHashAndPassword compares a hashed password with a plain password
func CompareHashAndPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}