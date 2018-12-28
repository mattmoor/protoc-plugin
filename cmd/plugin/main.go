package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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

		var messages []string
		for _, dp := range fd.MessageType {
			messages = append(messages, fmt.Sprintf(" - %s/%s", fd.GetPackage(), dp.GetName()))
		}
		messageContent := strings.Join(messages, "\n")

		var enums []string
		for _, edp := range fd.EnumType {
			enums = append(enums, fmt.Sprintf(" - %s/%s", fd.GetPackage(), edp.GetName()))
		}
		enumContent := strings.Join(enums, "\n")

		var services []string
		for _, sdp := range fd.Service {
			services = append(services, fmt.Sprintf(" - %s/%s", fd.GetPackage(), sdp.GetName()))
		}
		serviceContent := strings.Join(services, "\n")

		var extensions []string
		for _, fdp := range fd.Extension {
			extensions = append(extensions, fmt.Sprintf(" - %s/%s", fd.GetPackage(), fdp.GetName()))
		}
		extensionContent := strings.Join(extensions, "\n")

		name := fd.GetName() + ".output"
		content := strings.Join([]string{
			"Messages:\n" + messageContent,
			"Enums:\n" + enumContent,
			"Services:\n" + serviceContent,
			"Extensions:\n" + extensionContent,
		}, "\n\n")

		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name:    &name,
			Content: &content,
		})
	}
	return &resp, nil
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Unable to read stdin: %v", err)
	}

	var req plugin_go.CodeGeneratorRequest
	if err := req.Unmarshal(bytes); err != nil {
		log.Fatalf("Unable to unmarshal codegen request: %v", err)
	}

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
