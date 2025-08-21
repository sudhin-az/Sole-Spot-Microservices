package client

import (
	"context"
	"fmt"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/config"
	pb "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/pb/product"
	"google.golang.org/grpc"
)

type clientProduct struct {
	Client pb.ProductClient
}

func NewProductClient(c *config.Config) *clientProduct {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("could not connect: ", err)
	}

	pbClient := pb.NewProductClient(cc)

	return &clientProduct{
		Client: pbClient,
	}
}

func (c *clientProduct) GetQuantityFromProductID(id int) (int, error) {
	res, err := c.Client.GetQuantityFromProductID(context.Background(), &pb.GetQuantityFromProductIDRequest{
		ID: int64(id),
	})
	if err != nil {
		return 0, err
	}
	quantity := res.Quantity
	return int(quantity), nil
}

func (c *clientProduct) GetPriceOfProductFromID(id int) (float64, error) {
	res, err := c.Client.GetPriceofProductFromID(context.Background(), &pb.GetPriceofProductFromIDRequest{
		ID: int64(id),
	})
	if err != nil {
		return 0, err
	}
	price := res.Price
	return float64(price), nil
}
func (c *clientProduct) ProductStockMinus(productID, stock int) error {
	_, err := c.Client.ProductStockMinus(context.Background(), &pb.ProductStockMinusRequest{
		ID: int64(productID),
		Stock: int64(stock),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *clientProduct) CheckProduct(productID int) (bool, error) {
	ok, err := c.Client.CheckProduct(context.Background(), &pb.CheckProductRequest{
		ProductID: int64(productID),
	})
	if err != nil {
		return false, err
	}
	return ok.Bool, nil
}