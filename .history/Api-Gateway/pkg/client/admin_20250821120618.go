package client

import (
	"context"
	"fmt"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/client/interfaces"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/config"
	pb "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/pb/admin"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"
	"google.golang.org/grpc"
)

type AdminClient struct {
	Client pb.AdminClient
}

func NewAdminClient(cfg config.Config) interfaces.AdminClient {

	grpcConnection, err := grpc.Dial(cfg.AdminSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewAdminClient(grpcConnection)

	return &AdminClient{
		Client: grpcClient,
	}
}

func (ad *AdminClient) AdminSignUp(adminDetails models.AdminSignUp) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminSignup(context.Background(), &pb.AdminSignupRequest{
		Firstname: adminDetails.Firstname,
		Lastname: adminDetails.Lastname,
		Email: adminDetails.Email,
		Password: adminDetails.Password,
	})
	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID: uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname: admin.AdminDetails.Lastname,
			Email: admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}

func (ad *AdminClient) AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminLogin(context.Background(), &pb.AdminLoginRequest{
		Email: adminDetails.Email,
		Password: adminDetails.Password,
	})
	if err != nil {
		return models.TokenAdmin{}, err
	}

	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID: uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname: admin.AdminDetails.Lastname,
			Email: admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}
