package di

import (
	server "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/api"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/api/services"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/client"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/config"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/db"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/repository"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	orderRepository := repository.NewOrderRepository(gormDB)
	cartClient := client.NewCartClient(&cfg)
	productClient := client.NewProductClient(&cfg)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartClient, productClient)

	orderServiceServer := services.NewOrderServer(orderUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, orderServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil
}
