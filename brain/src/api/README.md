# The Block Services
Here we store all of the services to support the backend for The Block. We are using gRPCs to maintain our APIs.

## Handling Proto Files
To create proto files to represent request and response messages please follow the steps in whatever order is needed:
- create a .proto file defining package, go_package, syntax=proto3, and define request and response messages.
- using `protoc --go_out=. --go-grpc_out=. example.proto` generate proto files in the same relative path
    - please use a standard naming convention such as [ package ]pb
- define your endpoints under an internal folder

Please feel free to take a look at services already made

## Test Suite
To run go unit tests under brain:
```
go test ./...
```
### Generating Mocks
To generate mocks for unit testing, please run:
```
mockgen -destination=mock_destination.go -package=mockpackage [directory] [package interface you want to mock]
```
example: `mockgen -destination=mocks/mock_authdatastore.go -package=mockauthdatastore TheBlock/src/api/auth/internal/datastore AuthDatastore`

This must be updated when datastore interface is updated.
