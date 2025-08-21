package server

import (
	"fmt"
	"net"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/config"
	pb "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Order-Service/pkg/pb/order"
	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.OrderServer) (*Server, error) {

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	newServer := grpc.NewServer()
	pb.RegisterOrderServer(newServer, server)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50056")
	return c.server.Serve(c.listener)
}
