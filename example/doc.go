package mattmoor

//go:generate go generate github.com/mattmoor/protoc-plugin/proto
//go:generate go install github.com/mattmoor/protoc-plugin/cmd/krpc-yaml
//go:generate go install github.com/mattmoor/protoc-plugin/cmd/krpc-yaml
//go:generate mkdir -p ./build
//go:generate protoc -I${GOPATH}/src -I. --plugin=protoc-gen-krpc-yaml=${GOPATH}/bin/krpc-yaml --krpc-yaml_opt=github.com/blurg/bleh/cmd/thing --krpc-yaml_out=./build hello.proto
//go:generate protoc -I${GOPATH}/src -I. --plugin=protoc-gen-krpc-go=${GOPATH}/bin/krpc-go --krpc-go_out=./build hello.proto
