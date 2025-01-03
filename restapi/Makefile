GOCMD=GO111MODULE=on GOARCH=amd64 go
GOBUILD=$(GOCMD) build -v 
GOTEST=$(GOCMD) test -parallel 3
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

all: clean tidy format swag openapi go-sdk build test

clean:
	go clean
	rm -rf bin

.PHONY: tidy

# clean up and organize
tidy:
	go mod tidy

format:
	gofmt -s -w internal/ cmd/ 

# code generation
swag:
	swag init -g ../../internal/api/v1/gin/gin.go -o internal/api/v1/docs -d cmd/main,internal/types,internal/api/v1/gin
	swag fmt

openapi:
	docker run --rm -v /workspaces/personacmms/restapi/internal/api/v1/docs:/usr/src/app mermade/swagger2openapi swagger2openapi --yaml --outfile openapi.yaml swagger.json

go-sdk: 
	rm -rf /workspaces/personacmms/sdk/v1/go
	docker run --rm -v "/workspaces/personacmms:/local" openapitools/openapi-generator-cli generate \
    -i /local/restapi/internal/api/v1/docs/openapi.yaml \
    -g go \
	--git-repo-id personacmms/sdk/v1/go \
	--git-user-id jtcarden0001 \
    -o /local/sdk/v1/go
	cd /workspaces/personacmms/sdk/v1/go; go get github.com/stretchr/testify/assert; go get golang.org/x/net/context; go test ./...
	

# build
build: build-windows build-linux

build-windows:
	GOOS=windows $(GOBUILD) -ldflags $(LDFLAGS) -o $(OUTWIN) github.com/jtcarden0001/personacmms/restapi/cmd/main

build-windows-debug:
	GOOS=windows $(GOBUILD) -gcflags "all=-N -l" -ldflags $(LDFLAGS) -o $(OUTWINDEBUG) github.com/jtcarden0001/personacmms/restapi/cmd/main

build-linux:
	GOOS=linux $(GOBUILD) -ldflags $(LDFLAGS) -o $(OUTLIN) github.com/jtcarden0001/personacmms/restapi/cmd/main

build-linux-debug:
	GOOS=linux $(GOBUILD) -gcflags "all=-N -l" -ldflags $(LDFLAGS) -o $(OUTLINDEBUG) github.com/jtcarden0001/personacmms/restapi/cmd/main

# test
test: store-test app-test

store-test:
	$(GOTEST) ./internal/store/postgres/...

app-test:
	$(GOTEST) ./internal/app/cmmsapp/...

# run
dev-up: clean
	docker-compose -f ./tools/docker/dev.compose.yaml --env-file ./tools/docker/dev.env up --build -d

dev-down:
	docker-compose -f ./tools/docker/dev.compose.yaml --env-file ./tools/docker/dev.env down -v

