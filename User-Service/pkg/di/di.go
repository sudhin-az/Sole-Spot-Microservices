package di

import (
	server "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/api"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/api/services"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/config"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/db"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/repository"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userServiceServer := services.NewUserServer(userUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, userServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}