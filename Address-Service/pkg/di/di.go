package di

import (
	server "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/api"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/api/services"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/config"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/db"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/repository"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	addressRepository := repository.NewAddressRepository(gormDB)
	addressUseCase := usecase.NewAddressUseCase(addressRepository)
	addressServiceServer := services.NewAddressServer(addressUseCase)

	grpcServer, err := server.NewGRPCServer(cfg, addressServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}