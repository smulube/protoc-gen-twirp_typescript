import {createTwirpHaberdasher} from './service.twirp';

const haberdasher = createTwirpHaberdasher('http://localhost:8080');

haberdasher.makeHat({inches: 10})
    .then((hat) => console.log(hat))
    .catch((err) => console.error(err));
