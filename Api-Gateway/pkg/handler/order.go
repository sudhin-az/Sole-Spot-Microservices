package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/client/interfaces"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/response"
)

type OrderHandler struct {
	GRPC_Client interfaces.OrderClient
}

func NewOrderHandler(client interfaces.OrderClient) *OrderHandler {
	return &OrderHandler{
		GRPC_Client: client,
	}
}

func (or *OrderHandler) OrderItemsFromCart(c *gin.Context) {
	id, _ := c.Get("user_id")
	userID := id.(int)
	var orderFromCart models.OrderFromCart
	if err := c.ShouldBindJSON(&orderFromCart); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	orderSuccessResponse, err := or.GRPC_Client.OrderItemsFromCart(orderFromCart, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully created the order", orderSuccessResponse, nil)
	c.JSON(http.StatusOK, successRes)
}

func (or *OrderHandler) GetOrderDetails(c *gin.Context) {
	id, _ := c.Get("user_id")
	UserID := id.(int)
	orderDetails, err := or.GRPC_Client.GetOrderDetails(UserID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not get the orders", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Full Order Details", orderDetails, nil)
	c.JSON(http.StatusOK, successRes)
}
