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

export const createTwirpHaberdasher = (hostname: string): twitch.twirp.example.Haberdasher => {
    return twitch.twirp.example.Haberdasher.create(createTwirpAdapter(hostname + '/twirp/twitch.twirp.example.Haberdasher/'));
};
