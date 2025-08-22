package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/client/interfaces"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/response"
)

type UserHandler struct {
	GRPC_Client interfaces.UserClient
}

func NewUserHandler(UserClient interfaces.UserClient) *UserHandler {
	return &UserHandler{
		GRPC_Client: UserClient,
	}
}

// UserSignup handles user signup by making a call to the gRPC service
func (ur *UserHandler) UserSignUp(c *gin.Context) {
	var signupDetail models.UserSignUp
	if err := c.ShouldBindJSON(&signupDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Validate the struct fields
	err := validator.New().Struct(signupDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Call the gRPC client for user signup
	user, err := ur.GRPC_Client.UserSignUp(signupDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Error during signup", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	success := response.ClientResponse(http.StatusCreated, "User successfully signed up", user, nil)
	c.JSON(http.StatusCreated, success)
}

// UserLogin handles user login by calling the gRPC service
func (ur *UserHandler) Userlogin(c *gin.Context) {
	var userLoginDetail models.UserLogin
	if err := c.ShouldBindJSON(&userLoginDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Validate the login details
	err := validator.New().Struct(userLoginDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Call the gRPC client for user login
	user, err := ur.GRPC_Client.UserLogin(userLoginDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Error during login", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "User successfully logged in", user, nil)
	c.JSON(http.StatusOK, success)
}
