package services

import (
	"context"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/models"
	pb "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/pb/order"
	interfaceUseCase "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/usecase/interfaces"
)

type OrderServer struct {
	UseCase interfaceUseCase.OrderUseCase
	pb.UnimplementedOrderServer
}

func NewOrderServer(useCase interfaceUseCase.OrderUseCase) pb.OrderServer {
	return &OrderServer{
		UseCase: useCase,
	}
}

func (or *OrderServer) OrderItemsFromCart(ctx context.Context, req *pb.OrderItemsFromCartRequest) (*pb.OrderItemsFromCartResponse, error) {
	model := &models.OrderFromCart{
		AddressID: uint(req.OrderFromCart.AddressID),
		PaymentID: uint(req.OrderFromCart.AddressID),
	}
	userID := req.UserID
	result, err := or.UseCase.OrderItemsFromCart(*model, int(userID))
	if err != nil {
		return &pb.OrderItemsFromCartResponse{}, err
	}
	return &pb.OrderItemsFromCartResponse{
		OrderID: int64(result.OrderID),
		ShipmentStatus: result.ShipmentStatus,
	}, nil
}

func (or *OrderServer) GetOrderDetails(ctx context.Context, req *pb.GetOrderDetailsRequest) (*pb.GetOrderDetailsResponse, error) {
	details, err := or.UseCase.GetOrderDetails(int(req.UserID))
	if err != nil {
		return nil, err
	}

	var result pb.GetOrderDetailsResponse

	for _, v := range details {
		var orderDetails pb.OrderDetails
		orderDetails.OrderID = int64(v.OrderDetails.OrderId)
		orderDetails.Price = float32(v.OrderDetails.FinalPrice)
		orderDetails.Shipmentstatus = v.OrderDetails.ShipmentStatus
		orderDetails.Paymentstatus = v.OrderDetails.PaymentStatus

		var orderProductDetails []*pb.OrderProductDetails
		for _, product := range v.OrderProductDetails {
			orderProduct := &pb.OrderProductDetails{
				ProductID: int64(product.ProductID),
				Quantity: int64(product.Quantity),
				Price: float32(product.TotalPrice),
			}
			orderProductDetails = append(orderProductDetails, orderProduct)
		}

		fullOrderDetails := &pb.FullOrderDetails{
			OrderDetails: &orderDetails,
			OrderProductDetails: orderProductDetails,
		}
		result.Details = append(result.Details, fullOrderDetails)
	}
	return &result, nil
}