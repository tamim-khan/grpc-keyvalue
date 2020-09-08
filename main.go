package main

import (
	"fmt"
	keyvalue "github.com/tamim-khan/grpc-keyvalue/protos"
	"github.com/tamim-khan/grpc-keyvalue/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	address := ":5005"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	keyValueServer, err := server.New()
	if err != nil {
		panic(err)
	}
	defer keyValueServer.Shutdown()
	keyvalue.RegisterKeyValueStoreServer(grpcServer, keyValueServer)

	reflection.Register(grpcServer)

	fmt.Printf("Listening on %s\n", address)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
