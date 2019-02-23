package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
)

func doit(request *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	codegen := make(map[string]struct{})
	for _, file := range request.FileToGenerate {
		codegen[file] = struct{}{}
	}

	var resp plugin_go.CodeGeneratorResponse
	for _, fd := range request.ProtoFile {
		if _, ok := codegen[fd.GetName()]; !ok {
			continue
		}

		for _, sdp := range fd.Service {
			// TODO(mattmoor): Add pb import

			// TODO(mattmoor): Generate a class definition.
			// e.g. type ServiceName struct {}

			// TODO(mattmoor): Generate a service registration.
			// e.g. svcName := &ServiceName{}
			//      pb.RegisterServiceNameServer(grpcServer, sycName)

			for _, mdp := range sdp.Method {
				// TODO(mattmoor): Generate methods off of the class definition.
				// e.g. non-streaming:
				//   func (s *ServiceName) MethodName(ctx context.Context, req *pb.MethodRequest) (*pb.MethodResponse, error) {
				//      // Implemented by the user.
				//      return someimportpath.MethodName(ctx, req)
				//   }
				//
				// e.g. streaming:
				//   func (s *ServiceName) MethodName(stream pb.ServiceName_MethodNameServer) error {
				//      // Implemented by the user.
				//      return someimportpath.MethodName(stream)
				//   }

				// "// TODO: " + fd.GetPackage() + "/" + sdp.GetName() + "/" + mdp.GetName(), nil

				// TODO(mattmoor): Consider optionally retrieving the import path for this
				// method from an annotation?
			}
		}

		mainName := "main.go"
		mainContent := fmt.Sprintf(
			mainTemplate,
			importpath,
			serviceDefinitions,
			registerServices,
		)

		log.Printf("Content[%q] = %s", mainName, mainContent)
		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name:    &mainName,
			Content: &mainContent,
		})
	}
	return &resp, nil
}

func main() {
	log.Printf("Args: %v", os.Args)
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Unable to read stdin: %v", err)
	}

	var req plugin_go.CodeGeneratorRequest
	if err := req.Unmarshal(bytes); err != nil {
		log.Fatalf("Unable to unmarshal codegen request: %v", err)
	}

	// log.Printf("Parameter: %q", req.GetParameter())
	resp, err := doit(&req)
	if err != nil {
		// Return plugin errors as an error response instead of a
		// non-zero exit code.
		s := err.Error()
		resp = &plugin_go.CodeGeneratorResponse{
			Error: &s,
		}
	}

	result, err := resp.Marshal(nil, true)
	if err != nil {
		log.Fatalf("Unable to marshal codegen response: %v", err)
	}
	os.Stdout.Write(result)
}

const (
	mainTemplate = `
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"os"

	// Inject import paths for relevant definitions.
	%s

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Inject service structs and methods here.
%s

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%%d", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %%v", err)
	}

	// The grpcServer is currently configured to serve h2c traffic by default.
	// To configure credentials or encryption, see: https://grpc.io/docs/guides/auth.html#go
	grpcServer := grpc.NewServer()

	// Inject service registration here.
	%s

	grpcServer.Serve(lis)
}
`
)
