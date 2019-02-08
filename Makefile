# The name of the binary.
# Also the default name for Docker.
BINARY_NAME:=go-endpoint-testing-with-godog

.PHONY: all
all: build run

.PHONY: initgo
initgo:
	GOARCH=$(DEFAULT_GOARCH) GOOS=$(DEFAULT_GOOS) GOPATH=$${HOME}/go GOBIN=$${HOME}/go/bin go get -v github.com/DATA-DOG/godog
	GOARCH=$(DEFAULT_GOARCH) GOOS=$(DEFAULT_GOOS) GOPATH=$${HOME}/go GOBIN=$${HOME}/go/bin go get -v github.com/christianhujer/assert
	go generate
	GOARCH=$(DEFAULT_GOARCH) GOOS=$(DEFAULT_GOOS) GOPATH=$${HOME}/go GOBIN=$${HOME}/go/bin go get -v .

.PHONY: build
## Performs the build only, without running tests.
build:
	go fmt
	go build -o $(BINARY_NAME)

.PHONY: clean
clean::
	$(RM) -r $(BINARY_NAME)-* target/


.PHONY: run
run:
	./$(BINARY_NAME)

#-include User.mk
#
#.PHONY: test
#test:
#	go test