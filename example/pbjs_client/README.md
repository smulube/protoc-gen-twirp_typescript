## Generating protobuf.js client and testing

    npm install
    pbjs -t static-module -w commonjs -o ./service.pb.js ../service.proto
    pbts --no-comments -o ./service.pb.d.ts ./service.pb.js
    
    tsc main.ts && webpack-cli --entry ./main.js --output browser.js --mode development
    python -m SimpleHTTPServer
    
    http://localhost:8000
