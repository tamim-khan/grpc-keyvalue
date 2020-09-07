package main

import (
	"context"
	keyvalue "github.com/tamim-khan/grpc-keyvalue/protos"
	"github.com/tamim-khan/grpc-keyvalue/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var listener *bufconn.Listener

func init() {
	listener = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()
	keyvalue.RegisterKeyValueStoreServer(grpcServer, &server.Server{})
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := keyvalue.NewKeyValueStoreClient(conn)
	resp, err := client.Get(ctx, &keyvalue.GetRequest{Key: "kee"})
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
	if resp.Value != "empty" {
		t.Fatalf("Get failed: Expecting value = \"empty\", Got value = %v", resp.Value)
	}
}
