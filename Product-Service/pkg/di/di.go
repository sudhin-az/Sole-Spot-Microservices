package di

import (
	server "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/api"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/api/services"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/config"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/db"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/repository"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	ProductRepository := repository.NewProductRepository(gormDB)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductServer := services.NewProductServer(ProductUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, ProductServer)

	if err != nil {
		return &server.Server{}, nil
	}
	return grpcServer, nil
}
