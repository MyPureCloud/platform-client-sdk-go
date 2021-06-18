PACKAGE_NAME = platformclientv2

.PHONY: all test

all: test build ## Build and run tests

build: windows linux darwin ## Build binaries

windows:
	env GOOS=windows GOARCH=amd64 go build -v ./${PACKAGE_NAME}

linux:
	env GOOS=linux GOARCH=amd64 go build -v ./${PACKAGE_NAME}

darwin:
	env GOOS=darwin GOARCH=amd64 go build -v ./${PACKAGE_NAME}

test:
	go test ./${PACKAGE_NAME} -v -failfast
