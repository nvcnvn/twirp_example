```sh
go get github.com/nvcnvn/twirp_example
go run $GOPATH/src/github.com/nvcnvn/twirp_example/cmd/http/main.go
```
This program will listen on port 9990.  
In another terminal:
```sh
# make sure you have curl and protobuff tools
# for testing protobuff
$GOPATH/src/github.com/nvcnvn/twirp_example/tools/scripts/testscripts/make_hat_protobuff.sh
# or json
$GOPATH/src/github.com/nvcnvn/twirp_example/tools/scripts/testscripts/make_hat_json.sh
```