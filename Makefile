IMPORT_PATH := github.com/macroadster/hicc
GO          := $(or $(shell command -v go 2> /dev/null),which go)
BUILD_DIR := $(CURDIR)/build

all: init build

build: init fmt vet
	$(GO) build -o ${BUILD_DIR}/hicc
docker-build: init fmt vet
init:
	mkdir -p ${BUILD_DIR}
fmt:
	$(GO) fmt $(IMPORT_PATH)...
vet:
	$(GO) vet -composites=false $(IMPORT_PATH)...
docker: init fmt vet
	env GOOS=linux GOARCH=amd64 $(GO) build -o ${BUILD_DIR}/hicc
	docker build -t hicc .
clean:
	rm -rf ${BUILD_DIR}
