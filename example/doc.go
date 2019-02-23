package mattmoor

//go:generate go generate github.com/mattmoor/protoc-plugin/proto
//go:generate go install github.com/mattmoor/protoc-plugin/cmd/krpc-yaml
//go:generate go install github.com/mattmoor/protoc-plugin/cmd/krpc-go
//go:generate mkdir -p ./build
//go:generate ${HOME}/protoc-3.6.1/bin/protoc -I${HOME}/protoc-3.6.1/include -I${GOPATH}/src -I. --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --go_out=./build hello.proto
//go:generate ${HOME}/protoc-3.6.1/bin/protoc -I${HOME}/protoc-3.6.1/include -I${GOPATH}/src -I. --plugin=protoc-gen-krpc-yaml=${GOPATH}/bin/krpc-yaml --krpc-yaml_opt=docker.io/mattmoor/grpc-ping-go               --krpc-yaml_out=./build hello.proto
//go:generate ${HOME}/protoc-3.6.1/bin/protoc -I${HOME}/protoc-3.6.1/include -I${GOPATH}/src -I. --plugin=protoc-gen-krpc-go=${GOPATH}/bin/krpc-go     --krpc-go_opt=github.com/mattmoor/protoc-plugin/example/build --krpc-go_out=./build hello.proto
