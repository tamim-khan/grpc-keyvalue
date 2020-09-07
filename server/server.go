package server

import (
	"context"
	keyvalue "github.com/tamim-khan/grpc-keyvalue/protos"
)

type Server struct{}

func (s *Server) Get(ctx context.Context, req *keyvalue.GetRequest) (*keyvalue.GetResponse, error) {
	return &keyvalue.GetResponse{
		Value: "empty",
	}, nil
}

func (s *Server) Set(ctx context.Context, req *keyvalue.SetRequest) (*keyvalue.SetResponse, error) {
	return &keyvalue.SetResponse{}, nil
}

func (s *Server) Delete(ctx context.Context, req *keyvalue.DeleteRequest) (*keyvalue.DeleteResponse, error) {
	return &keyvalue.DeleteResponse{
		Status: keyvalue.DeleteResponse_NOT_FOUND,
	}, nil
}
