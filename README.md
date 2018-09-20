```sh
go get github.com/nvcnvn/twirp_example
go run $GOPATH/src/github.com/nvcnvn/twirp_example/cmd/supervisorserver/main.go
# in another terminal
go run $GOPATH/src/github.com/nvcnvn/twirp_example/cmd/haberdasher/main.go
```
This two programs will listen on port 9990 and 9991.  
In another terminal:
```sh
# make sure you have curl and protobuff tools
# for testing protobuff
$GOPATH/src/github.com/nvcnvn/twirp_example/tools/scripts/testscripts/make_hat_protobuff.sh
# or json
$GOPATH/src/github.com/nvcnvn/twirp_example/tools/scripts/testscripts/make_hat_json.sh
```