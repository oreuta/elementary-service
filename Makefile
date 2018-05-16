all:get install vet test run

test:
	-go test -v -cover ./... > tests.xml
get:
	go get -u gopkg.in/alecthomas/gometalinter.v2
install:
	gometalinter.v2 --install
vet:
	-gometalinter.v2 --vendor --tests --skip=mocks --exclude='_gen.go' --exclude='docs.go' --disable=gotype --disable=errcheck --disable=gas --disable=dupl --deadline=1500s --checkstyle --sort=linter ./... > static-analysis.xml
run:
	go run src/main.go
    
