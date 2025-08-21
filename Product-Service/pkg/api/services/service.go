package services

import (
	"context"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/models"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/pb"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/usecase"
)

type ProductServer struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServer
}

func NewProductServer(useCase *usecase.ProductUseCase) pb.ProductServer {
	return &ProductServer{
		ProductUseCase: *useCase,
	}
}

func (p *ProductServer) AddProducts(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	model := models.Product{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  uint(req.CategoryID),
		Size:        int(req.Size),
		Stock:       int(req.Stock),
		Price:       float64(req.Price),
	}
	data, err := p.ProductUseCase.AddProducts(model)
	if err != nil {
		return &pb.AddProductResponse{}, err
	}
	return &pb.AddProductResponse{
		ID:          int64(data.ID),
		Name:        data.Name,
		Description: data.Description,
		CategoryID:  int64(data.CategoryID),
		Size:        int64(data.Size),
		Stock:       int64(data.Stock),
		Price:       float32(data.Price),
	}, nil
}

func (p *ProductServer) ListProducts(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductResponse, error) {
	products, err := p.ProductUseCase.ListProducts(int(req.UserID))
	if err != nil {
		return &pb.ListProductResponse{}, err
	}
	var result pb.ListProductResponse
	for _, v := range products {
		result.Details = append(result.Details, &pb.ProductDetails{
			ID:            int64(v.ID),
			Name:          v.Name,
			Description:   v.Description,
			CategoryID:    int64(v.CategoryID),
			Size:          int64(v.Size),
			Stock:         int64(v.Stock),
			Price:         float32(v.Price),
			ProductStatus: v.ProductStatus,
		})
	}
	return &result, nil
}

func (p *ProductServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	product := models.Product{
		ID:          uint(req.ID),
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  uint(req.CategoryID),
		Size:        int(req.Size),
		Stock:       int(req.Stock),
		Price:       float64(req.Price),
	}
	updated, err := p.ProductUseCase.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProductResponse{
		ID:          int64(updated.ID),
		Name:        updated.Name,
		Description: updated.Description,
		CategoryID:  int64(updated.CategoryID),
		Size:        int64(updated.Size),
		Stock:       int64(updated.Stock),
		Price:       float32(updated.Price),
	}, nil
}

func (p *ProductServer) DeleteProducts(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {

	err := p.ProductUseCase.DeleteProducts(int(req.ID))
	if err != nil {
		return &pb.DeleteProductResponse{}, err
	}
	return &pb.DeleteProductResponse{}, nil
}

func (p *ProductServer) GetQuantityFromProductID(ctx context.Context, req *pb.GetQuantityFromProductIDRequest) (*pb.GetQuantityFromProductIDResponse, error) {
	res, err := p.ProductUseCase.GetQuantityFromProductID(int(req.ID))
	if err != nil {
		return &pb.GetQuantityFromProductIDResponse{}, err
	}
	return &pb.GetQuantityFromProductIDResponse{
		Quantity: int64(res),
	}, nil
}

func (p *ProductServer) GetPriceofProductFromID(ctx context.Context, req *pb.GetPriceofProductFromIDRequest) (*pb.GetPriceofProductFromIDResponse, error) {
	res, err := p.ProductUseCase.GetPriceofProductFromID(int(req.ID))
	if err != nil {
		return &pb.GetPriceofProductFromIDResponse{}, err
	}
	return &pb.GetPriceofProductFromIDResponse{
		Price: int64(res),
	}, nil
}

func (p *ProductServer) ProductStockMinus(ctx context.Context, req *pb.ProductStockMinusRequest) (*pb.ProductStockMinusResponse, error) {
	err := p.ProductUseCase.ProductStockMinus(int(req.ID), int(req.Stock))
	if err != nil {
		return &pb.ProductStockMinusResponse{}, err
	}
	return &pb.ProductStockMinusResponse{}, nil
}

func (p *ProductServer) CheckProduct(ctx context.Context, req *pb.CheckProductRequest) (*pb.CheckProductResponse, error) {
	ok, err := p.ProductUseCase.CheckProduct(int(req.ProductID))
	if err != nil {
		return &pb.CheckProductResponse{
			Bool: false,
		}, err
	}
	if !ok {
		return &pb.CheckProductResponse{
			Bool: false,
		}, err
	}
	return &pb.CheckProductResponse{
		Bool: true,
	}, nil
}
