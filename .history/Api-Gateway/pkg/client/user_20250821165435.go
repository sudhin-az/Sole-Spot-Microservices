package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/client/interfaces"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/config"
	pb "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/pb/user"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"
	"google.golang.org/grpc"
)

type UserClient struct {
	Client pb.UserClient
}

func NewUserClient(cfg config.Config) interfaces.UserClient {
	// Establish a secure gRPC connection with proper error handling.
	grpcConnection, err := grpc.Dial(cfg.UserSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error connecting to user service: %v", err)
		return nil
	}

	grpcClient := pb.NewUserClient(grpcConnection)

	return &UserClient{
		Client: grpcClient,
	}
}

func (c *UserClient) UsersSignUp(user models.UserSignUp) (models.TokenUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Client.UserSignUp(ctx, &pb.UserSignUpRequest{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
	})
	if err != nil {
		return models.TokenUser{}, err
	}

	if res.UserDetails == nil {
		return models.TokenUser{}, fmt.Errorf("user details are missing in the response")
	}

	userDetails := models.UserDetails{
		ID:        uint(res.UserDetails.Id),
		Firstname: res.UserDetails.Firstname,
		Lastname:  res.UserDetails.Lastname,
		Email:     res.UserDetails.Email,
		Phone:     res.UserDetails.Phone,
	}
	return models.TokenUser{
		User:        userDetails,
		AccessToken: res.AccessToken,
	}, nil
}

func (c *UserClient) UserLogin(user models.UserLogin) (models.TokenUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Client.UserLogin(ctx, &pb.UserLoginRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return models.TokenUser{}, err
	}
	if res.UserDetails == nil {
		return models.TokenUser{}, fmt.Errorf("user details are missing in the response")
	}
	userDetails := models.UserDetails{
		ID:        uint(res.UserDetails.Id),
		Firstname: res.UserDetails.Firstname,
		Lastname:  res.UserDetails.Lastname,
		Email:     res.UserDetails.Email,
		Phone:     res.UserDetails.Phone,
	}

	return models.TokenUser{
		User:        userDetails,
		AccessToken: res.AccessToken,
	}, nil
}
