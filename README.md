# grpc-keyvalue

[![Build Status](https://travis-ci.com/tamim-khan/grpc-keyvalue.svg?token=eqFaMMfDXJ7bNqqNdA8B&branch=master)](https://travis-ci.com/tamim-khan/grpc-keyvalue)

## Running the server

### curl for gRPC
```shell script
brew install grpcurl
```

### list methods
```shell script
grpcurl --plaintext localhost:5005 list
```

### Call a method
```shell script
grpcurl --plaintext -d '{ "key":"kee" }' localhost:5005 keyvalue.KeyValueStore.Get
```

## Supported Methods
- [x] `GET`
- [x] `SET`
- [x] `DELETE`

See `proto` definition for more details
