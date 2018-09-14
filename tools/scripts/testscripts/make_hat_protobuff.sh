echo "inches:101" \
    | protoc --proto_path=$GOPATH/src/github.com/nvcnvn/twirp_example/rpc/haberdasher --encode twirp.example.haberdasher.Size service.proto \
    | curl -s --request POST \
                  --location http://localhost:9990/twirp/twirp.example.haberdasher.Haberdasher/MakeHat \
                  --header "Content-Type: application/protobuf" \
                  --data-binary @- \
    | protoc --proto_path=$GOPATH/src/github.com/nvcnvn/twirp_example/rpc/haberdasher --decode twirp.example.haberdasher.Hat service.proto