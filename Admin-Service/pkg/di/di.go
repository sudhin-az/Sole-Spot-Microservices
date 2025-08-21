package di

import (
	server "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/api"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/api/services"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/config"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/db"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/repository"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	AdminRepository := repository.NewAdminRepository(gormDB)
	AdminUseCase := usecase.NewAdminUseCase(AdminRepository)
	AdminServiceServer := services.NewAdminServer(AdminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, AdminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}