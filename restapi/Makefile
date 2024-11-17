GOCMD=GO111MODULE=on GOARCH=amd64 go
GOBUILD=$(GOCMD) build -v 
GOTEST=$(GOCMD) test -v 
GOHOSTOS=$(strip $(shell $(GOCMD) env get GOHOSTOS))


# Active module mode, as we use go modules to manage dependencies
export GO111MODULE=on

COMMIT ?= $(shell git describe --always)
BUILD_DATE ?= $(shell date -u +%m/%d/%Y)
LDFLAGS="-X main.commit=${COMMIT} -X main.date=${BUILD_DATE}"

OUTWIN=bin/personacmms.exe
OUTWINDEBUG=bin/personacmms-debug.exe
OUTLIN=bin/personacmms
OUTLINDEBUG=bin/personacmms-debug

#
PKG := 

all: tidy format build test

clean:
	go clean
	rm -rf bin

.PHONY: tidy
tidy:
	go mod tidy

format:
	gofmt -s -w internal/ cmd/ 

build: build-windows build-linux

build-windows:
	GOOS=windows $(GOBUILD) -ldflags $(LDFLAGS) -o $(OUTWIN) github.com/jtcarden0001/personacmms/restapi/cmd/main

build-windows-debug:
	GOOS=windows $(GOBUILD) -gcflags "all=-N -l" -ldflags $(LDFLAGS) -o $(OUTWINDEBUG) github.com/jtcarden0001/personacmms/restapi/cmd/main

build-linux:
	GOOS=linux $(GOBUILD) -ldflags $(LDFLAGS) -o $(OUTLIN) github.com/jtcarden0001/personacmms/restapi/cmd/main

build-linux-debug:
	GOOS=linux $(GOBUILD) -gcflags "all=-N -l" -ldflags $(LDFLAGS) -o $(OUTLINDEBUG) github.com/jtcarden0001/personacmms/restapi/cmd/main

swag:
	swag init -g ../../internal/api/v1/gin/gin.go -o internal/api/v1/docs -d cmd/main,internal/types,internal/api/v1/gin

test: unittest

unittest: 
	$(GOTEST) ./internal/...

