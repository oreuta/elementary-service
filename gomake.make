all : test get run
  
get:
    go get -u gopkg.in/alecthomas/gometalinter.v2
    gometalinter.v2 --install

test: 
    gometalinter ./...
    go test

run:
    go run main.go
    
