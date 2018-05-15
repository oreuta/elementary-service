all:test get install vet run
get:
	go get -u gopkg.in/alecthomas/gometalinter.v2
install:
	gometalinter.v2 --install
vet:
	-gometalinter.v2 ./...
test:
	-go test -v -cover ./...
run:
	go run src/main.go
    
