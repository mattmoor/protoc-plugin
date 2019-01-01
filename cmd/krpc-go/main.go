package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func method(fd *descriptor.FileDescriptorProto, sdp *descriptor.ServiceDescriptorProto, mdp *descriptor.MethodDescriptorProto) (string, error) {
        return "// TODO: " + fd.GetPackage() + "/" + sdp.GetName() + "/" + mdp.GetName(), nil
}

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

		mainParts := []string{
			  "package main",
			  `import "log"`,
			  "func main() {",
		}

		for _, sdp := range fd.Service {
		        for _, mdp := range sdp.Method {
			        content, err := method(fd, sdp, mdp)
				if err != nil {
				   return nil, err
				}
				mainParts = append(mainParts, content)
			}
		}

		mainParts = append(mainParts,
			  `	log.Fatal("NYI")`,
			  "}",
		)
		mainName := "main.go"
		mainContent := strings.Join(mainParts, "\n")

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

	log.Printf("Parameter: %q", req.GetParameter())
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
