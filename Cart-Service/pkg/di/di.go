package di

import (
	server "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/api"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/api/services"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/client"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/config"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/db"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/repository"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	cartRepository := repository.NewCartRepository(gormDB)
	productClient := client.NewProductClient(&cfg)
	cartUseCase := usecase.NewCartUseCase(*cartRepository, productClient)

	cartServiceServer := services.NewCartServer(cartUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, cartServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil
}