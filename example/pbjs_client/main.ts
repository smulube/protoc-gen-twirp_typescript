import 'isomorphic-fetch';
import {twitch} from './service.pb';
import Haberdasher = twitch.twirp.example.Haberdasher;
import {createTwirpClient} from './service.twirp';

const hostname = 'http://localhost:8080';

console.log('Using: ' + hostname);

const haberdasher = Haberdasher.create(createTwirpClient(hostname));

haberdasher.makeHat({inches: 10})
    .then((hat) => {
        console.log('got hat!');
        console.log(hat);
    })
    .catch((err) => console.error(err));
