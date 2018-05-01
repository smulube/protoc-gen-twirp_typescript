import {Message, rpc} from 'protobufjs';
import * as $protobufjs from 'protobufjs';
import axios, {AxiosResponse} from 'axios';

const fnNameMatcher = /^\s*function\s+([^\(\s]*)\s*/;

const getServiceMethodName = (fn: Function): string => {
    let fnName = "";

    if (fn['name']) {
        fnName = fn['name'];
    }

    const match = fn.toString().match(fnNameMatcher);

    if (match !== null && match.length > 0) {
        fnName = match[1];
    }

    return fnName.slice(0, 1).toUpperCase() + fnName.slice(1);
};

export const createTwirpAdapter = (hostname: string): $protobufjs.RPCImpl => {
    return (method: rpc.ServiceMethod<Message<{}>,Message<{}>>, requestData: Uint8Array, callback: $protobufjs.RPCImplCallback) => {
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
