GOCMD=GO111MODULE=on GOARCH=amd64 go
GOBUILD=$(GOCMD) build -v 
GOTEST=$(GOCMD) test -v 
GOHOSTOS=$(strip $(shell $(GOCMD) env get GOHOSTOS))


# Active module mode, as we use go modules to manage dependencies
export GO111MODULE=on

COMMIT ?= $(shell git describe --always)
BUILD_DATE ?= $(shell date -u +%m/%d/%Y)
LDFLAGS="-X main.commit=${COMMIT} -X main.date=${BUILD_DATE}"

OUTEXE=bin/cmmsctl.exe
OUT=bin/cmmsctl

#
PKG := 

all: tidy format build

clean:
	go clean
	rm -rf bin

.PHONY: tidy
tidy:
	go mod tidy

format:
	gofmt -s -w cmd/ 

build: build-windows build-linux

build-windows:
	GOOS=windows $(GOBUILD) -ldflags $(LDFLAGS) -o $(OUTEXE) github.com/jtcarden0001/personacmms/cli/cmd/main

build-linux:
	GOOS=linux $(GOBUILD) -ldflags $(LDFLAGS) -o $(OUT) github.com/jtcarden0001/personacmms/cli/cmd/main