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

type ProductHandler struct {
	GRPC_Client interfaces.ProductClient
}

func NewProductHandler(productClient interfaces.ProductClient) *ProductHandler {
	return &ProductHandler{
		GRPC_Client: productClient,
	}
}

// isAdmin checks if the request context is from an admin user
func isAdmin(c *gin.Context) bool {
	// You can check the user role from the context here
	// Assuming you have set the user role in the context
	role, exists := c.Get("user_role")
	return exists && role == "admin"
}

func (pt *ProductHandler) AddProduct(c *gin.Context) {
	// Ensure the user is an admin
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, response.ClientResponse(http.StatusForbidden, "Access denied", nil, nil))
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if err := validator.New().Struct(product); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if product.Stock < 1 {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid stock quantity", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	products, err := pt.GRPC_Client.AddProduct(product)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Could not add the product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Successfully added product", products, nil)
	c.JSON(http.StatusOK, success)
}

func (pt *ProductHandler) ListProducts(c *gin.Context) {
	userID, _ := c.Get("")
	products, err := pt.GRPC_Client.ListProducts(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Could not get the products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Successfully retrieved the products", products, nil)
	c.JSON(http.StatusOK, success)
}

func (pt *ProductHandler) DeleteProduct(c *gin.Context) {
	// Ensure the user is an admin
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, response.ClientResponse(http.StatusForbidden, "Access denied", nil, nil))
		return
	}

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid product ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = pt.GRPC_Client.DeleteProduct(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Could not delete the specified product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Successfully deleted the product", nil, nil)
	c.JSON(http.StatusOK, success)
}

func (pt *ProductHandler) UpdateProducts(c *gin.Context) {
	// Ensure the user is an admin
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, response.ClientResponse(http.StatusForbidden, "Access denied", nil, nil))
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if err := validator.New().Struct(product); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	products, err := pt.GRPC_Client.UpdateProducts(product)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Could not update the product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Successfully updated the product", products, nil)
	c.JSON(http.StatusOK, success)
}
