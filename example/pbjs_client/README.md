## Generating protobuf.js client

    npm install
    pbjs -t static-module -o ./service.pb.js -w commonjs ../service.proto
    pbts -o ./service.pb.d.ts ./service.pb.js

