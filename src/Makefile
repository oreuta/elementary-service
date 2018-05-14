all:get metal  run
  
get:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install

metal: 
	gometalinter.v2 ./...
	go test ./...
run:
	go run main.go
    
