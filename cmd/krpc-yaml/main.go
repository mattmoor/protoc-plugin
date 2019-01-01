package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	kpb "github.com/mattmoor/protoc-plugin/proto/krpc"
)

const domain = "mattmoor.io"

// // TODO(mattmoor): Use codegen when for real, but this will do for prototyping.
// // TODO(mattmoor): We should use a single struct extension to annotate all of
// // the stuff we need.
// var (
//     serviceAccount = stringOption("service_account", 50001)
//     // TODO(mattmoor): Environment Variables (can't do without codegen)
//     // TODO(mattmoor): ContainerConcurrency
//     // TODO(mattmoor): TimeoutSeconds
//     // TODO(mattmoor): Resources
// )

// func stringOption(name string, number int32) *proto.ExtensionDesc {
//      return &proto.ExtensionDesc{
// 	ExtendedType: (*descriptor.MethodOptions)(nil),
// 	ExtensionType: (*string)(nil),
// 	Field: number,
// 	Name: fmt.Sprintf("proto.%s", name),
//         Tag: fmt.Sprintf("bytes,%d,opt,name=%s", number, name),
//     }
// }

func method(importpath string, fd *descriptor.FileDescriptorProto, sdp *descriptor.ServiceDescriptorProto, mdp *descriptor.MethodDescriptorProto,
            resp *plugin_go.CodeGeneratorResponse) (string, string, error) {
        // directory := fd.GetPackage() + "/" + sdp.GetName() + "/" + mdp.GetName()

	// TODO(mattmoor): If we build-in the `ko`-like functionality, then this need not be a full import path.
	// importpath := "github.com/fooxbar/baz/outpath/" + directory

	serviceName := strings.ToLower(sdp.GetName() + "-" + mdp.GetName())

	var options kpb.Options
	addr, err := proto.GetExtension(mdp.Options, kpb.E_Options)
	if err == nil {
	   options = *(addr.(*kpb.Options))
	}

	yamlContent := fmt.Sprintf(`
apiVersion: serving.knative.dev/v1alpha1
kind: Service
metadata:
  name: %s
spec:
  revisionTemplate:
    spec:
      serviceAccountName: %s
      containerConcurrency: %d
      container:
        image: %s
        env:
`, serviceName, options.GetServiceAccount(), options.ContainerConcurrency, importpath)

        for _, kv := range options.Env {
	  yamlContent = yamlContent + fmt.Sprintf(`
        - name: %s
          value: %s        
`, kv.GetName(), kv.GetValue())
 	}

	return serviceName, yamlContent, nil
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

		type routingRule struct{
		    Method string
                    Path string
                    ServiceName string
		    YAMLContent string
		}

		var rules []routingRule
		for _, sdp := range fd.Service {
		        for _, mdp := range sdp.Method {
			        serviceName, yamlContent, err := method(request.GetParameter(), fd, sdp, mdp, &resp)
				if err != nil {
				   return nil, err
				}
				rules = append(rules, routingRule{
				  Method: "POST",
				  Path: fmt.Sprintf("/%s.%s/%s", fd.GetPackage(), sdp.GetName(), mdp.GetName()),
				  ServiceName: serviceName,
				  YAMLContent: yamlContent,
				})
			}
		}

		// Based on the accumulated rules generate the dispatch yaml.
		content := fmt.Sprintf(`
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: grpc-gateway
spec:
  gateways:
  - knative-shared-gateway.knative-serving.svc.cluster.local
  hosts:
  - %s
  http:
`, domain)

		for _, rr := range rules {
			content = rr.YAMLContent + "\n---\n" + content + fmt.Sprintf(`
  - match:
    - uri:
        prefix: %s
    rewrite:
      authority: %s.%s.svc.cluster.local
    route:
      - destination:
          host: knative-ingressgateway.istio-system.svc.cluster.local
        weight: 100
`, rr.Path, rr.ServiceName, "default")
   	        }

		name := "service.yaml"
		// log.Printf("Content[%q] = %s", name, content)

		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name:    &name,
			Content: &content,
		})
	}
	return &resp, nil
}

func main() {
        // proto.RegisterExtension(serviceAccount)

        // log.Printf("Args: %v", os.Args)
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
