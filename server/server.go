package server

import (
	"context"
	"github.com/tamim-khan/grpc-keyvalue/database"
	keyvalue "github.com/tamim-khan/grpc-keyvalue/protos"
)

type Server struct {
	db *database.Database
}

func New() (*Server, error) {
	db, err := database.Start("file.db")
	if err != nil {
		return nil, err
	}
	return &Server{db: db}, nil
}

func (s *Server) Shutdown() error {
	return s.db.Close()
}

func (s *Server) Get(ctx context.Context, req *keyvalue.GetRequest) (*keyvalue.GetResponse, error) {
	if value, err := s.db.Get(req.Key); err != nil {
		return nil, err
	} else {
		return &keyvalue.GetResponse{
			Value: *value,
		}, nil
	}
}

func (s *Server) Set(ctx context.Context, req *keyvalue.SetRequest) (*keyvalue.SetResponse, error) {
	if err := s.db.Set(req.Key, req.Value); err != nil {
		return nil, err
	} else {
		return &keyvalue.SetResponse{}, nil
	}
}

func (s *Server) Delete(ctx context.Context, req *keyvalue.DeleteRequest) (*keyvalue.DeleteResponse, error) {
	if err := s.db.Delete(req.Key); err != nil {
		return nil, err
	} else {
		return &keyvalue.DeleteResponse{
			Status: keyvalue.DeleteResponse_DELETED,
		}, nil
	}
}
