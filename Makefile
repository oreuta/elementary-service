<<<<<<< HEAD
GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/go-getting-started

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web
=======
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
    
>>>>>>> e196e071311a4665f7277bf6c1efe6c77f61241d
