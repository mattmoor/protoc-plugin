package proto

//go:generate go get -u github.com/golang/protobuf/protoc-gen-go
//go:generate mkdir -p ./krpc
//go:generate ${HOME}/protoc-3.6.1/bin/protoc -I${HOME}/protoc-3.6.1/include -I. --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --go_out=./krpc krpc.proto
