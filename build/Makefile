PACKAGE_NAME = platformclientv2
ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

build:
	cd ${ROOT_DIR}; go build -o ./bin/${PACKAGE_NAME} ./${PACKAGE_NAME}

test:
	cd ${ROOT_DIR}; go test -v -failfast ./${PACKAGE_NAME}
