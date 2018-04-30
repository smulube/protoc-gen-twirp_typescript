import {twitch} from './service.pb';
import Haberdasher = twitch.twirp.example.Haberdasher;
import {createTwirpAdapter} from './service.twirp';

const hostname = 'http://localhost:8080';
const twirpPrefix = '/twirp/twitch.twirp.example.Haberdasher/';
const haberdasher = Haberdasher.create(createTwirpAdapter(hostname + twirpPrefix));

haberdasher.makeHat({inches: 10})
    .then((hat) => console.log(hat))
    .catch((err) => console.error(err));
