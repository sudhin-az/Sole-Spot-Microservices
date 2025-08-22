package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/client/interfaces"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/config"
	pb "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/pb/address"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"
	"google.golang.org/grpc"
)

type AddressClient struct {
	Client pb.AddressClient
}

func NewAddressClient(cfg config.Config) interfaces.AddressClient {
	// Establish a secure gRPC connection with proper error handling.
	grpcConnection, err := grpc.Dial(cfg.AddressSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error connecting to address service: %v", err)
		return nil
	}

	grpcClient := pb.NewAddressClient(grpcConnection)

	return &AddressClient{
		Client: grpcClient,
	}
}

func (c *AddressClient) AddAddress(address models.Address) (models.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Client.AddAddress(ctx, &pb.AddAddressRequest{
		UserId:   uint64(address.UserID),
		Street:   address.Street,
		City:     address.City,
		State:    address.State,
		District: address.District,
		ZipCode:  address.ZipCode,
		Country:  address.Country,
	})

	if err != nil {
		return models.Address{}, err
	}

	if res.AddressDetails == nil {
		return models.Address{}, fmt.Errorf("address details are missing in the response")
	}

	return models.Address{
		ID:       uint(res.AddressDetails.Id),
		UserID:   uint(res.AddressDetails.UserId),
		Street:   res.AddressDetails.Street,
		City:     res.AddressDetails.City,
		State:    res.AddressDetails.State,
		District: res.AddressDetails.District,
		ZipCode:  res.AddressDetails.ZipCode,
		Country:  res.AddressDetails.Country,
	}, nil
}

func (c *AddressClient) GetAddressByID(id uint) (models.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Client.GetAddressByID(ctx, &pb.GetAddressRequest{
		Id: uint64(id),
	})
	if err != nil {
		return models.Address{}, err
	}
	if res.Status != 200 || res.AddressDetails == nil {
		return models.Address{}, fmt.Errorf("address not found with ID %d", id)
	}

	return models.Address{
		ID:       uint(res.AddressDetails.Id),
		UserID:   uint(res.AddressDetails.UserId),
		Street:   res.AddressDetails.Street,
		City:     res.AddressDetails.City,
		State:    res.AddressDetails.State,
		District: res.AddressDetails.District,
		ZipCode:  res.AddressDetails.ZipCode,
		Country:  res.AddressDetails.Country,
	}, nil
}

func (c *AddressClient) UpdateAddress(address models.Address) (models.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Client.UpdateAddress(ctx, &pb.UpdateAddressRequest{
		Id:       uint64(address.ID),
		UserId:   uint64(address.UserID),
		Street:   address.Street,
		City:     address.City,
		State:    address.State,
		District: address.District,
		ZipCode:  address.ZipCode,
		Country:  address.Country,
	})
	if err != nil {
		return models.Address{}, err
	}

	if res.Status != 200 || res.AddressDetails == nil {
		return models.Address{}, fmt.Errorf("failed to update address with ID %d", address.ID)
	}

	return models.Address{
		ID:       uint(res.AddressDetails.Id),
		UserID:   uint(res.AddressDetails.UserId),
		Street:   res.AddressDetails.Street,
		City:     res.AddressDetails.City,
		State:    res.AddressDetails.State,
		District: res.AddressDetails.District,
		ZipCode:  res.AddressDetails.ZipCode,
		Country:  res.AddressDetails.Country,
	}, nil
}

func (c *AddressClient) DeleteAddress(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Client.DeleteAddress(ctx, &pb.DeleteAddressRequest{
		Id: uint64(id),
	})
	if err != nil {
		return err
	}

	if res.Status != 200 {
		return fmt.Errorf("failed to delete address with ID %d", id)
	}

	return nil
}
