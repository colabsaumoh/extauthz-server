# This is just for notes...need to plug this into the bazel build system.
dep ensure
# Get the grpc proto
go get -u google.golang.org/grpc
# Assume you have the protoc installated and in the $PATH
# SEE: https://grpc.io/docs/quickstart/go.html
# Get the go protoc generator
# Assuming that the includes are moved to /usr/local/incluge/google....
get -u github.com/golang/protobuf/protoc-gen-go
# For now local copies of the gogoproto, google/rpc and validate/
## https://github.com/lyft/protoc-gen-validate
## https://github.com/gogo/protobuf/
## https://github.com/googleapis/googleapis/

# Build protoc for the external auth gRPC server
~/go/bin/protoc -I ./proto -I /usr/local/include ./proto/api/auth/external_auth.proto --go_out=plugins=grpc:ext_auth_v1
~/go/bin/protoc -I ./proto -I /usr/local/include ./proto/api/address.proto --go_out=plugins=grpc:./
~/go/bin/protoc -I ./proto  -I /usr/local/include ./proto/validate/validate.proto --go_out=plugins=grpc:vendor

ln -s ${PWD}/api vendor/api

## Build the server last
go build authzserver.go
