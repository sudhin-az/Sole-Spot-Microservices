package server

import (
	"fmt"
	"net"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/config"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Address-Service/pkg/pb"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, server pb) (*Server, error) {

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	newServer := grpc.NewServer()
	pb.RegisterAddressServer(newServer, server)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50052")
	return c.server.Serve(c.listener)
}
