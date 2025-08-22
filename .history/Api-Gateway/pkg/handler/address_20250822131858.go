package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/client/interfaces"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/response"
)

type AddressHandler struct {
	GRPC_Client interfaces.AddressClient
}

func NewAddressHandler(AddressClient interfaces.AddressClient) *AddressHandler {
	return &AddressHandler{
		GRPC_Client: AddressClient,
	}
}

// AddAddress handles the addition of a new address
func (ad *AddressHandler) AddAddress(c *gin.Context) {
	var addressDetail models.Address
	if err := c.ShouldBindJSON(&addressDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		errs := response.ClientResponse(http.sua)
	}

	// Validate the address details
	err := validator.New().Struct(addressDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Call the gRPC client to add the address
	address, err := ad.GRPC_Client.AddAddress(addressDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Error adding address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	success := response.ClientResponse(http.StatusCreated, "Address successfully added", address, nil)
	c.JSON(http.StatusCreated, success)
}

// GetAddress handles retrieval of an address by ID
func (ad *AddressHandler) GetAddress(c *gin.Context) {
	id := c.Param("id")
	addressID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid address ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	address, err := ad.GRPC_Client.GetAddress(uint(addressID))
	if err != nil {
		errs := response.ClientResponse(http.StatusNotFound, "Address not found", nil, err.Error())
		c.JSON(http.StatusNotFound, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Address retrieved successfully", address, nil)
	c.JSON(http.StatusOK, success)
}

func (ad *AddressHandler) UpdateAddress(c *gin.Context) {
	var addressDetail models.Address
	if err := c.ShouldBindJSON(&addressDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Validate the address details
	err := validator.New().Struct(addressDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	address, err := ad.GRPC_Client.UpdateAddress(addressDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Error updating address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Address successfully updated", address, nil)
	c.JSON(http.StatusOK, success)
}

func (ad *AddressHandler) DeleteAddress(c *gin.Context) {
	id := c.Param("id")
	addressID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid address ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ad.GRPC_Client.DeleteAddress(int(addressID))
	if err != nil {
		errs := response.ClientResponse(http.StatusNotFound, "Error deleting address", nil, err.Error())
		c.JSON(http.StatusNotFound, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Address successfully deleted", nil, nil)
	c.JSON(http.StatusOK, success)
}
