package client

import (
	"context"
	"fmt"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/config"
	pb "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/pb/product"
	"google.golang.org/grpc"
)

type clientProduct struct {
	client pb.ProductClient
}

func NewProductClient(cfg *config.Config) *clientProduct {
	grpcConnection, err := grpc.Dial(cfg.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewProductClient(grpcConnection)

	return &clientProduct{
		client: grpcClient,
	}
}

func (c *clientProduct) ProductStockMinus(productID, stock int) error {
	_, err := c.client.ProductStockMinus(context.Background(), &pb.ProductStockMinusRequest{
		ID:    int64(productID),
		Stock: int64(stock),
	})
	if err != nil {
		return err
	}
	return nil
}
