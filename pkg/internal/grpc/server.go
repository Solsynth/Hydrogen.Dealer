package grpc

import (
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	health "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	proto.UnimplementedServiceDirectoryServer
	proto.UnimplementedAuthServer

	srv *grpc.Server
}

func NewServer() *Server {
	server := &Server{
		srv: grpc.NewServer(),
	}

	proto.RegisterServiceDirectoryServer(server.srv, &Server{})
	proto.RegisterAuthServer(server.srv, &Server{})
	health.RegisterHealthServer(server.srv, &Server{})

	reflection.Register(server.srv)

	return server
}

func (v *Server) Listen() error {
	listener, err := net.Listen("tcp", viper.GetString("grpc_bind"))
	if err != nil {
		return err
	}

	return v.srv.Serve(listener)
}
