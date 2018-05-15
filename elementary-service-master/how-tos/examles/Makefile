TARGETNAME=bla-bla-service

packages =  ./models \
            ./services

# global command
.PHONY: all
all: dependencies localbuild test code-quality cover

.PHONY: local
local: dependencies localbuild test

.PHONY: localbuild
localbuild:
	go build -o $(TARGETPATH)

.PHONY: service-linux
platform-webroot-service-linux: bindata
	GOOS=linux GOARCH=amd64 go build -o $(GOPATH)/bin/linux_amd64/$(TARGETNAME)

.PHONY: service-mac
platform-webroot-service-mac: bindata
	GOOS=darwin GOARCH=amd64 go build -o $(GOPATH)/bin/darwin_amd64/$(TARGETNAME)


.PHONY: dependencies
dependencies:
	echo "Installing dependencies"
    go get code.google.com/p/go.tools/cmd/cover
    go get github.com/t-yuki/gocover-cobertura
    go get -u gopkg.in/alecthomas/gometalinter.v2
	glide install

.PHONY: test
test:
	@$(foreach package,$(packages), \
		set -e; \
		go test -coverprofile $(package)/cover.out -covermode=count $(package);)

.PHONY: cover
cover: test
	echo "mode: count" > cover-all.out
	@$(foreach package,$(packages), \
		tail -n +2 $(package)/cover.out >> cover-all.out;)
	gocover-cobertura < cover-all.out > cover-cobertura.xml

.PHONY: show
show:
	echo "Launching web browser to show overall coverage..."
	go tool cover -html=cover-all.out

.PHONY: code-quality
code-quality:
	gometalinter --vendor --tests --skip=mocks \
		--exclude='_gen.go' --exclude='docs.go' \
		--disable=gotype --disable=errcheck --disable=gas --disable=dupl \
		--deadline=1500s --checkstyle --sort=linter ./... > static-analysis.xml
