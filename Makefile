all:test get vet run
  
get:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install
vet:
	-gometalinter.v2 ./...
test:
	-go test -v -cover ./...
run:
	go run src/main.go
    
