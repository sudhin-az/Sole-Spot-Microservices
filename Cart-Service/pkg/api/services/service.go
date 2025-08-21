package services

import (
	"context"

	pb "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/pb/cart"
	interfaceUsecase "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Cart-Service/pkg/usecase/interfaces"
)

type CartServer struct {
	CartUseCase interfaceUsecase.CartUseCase
	pb.UnimplementedCartServer
}

func NewCartServer(cartUseCase interfaceUsecase.CartUseCase) pb.CartServer {
	return &CartServer{
		CartUseCase: cartUseCase,
	}
}

func (c *CartServer) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	res, err := c.CartUseCase.AddToCart(int(req.ProductID), int(req.UserID), int(req.Quantity))
	if err != nil {
		return &pb.AddToCartResponse{}, err
	}
	var result pb.AddToCartResponse
	var cartDetails []*pb.CartDetails
	for _, v := range res.Cart {
		details := &pb.CartDetails{
			ProductID: int64(v.ProductID),
			Quantity: float32(v.Quantity),
			TotalPrice: float32(v.TotalPrice),
		}
		cartDetails = append(cartDetails, details)
	}
	result.Price = float32(res.TotalPrice)
	result.Cart = cartDetails
	return &result, nil
}

func (c *CartServer) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	res, err := c.CartUseCase.DisplayCart(int(req.UserID))
	if err != nil {
		return &pb.GetCartResponse{}, err
	}
	var result pb.GetCartResponse
	var cartDetails []*pb.CartDetails
	for _, v := range res.Cart {
		details := &pb.CartDetails{
			ProductID: int64(v.ProductID),
			Quantity: float32(v.Quantity),
			TotalPrice: float32(v.TotalPrice),
		}
		cartDetails = append(cartDetails, details)
	}
	result.Price = float32(res.TotalPrice)
	result.Cart = cartDetails
	return &result, nil
}

func (c *CartServer) DoesCartExist(ctx context.Context, req *pb.DoesCartExistRequest) (*pb.DoesCartExistResponse, error) {
	exists, err := c.CartUseCase.DoesCartExist(int(req.UserID))
	if err != nil {
		return &pb.DoesCartExistResponse{}, err
	}
	return &pb.DoesCartExistResponse{
		Data: exists,
	}, nil
}

func (c *CartServer) GetAllItemsFromCart(ctx context.Context, req *pb.GetAllItemsFromCartRequest) (*pb.GetAllItemsFromCartResponse, error) {
	res, err := c.CartUseCase.GetAllItemsFromCart(int(req.UserID))
	if err != nil {
		return &pb.GetAllItemsFromCartResponse{}, err
	}

	var cartDetails []*pb.CartDetails
	for _, cartItem := range res {
		cartDetail := &pb.CartDetails{
			ProductID: int64(cartItem.ProductID),
			Quantity: float32(cartItem.Quantity),
			TotalPrice: float32(cartItem.TotalPrice),
		}
		cartDetails = append(cartDetails, cartDetail)
	}
	return &pb.GetAllItemsFromCartResponse{
		Cart: cartDetails,
	}, nil
}

func (c *CartServer) TotalAmountInCart(ctx context.Context, req *pb.TotalAmountInCartRequest) (*pb.TotalAmountInCartResponse, error) {
	res, err := c.CartUseCase.TotalAmountInCart(int(req.UserID))
	if err != nil {
		return &pb.TotalAmountInCartResponse{}, err
	}
	return &pb.TotalAmountInCartResponse{
		Data: float32(res),
	}, nil
}

func (c *CartServer) UpdateCartAfterOrder(ctx context.Context, req *pb.UpdateCartAfterOrderRequest) (*pb.UpdateCartAfterOrderResponse, error) {
	err := c.CartUseCase.UpdateCartAfterOrder(int(req.UserID), int(req.ProductID), float64(req.Quantity))
	if err != nil {
		return &pb.UpdateCartAfterOrderResponse{}, err
	}
	return &pb.UpdateCartAfterOrderResponse{}, nil
}