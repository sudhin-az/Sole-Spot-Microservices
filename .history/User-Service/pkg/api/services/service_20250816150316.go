package services

import (
	"context"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/models"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/pb"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/User-Service/pkg/usecase/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	userUseCase interfaces.UserUseCase
	pb.UnimplementedUserServer
}

func NewUserServer(usecase interfaces.UserUseCase) pb.UserServer {
	return &UserServer{
		userUseCase: usecase,
	}
}

// UserSignUp handles the signup process for new users.
func (s *UserServer) UserSignUp(ctx context.Context, userSignUpDetails *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	userCreateDetails := models.UserSignUp{
		FirstName: userSignUpDetails.Firstname,
		LastName:  userSignUpDetails.Lastname,
		Email:     userSignUpDetails.Email,
		Phone:     userSignUpDetails.Phone,
		Password:  userSignUpDetails.Password,
	}
	data, err := s.userUseCase.UserSignUp(userCreateDetails)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to sign up user: %v", err)
	}

	userDetails := &pb.UserDetails{
		Id:        uint64(data.User.ID),
		Firstname: data.User.FirstName,
		Lastname:  data.User.LastName,
		Email:     data.User.Email,
		Phone:     data.User.Phone,
	}

	return &pb.UserSignUpResponse{
		Status:      200,
		UserDetails: userDetails,
		AccessToken: data.AccessToken,
	}, nil
}

// UserLogin handles the login process for users.
func (s *UserServer) UserLogin(ctx context.Context, loginDetails *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	login := models.UserLogin{
		Email:    loginDetails.Email,
		Password: loginDetails.Password,
	}

	data, err := s.userUseCase.UserLogin(login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to login: %v", err)
	}

	userDetails := &pb.UserDetails{
		Id:        uint64(data.User.ID),
		Firstname: data.User.FirstName,
		Lastname:  data.User.LastName,
		Email:     data.User.Email,
		Phone:     data.User.Phone,
	}

	return &pb.UserLoginResponse{
		Status:      200,
		UserDetails: userDetails,
		AccessToken: data.AccessToken,
	}, nil
}
