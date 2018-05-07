package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"go.larrymyers.com/protoc-gen-twirp_typescript/generator"
	"go.larrymyers.com/protoc-gen-twirp_typescript/generator/minimal"
)

func main() {
	req := readRequest(os.Stdin)
	writeResponse(os.Stdout, generate(req))
}

func readRequest(r io.Reader) *plugin.CodeGeneratorRequest {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	req := new(plugin.CodeGeneratorRequest)
	if err = proto.Unmarshal(data, req); err != nil {
		panic(err)
	}

	if len(req.FileToGenerate) == 0 {
		panic(err)
	}

	return req
}

func generate(in *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	resp := &plugin.CodeGeneratorResponse{}

	for _, f := range in.GetProtoFile() {
		// skip google/protobuf/timestamp, we don't do any special serialization for jsonpb.
		if *f.Name == "google/protobuf/timestamp.proto" {
			continue
		}

		cf, err := minimal.Generate(f)
		if err != nil {
			resp.Error = proto.String(err.Error())
			return resp
		}

		resp.File = append(resp.File, cf)
	}

	resp.File = append(resp.File, minimal.RuntimeLibrary())

	params := generator.GetParameters(in)

	if pkgName, ok := params["package_name"]; ok {
		idx, err := minimal.CreatePackageIndex(resp.File)
		if err != nil {
			resp.Error = proto.String(err.Error())
			return resp
		}

		resp.File = append(resp.File, idx)
		resp.File = append(resp.File, minimal.CreateTSConfig())
		resp.File = append(resp.File, minimal.CreatePackageJSON(pkgName))
	}

	return resp
}

func writeResponse(w io.Writer, resp *plugin.CodeGeneratorResponse) {
	data, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(data)
	if err != nil {

	}
}
