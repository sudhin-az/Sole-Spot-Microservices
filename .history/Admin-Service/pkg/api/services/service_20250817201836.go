package services

import (
	"context"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/models"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/pb"
	interfacesUseCase "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Admin-Service/pkg/usecase/interfaces"
)

type AdminServer struct {
	AdminUseCase interfacesUseCase.AdminUseCase
	pb.UnimplementedAdminServer
}

func NewAdminServer(useCase interfacesUseCase.AdminUseCase) *AdminServer {
	return &AdminServer{
		AdminUseCase: useCase,
	}
}

func (ad *AdminServer) AdminSignUp(ctx context.Context, req *pb.AdminSignupRequest) (*pb.AdminSignupResponse, error) {
	adminSignUp := models.AdminSignUp{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
	}

	res, err := ad.AdminUseCase.AdminSignUp(adminSignUp)
	if err != nil {
		return &pb.AdminSignupResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(res.Admin.ID),
		Firstname: res.Admin.Firstname,
		Lastname:  res.Admin.Lastname,
		Email:     res.Admin.Email,
	}
	return &pb.AdminSignupResponse{
		Status:       201,
		AdminDetails: adminDetails,
		Token:        res.Token,
	}, nil
}

func (ad *AdminServer) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	adminLogin := models.AdminLogin{
		Email:    req.Email,
		Password: req.Password,
	}
	admin, err := ad.AdminUseCase.AdminLogin(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(admin.Admin.ID),
		Firstname: admin.Admin.Firstname,
		Lastname:  admin.Admin.Lastname,
		Email:     admin.Admin.Email,
	}
	return &pb.AdminLoginResponse{
		Status:       200,
		AdminDetails: adminDetails,
		Token:        admin.Token,
	}, nil
}
