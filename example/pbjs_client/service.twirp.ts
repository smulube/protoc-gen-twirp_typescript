import {Message, rpc} from 'protobufjs';
import * as $protobufjs from 'protobufjs';

export const createTwirpClient = (hostname: string): $protobufjs.RPCImpl => {
    return (method: rpc.ServiceMethod<Message<{}>,Message<{}>>, requestData: Uint8Array) => {
        console.log(method);
        console.log(requestData);

        const p = fetch(hostname + '/twirp/twitch.twirp.example.Haberdasher/MakeHat', {
            method: 'POST',
            body: requestData,
            headers: {
                'Content-Type': 'application/protobuf'
            }
        });

        console.log(p);

        return p;
    };
};
