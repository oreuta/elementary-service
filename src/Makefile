all:get glide install vet test run

test:
	-go test -v -cover ./... > tests.xml
get:
	curl https://glide.sh/get | sh
	go get -u gopkg.in/alecthomas/gometalinter.v2
glide:
	glide update
install:
	gometalinter.v2 --install
vet:
	-gometalinter.v2 --vendor --tests --skip=mocks --disable=gotype --disable=errcheck --disable=gas --disable=dupl --deadline=1500s --checkstyle --sort=linter ./... > static-analysis.xml
run:
	go run main.go
    
