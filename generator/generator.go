package generator

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"go.larrymyers.com/protoc-gen-twirp_typescript/generator/minimal"
)

type Params map[string]string

func GetParameters(in *plugin.CodeGeneratorRequest) Params {
	params := make(Params)

	if in.Parameter == nil {
		return params
	}

	pairs := strings.Split(*in.Parameter, ",")

	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		params[kv[0]] = kv[1]
	}

	return params
}

type Generator = func(d *descriptor.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)

func NewGenerator(p Params) Generator {
	return minimal.Generate
}
