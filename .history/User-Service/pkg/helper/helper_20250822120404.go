package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthUserClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func PasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
	hash := string(hashedPassword)
	return hash, nil
}

func GenerateTokenUsers(userID int, userEmail string, expirationTime time.Time) (string, error) {
	claims := &AuthUserClaims{
		Id:    userID,
		Email: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("123456789"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateAccessToken(user models.UserDetails) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	tokenString, err := GenerateTokenUsers(int(user.ID), user.Email, expirationTime)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func GenerateRefreshToken(user models.UserDetails) (string, error) {
	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	tokenString, err := GenerateTokenUsers(int(user.ID), user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}

func PasswordHashing(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
	hash := string(hashedPassword)
	return hash, nil
}

func GetTokenFromHeader(header string) string {
	if len(header) > 7 && header[:7] == "Bearer " {
		return header[7:]
	}
	return header
}

func ExtractUserIDFromToken(tokenString string) (int, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthUserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte("123456789"), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*AuthUserClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid token claims")
	}
	return claims.Id, claims.Email, nil
}
