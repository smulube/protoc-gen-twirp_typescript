import 'isomorphic-fetch';
import {twitch} from './service.pb';
import Haberdasher = twitch.twirp.example.Haberdasher;
import {createTwirpAdapter} from './service.twirp';

const hostname = 'http://localhost:8080';
const haberdasher = Haberdasher.create(createTwirpAdapter(hostname));

haberdasher.makeHat({inches: 10})
    .then((hat) => console.log(hat))
    .catch((err) => console.error(err));
