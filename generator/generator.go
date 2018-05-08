package generator

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"go.larrymyers.com/protoc-gen-twirp_typescript/generator/minimal"
	"go.larrymyers.com/protoc-gen-twirp_typescript/generator/pbjs"
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

type Generator interface {
	Generate(d *descriptor.FileDescriptorProto) ([]*plugin.CodeGeneratorResponse_File, error)
}

func NewGenerator(p Params) Generator {
	lib, ok := p["library"]
	if ok && lib == "pbjs" {
		return pbjs.NewGenerator()
	}

	return minimal.NewGenerator(p)
}
