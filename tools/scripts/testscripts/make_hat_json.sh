curl --request "POST" \
     --location "http://localhost:9990/twirp/twirp.example.haberdasher.Haberdasher/MakeHat" \
     --header "Content-Type:application/json" \
     --data '{"inches": 11}' \
     --verbose