package services

import (
	"context"
	"errors"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/pb"
	interfaceUseCase "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/usecase/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddressServer struct {
	AddressUseCase interfaceUseCase.AddressUseCase
	pb.UnimplementedAddressServer
}

func NewAddressServer(useCase interfaceUseCase.AddressUseCase) pb.AddressServer {
	return &AddressServer{
		AddressUseCase: useCase,
	}
}

// AddAddress adds a new address for a user.
func (s *AddressServer) AddAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error) {
	addressDetails := domain.Address{
		UserID:   uint(req.UserId),
		Street:   req.Street,
		City:     req.City,
		State:    req.State,
		District: req.District,
		ZipCode:  req.ZipCode,
		Country:  req.Country,
	}

	newAddress, err := s.AddressUseCase.AddAddress(addressDetails)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add address: %v", err)
	}
	return &pb.AddAddressResponse{
		Status: 201,
		AddressDetails: &pb.AddressDetails{
			Id:       uint64(newAddress.ID),
			UserId:   uint64(newAddress.UserID),
			Street:   newAddress.Street,
			City:     newAddress.City,
			State:    newAddress.State,
			District: newAddress.District,
			ZipCode:  newAddress.ZipCode,
			Country:  newAddress.Country,
		},
	}, nil
}

// GetAddress retrieves a user's address by its ID.
func (s *AddressServer) GetAddressByID(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressResponse, error) {
	address, err := s.AddressUseCase.GetAddressByID(uint(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get address: %v", err)
	}

	return &pb.GetAddressResponse{
		Status: 200,
		AddressDetails: &pb.AddressDetails{
			Id:       uint64(address.ID),
			UserId:   uint64(address.UserID),
			Street:   address.Street,
			City:     address.City,
			State:    address.State,
			District: address.District,
			ZipCode:  address.ZipCode,
			Country:  address.Country,
		},
	}, nil
}

// UpdateAddress updates an existing address.
func (s *AddressServer) UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressResponse, error) {
	addressDetails := domain.Address{
		ID: ,
		UserID: uint(req.UserId),
		Street:   req.Street,
		City:     req.City,
		State:    req.State,
		District: req.District,
		ZipCode:  req.ZipCode,
		Country:  req.Country,
	}

	updateAddress, err := s.AddressUseCase.UpdateAddress(addressDetails)
	if err != nil {
		if errors.Is(err, domain.ErrAddressNotFound) {
			return nil, status.Errorf(codes.NotFound, "address with ID %d not found", req.UserId)
		}
		return nil, status.Errorf(codes.Internal, "failed to update address: %v", err)
	}
	return &pb.UpdateAddressResponse{
		Status: 200,
		AddressDetails: &pb.AddressDetails{
			Id:       uint64(updateAddress.ID),
			UserId:   uint64(updateAddress.UserID),
			Street:   updateAddress.Street,
			City:     updateAddress.City,
			State:    updateAddress.State,
			District: updateAddress.District,
			ZipCode:  updateAddress.ZipCode,
			Country:  updateAddress.Country,
		},
	}, nil
}

// DeleteAddress deletes a user's address by its ID.
func (s *AddressServer) DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressResponse, error) {
	err := s.AddressUseCase.DeleteAddress(uint(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete address: %v", err)
	}

	return &pb.DeleteAddressResponse{
		Status: 200,
	}, nil
}
