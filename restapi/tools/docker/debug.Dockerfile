# docker build -t restapi-dev -f tools/docker/dev.Dockerfile -p 8080:8080 .
FROM golang:1.23.3

WORKDIR /app

COPY . .

RUN go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080
EXPOSE 40000

WORKDIR /app/cmd/main

CMD ["dlv", "debug", "--headless", "--listen=:40000", "--api-version=2", "--accept-multiclient", "."]