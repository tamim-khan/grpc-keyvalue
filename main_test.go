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
	keyValueServer, err := server.New()
	if err != nil {
		panic(err)
	}
	keyvalue.RegisterKeyValueStoreServer(grpcServer, keyValueServer)
	go func() {
		defer keyValueServer.Shutdown()
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func TestSetGetDelete(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := keyvalue.NewKeyValueStoreClient(conn)

	// Set
	_, err = client.Set(ctx, &keyvalue.SetRequest{
		Key:   "test_key",
		Value: "test_value",
	})
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// Get
	getResp, err := client.Get(ctx, &keyvalue.GetRequest{Key: "test_key"})
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if getResp.Value != "test_value" {
		t.Fatalf("Get failed: Expecting value = \"test_value\", Got value = %v", getResp.Value)
	}

	// Delete
	_, err = client.Delete(ctx, &keyvalue.DeleteRequest{Key: "test_key"})
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	_, err = client.Get(ctx, &keyvalue.GetRequest{Key: "test_key"})
	if err == nil {
		t.Fatalf("Key present after delete")
	}
}
