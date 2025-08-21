package client

import (
	"context"
	"fmt"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/client/interfaces"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/config"
	pb "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/pb/order"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Api-Gateway/pkg/utils/models"
	"google.golang.org/grpc"
)

type OrderClient struct {
	Client pb.OrderClient
}

func NewOrderClient(cfg config.Config) interfaces.OrderClient {
	grpcConnection, err := grpc.Dial(cfg.OrderSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewOrderClient(grpcConnection)

	return &OrderClient{
		Client: grpcClient,
	}
}

func (c *OrderClient) OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (models.OrderSuccessResponse, error) {
	model := &pb.OrderItem{
		AddressID: int64(orderFromCart.AddressID),
		PaymentID: int64(orderFromCart.PaymentID),
	}
	res, err := c.Client.OrderItemsFromCart(context.Background(), &pb.OrderItemsFromCartRequest{
		OrderFromCart: model,
		UserID: int64(userID),
	})
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}
	if res.Error != "" {
		return models.OrderSuccessResponse{}, err
	}
	return models.OrderSuccessResponse{
		OrderID: uint(res.OrderID),
		ShipmentStatus: "delivered",  // Set the default value for ShipmentStatus
	}, nil
}

func (c *OrderClient) GetOrderDetails(userID int) ([]models.FullOrderDetails, error) {
	res, err := c.Client.GetOrderDetails(context.Background(), &pb.GetOrderDetailsRequest{
		UserID: int64(userID),
	})
	if err != nil {
		return []models.FullOrderDetails{}, err
	}
	if res.Error != "" {
		return []models.FullOrderDetails{}, err
	}
	var result []models.FullOrderDetails

	for _, v := range res.Details {
		orderDetails := models.OrderDetails{
			OrderId: int(v.OrderDetails.OrderID),
			FinalPrice: float64(v.OrderDetails.Price),
			ShipmentStatus: "delivered",
			PaymentStatus: v.OrderDetails.Paymentstatus,
		}

		var orderProductDetails []models.OrderProductDetails
		for _, product := range v.OrderProductDetails {
			orderProduct := models.OrderProductDetails{
				ProductID: uint(product.ProductID),
				Quantity: int(product.Quantity),
				TotalPrice: float64(product.Price),
			}
			orderProductDetails = append(orderProductDetails, orderProduct)
		}
		fullOrderDetails := models.FullOrderDetails{
			OrderDetails: orderDetails,
			OrderProductDetails: orderProductDetails,
		}
		result = append(result, fullOrderDetails)
	}
	return result, nil
}
