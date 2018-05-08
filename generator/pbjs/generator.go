package pbjs

import (
	"bytes"
	"text/template"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"go.larrymyers.com/protoc-gen-twirp_typescript/generator"
)

const tmpl = `
import {Message, Method, rpc} from 'protobufjs';
import * as $protobufjs from 'protobufjs';
import axios, {AxiosResponse} from 'axios';
import {twitch} from './service.pb';

const fnNameMatcher = /^\s*function\s+([^(\s]*)\s*/;

const getServiceMethodName = (fn: Method|rpc.ServiceMethod<Message<{}>, Message<{}>>): string => {
    let fnName = "";

    // Function.name isn't cross browser (no IE support), so acknowledge that yes,
    // there is no index signature here, but it will work.

    // @ts-ignore
    if (fn['name']) {
        // @ts-ignore
        fnName = fn['name'];
    }

    const match = fn.toString().match(fnNameMatcher);

    if (match !== null && match.length > 0) {
        fnName = match[1];
    }

    return fnName.slice(0, 1).toUpperCase() + fnName.slice(1);
};

export const createTwirpAdapter = (hostname: string): $protobufjs.RPCImpl => {
    return (method: Method|rpc.ServiceMethod<Message<{}>, Message<{}>>, requestData: Uint8Array, callback: $protobufjs.RPCImplCallback) => {
        axios({
            method: 'POST',
            url: hostname + getServiceMethodName(method),
            headers: {
                'Content-Type': 'application/protobuf'
            },
            // required to get an arraybuffer of the actual size, not the 8192 buffer pool that protobuf.js uses
            // see: https://github.com/dcodeIO/protobuf.js/issues/852
            data: requestData.slice(),
            responseType: 'arraybuffer'
        })
        .then((resp: AxiosResponse<Uint8Array>) => {
            callback(null, new Uint8Array(resp.data));
        })
        .catch((err) => {
            callback(err, null);
        });
    };
};

{{range .Services}}
export const createTwirpHaberdasher = (hostname: string): {{.Package}}.{{.Name}} => {
    return {{.Package}}.{{.Name}}.create(createTwirpAdapter(hostname + '/twirp/{{.Package}}.{{.Name}}/'));
};
{{end}}
`

type service struct {
	Name    string
	Package string
}

type tmplContext struct {
	Services []service
}

func NewGenerator() generator.Generator {
	return &pbjsGenerator{}
}

type pbjsGenerator struct{}

func (g *pbjsGenerator) Generate(d *descriptor.FileDescriptorProto) ([]*plugin.CodeGeneratorResponse_File, error) {
	ctx := tmplContext{}
	pkg := d.GetPackage()

	for _, s := range d.Service {
		service := service{
			Name:    s.GetName(),
			Package: pkg,
		}

		ctx.Services = append(ctx.Services, service)
	}

	t, err := template.New("pbjs_client").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBufferString("")
	err = t.Execute(b, ctx)
	if err != nil {
		return nil, err
	}

	cf := &plugin.CodeGeneratorResponse_File{}
	cf.Name = proto.String("service.twirp.ts")
	cf.Content = proto.String(b.String())

	return []*plugin.CodeGeneratorResponse_File{cf}, nil
}
