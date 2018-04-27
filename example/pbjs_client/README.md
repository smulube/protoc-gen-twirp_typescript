## Generating protobuf.js client

    npm install
    pbjs -t static-module -w commonjs -o ./service.pb.js ../service.proto
    pbts --no-comments -o ./service.pb.d.ts ./service.pb.js

