package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/handler"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

// NewServerHTTP initializes the server with routes and handlers
func NewServerHTTP(adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, userHandler *handler.UserHandler, cartHandler *handler.CartHandler, addressHandler *handler.AddressHandler, orderHandler *handler.OrderHandler) *ServerHTTP {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery()) // Add recovery middleware to handle panics

	//Public routes
	router.POST("/admin/login", adminHandler.AdminLogin)
	router.POST("/admin/signup", adminHandler.AdminSignup)
	router.POST("/user/signup", userHandler.UserSignup)
	router.POST("/user/login", userHandler.Userlogin)

	// Admin routes
	adminRoutes := router.Group("/")
	adminRoutes.Use(middleware.AdminAuthMiddleware())
	{
		adminRoutes.POST("/product", productHandler.AddProducts)
		adminRoutes.PUT("/product", productHandler.UpdateProducts)
		adminRoutes.DELETE("/product", productHandler.DeleteProduct)
		adminRoutes.GET("/product", productHandler.ListProducts)
	}

	//User routes
	userRoutes := router.Group("/")
	userRoutes.Use(middleware.UserAuthMiddleware())
	{
		userRoutes.POST("/cart", cartHandler.AddToCart)
		userRoutes.GET("/cart", cartHandler.GetCart)
		userRoutes.POST("/order", orderHandler.OrderItemsFromCart)
		userRoutes.GET("/order", orderHandler.GetOrderDetails)

		// Address routes
		userRoutes.POST("/address", addressHandler.AddAddress)
		userRoutes.GET("/address/:id", addressHandler.GetAddress)
		userRoutes.PUT("/address", addressHandler.UpdateAddress)
		userRoutes.DELETE("/address/:id", addressHandler.DeleteAddress)
	}

	return &ServerHTTP{engine: router}
}

// Start runs the server on the specified port
func (s *ServerHTTP) Start() {
	log.Println("Starting server on :8080")
	err := s.engine.Run(":8080")
	if err != nil {
		log.Fatalf("Error while starting the server: %v", err)
	}
}
