package proto

//go:generate go get -u github.com/golang/protobuf/protoc-gen-go
//go:generate mkdir -p ./krpc
//go:generate protoc --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --go_out=./krpc krpc.proto
