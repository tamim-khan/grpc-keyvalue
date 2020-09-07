.PHONY: proto

PROTO_DIR=./protos

proto:
	protoc -I ${PROTO_DIR} ${PROTO_DIR}/*.proto --go_out=plugins=grpc:${PROTO_DIR}
