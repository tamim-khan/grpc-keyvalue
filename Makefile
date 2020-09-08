.PHONY: test build run proto

PROTO_DIR=./protos

test:
	go test -v

build:
	go build -o build/grpc-keyvalue

run: build
	./build/grpc-keyvalue

proto:
	protoc -I ${PROTO_DIR} ${PROTO_DIR}/*.proto --go_out=plugins=grpc:${PROTO_DIR}
